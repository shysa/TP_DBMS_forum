package thread

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shysa/TP_DBMS/app/database"
	"github.com/shysa/TP_DBMS/internal/models"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

var counter uint32

type Handler struct {
	repo *database.DB
}

func NewHandler(db *database.DB) *Handler {
	return &Handler{
		repo: db,
	}
}

func (h *Handler) CreateThreadPosts(c *gin.Context) {
	thread := c.Param("slug_or_id")
	t := &models.Thread{}
	var p []*models.Post
	if err := c.BindJSON(&p); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &models.Error{Error: "[BindJSON]: " + err.Error()})
		return
	}

	// 404 for thread
	query := "select id, slug, forum from thread where "
	if _, err := strconv.ParseInt(thread, 10, 64); err != nil {
		// search by slug
		query += "slug=$1"
	} else {
		// by id
		query += "id=$1"
	}
	if err := h.repo.QueryRow(context.Background(), query, thread).Scan(&t.Id, &t.Slug, &t.Forum); err != nil {
		c.JSON(http.StatusNotFound, errors.New("can't find thread with this id or thread"))
		return
	}
	if len(p) == 0 {
		c.JSON(http.StatusCreated, p)
		return
	}

	var insertQuery strings.Builder
	insertQuery.WriteString("insert into post(author, message, created, parent, tree, forum, thread) values ")
	postValues := make([]interface{}, 0, len(p)*5)
	v := 1
	crt, _ := time.Parse(time.RFC3339Nano, time.Now().Format(time.RFC3339Nano))

	var insertFUQuery strings.Builder
	insertFUQuery.WriteString("insert into forum_users values ")
	insertUFValues := make([]string, 0, len(p))
	ufValues := make([]interface{}, 0, len(p)+1)
	ufValues = append(ufValues, t.Forum)

	for i := range p {
		p[i].Forum = t.Forum
		p[i].Thread = int(t.Id)

		if p[i].Parent != 0 {
			var pid int
			if err := h.repo.QueryRow(context.Background(), "select id from post where id=$1 and thread=$2", p[i].Parent, p[i].Thread).Scan(&pid); err != nil {
				c.JSON(http.StatusConflict, errors.New("can't find parent post"))
				return
			}
		}
		var puid int
		if err := h.repo.QueryRow(context.Background(), "select id from users where nickname=$1", p[i].Author).Scan(&puid); err != nil {
			c.JSON(http.StatusNotFound, errors.New("can't find post author"))
			return
		}

		insertQuery.WriteString("($" + strconv.Itoa(v) + ", $" + strconv.Itoa(v+1) + ", $" + strconv.Itoa(v+2) + ", $" + strconv.Itoa(v+3) + ",")
		if p[i].Parent != 0 {
			insertQuery.WriteString(" array_append((select tree from post where id=$" + strconv.Itoa(v+3) + "), (select last_value from post_id_seq)),")
		} else {
			insertQuery.WriteString(" array[(select last_value from post_id_seq)],")
		}
		insertQuery.WriteString(" $" + strconv.Itoa(v+4) + "," + " $" + strconv.Itoa(v+5) + ")")

		if i < len(p) - 1 {
			insertQuery.WriteString(",")
		}
		v += 6
		postValues = append(postValues, p[i].Author, p[i].Message, crt, p[i].Parent, t.Forum, int(t.Id))

		ufValues = append(ufValues, p[i].Author)
		insertUFValues = append(insertUFValues, "($1, $" + strconv.Itoa(i + 2) + ")")
	}
	insertQuery.WriteString(" returning id, created")

	insertFUQuery.WriteString(strings.Join(insertUFValues, ","))
	insertFUQuery.WriteString(" on conflict do nothing")

	// ---------------- transaction begin ------------------------------------
	tx, err := h.repo.Begin(context.Background())
	if err != nil {
		log.Fatal("cant open transaction on creating posts: ", err.Error())
	}
	defer tx.Rollback(context.Background())

	resRows, _ := tx.Query(context.Background(), insertQuery.String(), postValues...)
	idx := 0
	for resRows.Next() {
		if err := resRows.Scan(&p[idx].Id, &p[idx].Created); err != nil {
			log.Fatal("cant insert so big wow sorry such error: ", err.Error())
		}
		idx++
	}
	resRows.Close()

	if _, err := tx.Exec(context.Background(), "update forum set posts=forum.posts+$1 where slug=$2", len(p), t.Forum); err != nil {
		log.Fatal("can't add posts to forum row: ", err.Error())
	}

	if _, err := tx.Exec(context.Background(), insertFUQuery.String(), ufValues...); err != nil {
		log.Fatal("can't add forum users to forum_users table: ", err.Error())
	}

	if err := tx.Commit(context.Background()); err != nil {
		log.Fatal("cant commit transaction on creating posts: ", err.Error())
	}
	// ---------------- transaction end ------------------------------------

	atomic.AddUint32(&counter, uint32(len(p)))
	if counter == 1500000 {
		time.Sleep(5 * time.Second)
		h.repo.Exec(context.Background(),"cluster users using users_nickname_index")
		h.repo.Exec(context.Background(),"cluster forum using forum_slug_id_index")
		h.repo.Exec(context.Background(),"cluster thread using thread_forum_created_index")
		h.repo.Exec(context.Background(),"cluster post using parent_tree_1")
		h.repo.Exec(context.Background(),"cluster forum_users using forum_users_pkey")
	}

	c.JSON(http.StatusCreated, p)
}

func (h *Handler) GetThreadDetails(c *gin.Context) {
	thread := c.Param("slug_or_id")
	t := &models.Thread{}

	// 404
	query := "select id, author, created, forum, message, slug, title, votes from thread where "
	if _, err := strconv.ParseInt(thread, 10, 64); err != nil {
		// search by slug
		query += "slug=$1"
	} else {
		// by id
		query += "id=$1"
	}
	if err := h.repo.QueryRow(context.Background(), query, thread).Scan(&t.Id, &t.Author, &t.Created, &t.Forum, &t.Message, &t.Slug, &t.Title, &t.Votes); err != nil {
		c.JSON(http.StatusNotFound, errors.New("can't find thread with this id or thread"))
		return
	}
	c.JSON(http.StatusOK, t)
}

func (h *Handler) UpdateThreadDetails(c *gin.Context) {
	thread := c.Param("slug_or_id")
	t := &models.Thread{}
	upd := &models.ThreadUpdate{}
	if err := c.BindJSON(&upd); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &models.Error{Error: "[BindJSON]: " + err.Error()})
		return
	}

	// 404
	query := "select id, author, created, forum, slug, votes, message, title from thread where "
	if _, err := strconv.ParseInt(thread, 10, 64); err != nil {
		// search by slug
		query += "slug=$1"
	} else {
		// by id
		query += "id=$1"
	}
	if err := h.repo.QueryRow(context.Background(), query, thread).Scan(&t.Id, &t.Author, &t.Created, &t.Forum, &t.Slug, &t.Votes, &t.Message, &t.Title); err != nil {
		c.JSON(http.StatusNotFound, errors.New("can't find thread with this id or thread"))
		return
	}
	if upd.Title == "" && upd.Message == "" {
		c.JSON(http.StatusOK, t)
		return
	}

	i := 0
	updQuery := "update thread set "
	var values []interface{}

	if upd.Title != "" {
		i++
		updQuery += "title=$" + strconv.Itoa(i)
		values = append(values, upd.Title)
	}
	if upd.Message != "" {
		i++
		if i == 2 {
			updQuery += ","
		}
		updQuery += "message=$" + strconv.Itoa(i)
		values = append(values, upd.Message)
	}

	i++
	updQuery += " where id=$" + strconv.Itoa(i) + " returning title, message"
	values = append(values, t.Id)

	if err := h.repo.QueryRow(context.Background(), updQuery, values...).Scan(&t.Title, &t.Message); err != nil {
		fmt.Println("something went wrong with updating thread: ", err.Error())
	}
	c.JSON(http.StatusOK, t)
}

func (h *Handler) GetThreadPosts(c *gin.Context) {
	thread := c.Param("slug_or_id")
	t := &models.Thread{}
	params := &models.Params{
		Limit: 100,
		Desc:  false,
		Sort:  "flat",
	}
	_ = c.Bind(&params)

	// 404
	query := "select id, forum, slug from thread where "
	if _, err := strconv.ParseInt(thread, 10, 64); err != nil {
		// search by slug
		query += "slug=$1"
	} else {
		// by id
		query += "id=$1"
	}
	if err := h.repo.QueryRow(context.Background(), query, thread).Scan(&t.Id, &t.Forum, &t.Slug); err != nil {
		c.JSON(http.StatusNotFound, errors.New("can't find thread with this id or thread"))
		return
	}

	i := 1
	var values []interface{}
	query = "select id, author, created, isedited, message, parent, forum, thread, tree from post where thread=$1"
	if params.Sort == "parent_tree" {
		query += " and tree[1] in (select tree[1] from post where thread=$1 and array_length(tree,1)=1 "
	}
	values = append(values, t.Id)

	if params.Since != "" {
		values = append(values, params.Since)
		i++

		var sign = ""
		if params.Desc {
			sign = "<"
		} else {
			sign = ">"
		}

		switch params.Sort {
		case "flat":
			query += " and id" + sign + "$" + strconv.Itoa(i)
			break
		case "tree":
			query += " and tree" + sign + "(select tree from post where id=$" + strconv.Itoa(i) + ")"
			break
		case "parent_tree":
			query += " and tree[1]" + sign + "(select tree[1] from post where id=$" + strconv.Itoa(i) + ")"
			break
		}

	}
	query += " order by "
	switch params.Sort {
	case "flat":
		query += "id"
		break
	case "tree":
		query += "tree"
		break
	case "parent_tree":
		query += "tree[1]"
		break
	}
	if params.Desc {
		query += " desc"
	}
	query += " limit " + strconv.Itoa(params.Limit)
	if params.Sort == "parent_tree" {
		query += ") order by tree[1]"
		if params.Desc {
			query += " desc"
		}
		query += ", tree"
	}

	pl := models.Posts{}
	rows, _ := h.repo.Query(context.Background(), query, values...)
	defer rows.Close()
	for rows.Next() {
		p := models.Post{}
		if err := rows.Scan(&p.Id, &p.Author, &p.Created, &p.IsEdited, &p.Message, &p.Parent, &p.Forum, &p.Thread, &p.Tree); err != nil {
			log.Fatal("cant scan rows: ", err.Error())
		}
		pl = append(pl, p)
	}
	c.JSON(http.StatusOK, pl)
}

func (h *Handler) VoteThread(c *gin.Context) {
	thread := c.Param("slug_or_id")
	t := &models.Thread{}
	v := &models.Vote{}
	if err := c.BindJSON(&v); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Error{Error: "[BindJSON]: " + err.Error()})
		return
	}

	// 404
	query := "select id, author, created, forum, message, slug, title from thread where "
	if _, err := strconv.ParseInt(thread, 10, 64); err != nil {
		// search by slug
		query += "slug=$1"
	} else {
		// by id
		query += "id=$1"
	}
	if err := h.repo.QueryRow(context.Background(), query, thread).Scan(&t.Id, &t.Author, &t.Created, &t.Forum, &t.Message, &t.Slug, &t.Title); err != nil {
		c.JSON(http.StatusNotFound, errors.New("can't find thread with this id or thread"))
		return
	}
	if err := h.repo.QueryRow(context.Background(), "select nickname from users where nickname=$1", v.Nickname).Scan(&v.Nickname); err != nil {
		c.JSON(http.StatusNotFound, errors.New("can't find user for vote"))
		return
	}

	if err := h.repo.QueryRow(
		context.Background(),
		"with inserting as ("+
			"insert into votes(nickname, thread, voice) values ($1, $2, $3)"+
			"on conflict on constraint unique_uservoice_for_thread do update set prev=votes.voice, voice=excluded.voice "+
			"returning prev, voice"+
			")"+
			"update thread set votes=votes-(select prev-voice from inserting) where id=$2 returning votes", v.Nickname, t.Id, v.Voice,
	).Scan(&t.Votes); err != nil {
		log.Fatal("something went wrong on voice for thread: ", err.Error())
	}

	c.JSON(http.StatusOK, t)
}
