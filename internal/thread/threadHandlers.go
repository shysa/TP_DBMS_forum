package thread

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/shysa/TP_DBMS/app/database"
	"github.com/shysa/TP_DBMS/internal/models"
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
	var t models.Thread
	p := make([]*models.Post, 0)

	if err := c.BindJSON(&p); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Error{Error: "[BindJSON]: " + err.Error()})
		return
	}

	// 404 for thread
	query := "select id, case slug is null when true then '' else slug end, forum from thread where "
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

	for i, post := range p {
		p[i].Forum = t.Forum
		p[i].Thread = int(t.Id)

		if post.Parent != 0 {
			var pid int
			if err := h.repo.QueryRow(context.Background(), "select id from post where id=$1 and thread=$2", post.Parent, post.Thread).Scan(&pid); err != nil {
				c.JSON(http.StatusConflict, errors.New("can't find parent post"))
				return
			}
		}
		var puid int
		if err := h.repo.QueryRow(context.Background(), "select id from users where nickname=$1", post.Author).Scan(&puid); err != nil {
			c.JSON(http.StatusNotFound, errors.New("can't find post author"))
			return
		}

		insertQuery.WriteString(fmt.Sprintf("($%d, $%d, $%d, $%d, array_append((select tree from post where id=$%d), (select last_value from post_id_seq)), $%d, $%d)", v, v+1, v+2, v+3, v+3, v+4, v+5))
		if i < len(p) - 1 {
			insertQuery.WriteString(",")
		}
		v += 6
		postValues = append(postValues, post.Author, post.Message, crt, post.Parent, t.Forum, int(t.Id))

		ufValues = append(ufValues, post.Author)
		insertUFValues = append(insertUFValues, fmt.Sprintf("($1, $%d)", i+2))
	}
	insertQuery.WriteString(" returning id, created")

	insertFUQuery.WriteString(strings.Join(insertUFValues, ","))
	insertFUQuery.WriteString(" on conflict do nothing")

	// ---------------- transaction begin ------------------------------------
	tx, err := h.repo.Begin(context.Background())
	if err != nil {
		fmt.Println("cant open transaction on creating posts: ", err.Error())
		return
	}
	defer tx.Rollback(context.Background())

	resRows, _ := tx.Query(context.Background(), insertQuery.String(), postValues...)
	idx := 0
	for resRows.Next() {
		if err := resRows.Scan(&p[idx].Id, &p[idx].Created); err != nil {
			fmt.Println("cant insert so big wow sorry such error: ", err.Error())
			return
		}
		idx++
	}
	resRows.Close()

	if _, err := tx.Exec(context.Background(), "update forum set posts=forum.posts+$1 where slug=$2", len(p), t.Forum); err != nil {
		fmt.Println("can't add posts to forum row: ", err.Error())
		return
	}

	if _, err := tx.Exec(context.Background(), insertFUQuery.String(), ufValues...); err != nil {
		fmt.Println("can't add forum users to forum_users table: ", err.Error())
		return
	}

	if err := tx.Commit(context.Background()); err != nil {
		fmt.Println("cant commit transaction on creating posts: ", err.Error())
		return
	}
	// ---------------- transaction end ------------------------------------

	atomic.AddUint32(&counter, uint32(len(p)))
	//if counter >= 1500000 {
	//	h.repo.Exec(context.Background(),"cluster users using users_pkey")
	//	h.repo.Exec(context.Background(),"cluster forum using forum_pkey")
	//	h.repo.Exec(context.Background(),"cluster thread using thread_forum_created_index")
	//	h.repo.Exec(context.Background(),"cluster post using parent_tree_1")
	//	h.repo.Exec(context.Background(),"cluster forum_users using forum_users_pkey")
	//}

	c.JSON(http.StatusCreated, p)
}

func (h *Handler) GetThreadDetails(c *gin.Context) {
	thread := c.Param("slug_or_id")
	t := models.Thread{}

	// 404
	query := "select id, author, created, forum, message, case slug is null when true then '' else slug end, title, votes from thread where "
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
	t := models.Thread{}
	upd := models.ThreadUpdate{}
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
	t := models.Thread{}
	params := models.Params{
		Limit: 100,
		Desc:  false,
		Sort:  "flat",
	}
	_ = c.Bind(&params)

	// 404
	tQuery := "select id, forum, case slug is null when true then '' else slug end from thread where "
	if _, err := strconv.ParseInt(thread, 10, 64); err != nil {
		// search by slug
		tQuery += "slug=$1"
	} else {
		// by id
		tQuery += "id=$1"
	}
	if err := h.repo.QueryRow(context.Background(), tQuery, thread).Scan(&t.Id, &t.Forum, &t.Slug); err != nil {
		c.JSON(http.StatusNotFound, errors.New("can't find thread with this id or thread"))
		return
	}

	i := 1
	var values []interface{}
	var query strings.Builder
	query.WriteString("select id, author, created, isedited, message, parent, forum, thread, tree from post where thread=$1")
	if params.Sort == "parent_tree" {
		query.WriteString(" and tree[1] in (select tree[1] from post where thread=$1 and array_length(tree,1)=1 ")
	}
	values = append(values, t.Id)

	if params.Since != "" {
		values = append(values, params.Since)
		i++
		if params.Desc {
			switch params.Sort {
			case "flat":
				query.WriteString(fmt.Sprintf(" and id < $%d", i))
				break
			case "tree":
				query.WriteString(fmt.Sprintf(" and tree < (select tree from post where id=$%d)", i))
				break
			case "parent_tree":
				query.WriteString(fmt.Sprintf(" and tree[1] < (select tree[1] from post where id=$%d)", i))
				break
			}
		} else {
			switch params.Sort {
			case "flat":
				query.WriteString(fmt.Sprintf(" and id > $%d", i))
				break
			case "tree":
				query.WriteString(fmt.Sprintf(" and tree > (select tree from post where id=$%d)", i))
				break
			case "parent_tree":
				query.WriteString(fmt.Sprintf(" and tree[1] > (select tree[1] from post where id=$%d)", i))
				break
			}
		}
	}
	query.WriteString(" order by ")
	switch params.Sort {
	case "flat":
		query.WriteString("id")
		break
	case "tree":
		query.WriteString("tree")
		break
	case "parent_tree":
		query.WriteString("tree[1]")
		break
	}
	if params.Desc {
		query.WriteString(" desc")
	}
	query.WriteString(fmt.Sprintf(" limit %d", params.Limit))
	if params.Sort == "parent_tree" {
		query.WriteString(") order by tree[1]")
		if params.Desc {
			query.WriteString(" desc")
		}
		query.WriteString(", tree")
	}

	pl := models.Posts{}
	rows, _ := h.repo.Query(context.Background(), query.String(), values...)
	for rows.Next() {
		p := models.Post{}
		if err := rows.Scan(&p.Id, &p.Author, &p.Created, &p.IsEdited, &p.Message, &p.Parent, &p.Forum, &p.Thread, &p.Tree); err != nil {
			fmt.Println("cant scan rows: ", err.Error())
			return
		}
		pl = append(pl, p)
	}
	rows.Close()

	c.JSON(http.StatusOK, pl)
}

func (h *Handler) VoteThread(c *gin.Context) {
	thread := c.Param("slug_or_id")
	t := models.Thread{}
	v := models.Vote{}
	if err := c.BindJSON(&v); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Error{Error: "[BindJSON]: " + err.Error()})
		return
	}

	// 404
	query := "select id, author, created, forum, message, case slug is null when true then '' else slug end, title, votes from thread where "
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

	prev := 0
	curr := 0
	if err := h.repo.QueryRow(context.Background(),
		"insert into votes(nickname, thread, voice) values ((select nickname from users where nickname=$1), $2, $3) on conflict on constraint unique_uservoice_for_thread do update set prev=votes.voice, voice=excluded.voice returning voice, prev",
		v.Nickname, t.Id, v.Voice).Scan(&curr, &prev); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.NotNullViolation {
				// 404 user not found
				c.JSON(http.StatusNotFound, errors.New(fmt.Sprintf("Can't find user with nickname: %s", v.Nickname)))
				return
			}
		}
	}
	t.Votes = t.Votes - (prev - curr)

	c.JSON(http.StatusOK, t)
}
