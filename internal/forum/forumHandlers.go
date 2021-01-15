package forum

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/shysa/TP_DBMS/app/database"
	"github.com/shysa/TP_DBMS/internal/models"
	"log"
	"net/http"
	"strings"
)

type Handler struct {
	repo *database.DB
}

func NewHandler(db *database.DB) *Handler {
	return &Handler{
		repo: db,
	}
}

func (h *Handler) CreateForum(c *gin.Context) {
	f := models.Forum{}
	if err := c.BindJSON(&f); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Error{Error: "[BindJSON]: " + err.Error()})
		return
	}

	if err := h.repo.QueryRow(context.Background(),
		`insert into forum(slug, title, "user") values($1, $2, (select nickname from users where nickname=$3)) returning "user"`, f.Slug, f.Title, f.User).Scan(&f.User); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				switch pgErr.Code {
				case pgerrcode.NotNullViolation:
					// 404 user not found
					c.JSON(http.StatusNotFound, errors.New(fmt.Sprintf("Can't find user with nickname: %s", f.User)))
					return
				case pgerrcode.UniqueViolation:
					// 409 forum exists
					if err := h.repo.QueryRow(context.Background(),`select slug, title, "user" from forum where slug=$1`, f.Slug).Scan(&f.Slug, &f.Title, &f.User); err != nil {
						fmt.Println(err.Error())
					}
					c.JSON(http.StatusConflict, f)
					return
				}
			}
	}
	c.JSON(http.StatusCreated, f)
}

func (h *Handler) GetForumDetails(c *gin.Context) {
	slug := c.Param("slug")
	f := models.Forum{}

	if err := h.repo.QueryRow(context.Background(),`select slug, title, "user", posts, threads from forum where slug=$1`, slug).Scan(&f.Slug, &f.Title, &f.User, &f.Posts, &f.Threads); err != nil {
		c.JSON(http.StatusNotFound, errors.New(fmt.Sprintf("Can't find forum with slug: %s", slug)))
		return
	}
	c.JSON(http.StatusOK, f)
}

func (h *Handler) CreateForumThread(c *gin.Context) {
	slug := c.Param("slug")
	t := models.Thread{}
	if err := c.BindJSON(&t); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Error{Error: "[BindJSON]: " + err.Error()})
		return
	}

	var query strings.Builder
	query.WriteString("insert into thread(author, created, forum, message, title")
	var values strings.Builder
	values.WriteString(") values ((select nickname from users u where u.nickname=$1), $2, (select slug from forum f where f.slug=$3), $4, $5")
	qValues := make([]interface{}, 0, 6)
	qValues = append(qValues, t.Author, t.Created, slug, t.Message, t.Title)

	if t.Slug != "" {
		query.WriteString(", slug")
		values.WriteString(" , $6")
		qValues = append(qValues, t.Slug)
	}

	values.WriteString(") returning id, forum, author")

	if err := h.repo.QueryRow(context.Background(), query.String() + values.String(), qValues...).Scan(&t.Id, &t.Forum, &t.Author); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				if err := h.repo.QueryRow(context.Background(),
					"select id, author, created, forum, message, title, slug from thread where slug=$1", t.Slug).Scan(&t.Id, &t.Author, &t.Created, &t.Forum, &t.Message, &t.Title, &t.Slug); err != nil {
					fmt.Println(err.Error())
				}
				c.JSON(http.StatusConflict, t)
				return
			}
			if pgErr.Code == pgerrcode.NotNullViolation {
				c.JSON(http.StatusNotFound, errors.New(fmt.Sprintf("Can't find forum with slug: %s, or username: %s", slug, t.Author)))
				return
			}
		}
	}

	if _, err := h.repo.Exec(context.Background(), "insert into forum_users(forum, nickname) values ($1, $2) on conflict do nothing", slug, t.Author); err != nil {
		log.Fatal("can't add forum users to forum_users table: ", err.Error())
	}

	c.JSON(http.StatusCreated, t)
}

func (h *Handler) GetForumUsers(c *gin.Context) {
	slug := c.Param("slug")
	params := models.Params{
		Limit: 100,
		Desc: false,
	}
	_ = c.Bind(&params)

	i := 1
	var values []interface{}

	var query strings.Builder
	query.WriteString("select u.nickname, u.about, u.email, u.fullname from users u join forum_users fu on u.nickname=fu.nickname where fu.forum=$1")
	values = append(values, slug)

	if params.Since != "" {
		values = append(values, params.Since)
		i++
		if params.Desc {
			query.WriteString(fmt.Sprintf(" and fu.nickname < $%d", i))
		} else {
			query.WriteString(fmt.Sprintf(" and fu.nickname > $%d", i))
		}
	}
	query.WriteString(" order by fu.nickname")
	if params.Desc {
		query.WriteString(" desc")
	}
	query.WriteString(fmt.Sprintf(" limit %d", params.Limit))

	ul := models.Users{}
	rows, _ := h.repo.Query(context.Background(), query.String(), values...)
	for rows.Next() {
		u := models.User{}
		if err := rows.Scan(&u.Nickname, &u.About, &u.Email, &u.Fullname); err != nil {
			log.Fatal("cant scan rows: ", err.Error())
		}
		ul = append(ul, u)
	}
	rows.Close()

	if len(ul) == 0 {
		var existing string
		if err := h.repo.QueryRow(context.Background(),"select slug from forum where slug=$1", slug).Scan(&existing); err != nil {
			c.JSON(http.StatusNotFound, errors.New(fmt.Sprintf("Can't find forum with slug: %s", slug)))
			return
		}
	}

	c.JSON(http.StatusOK, ul)
}

func (h *Handler) GetForumThreads(c *gin.Context) {
	slug := c.Param("slug")
	params := models.Params{
		Limit: 100,
		Desc: false,
	}
	_ = c.Bind(&params)

	t := models.Threads{}

	i := 1
	var values []interface{}
	var query strings.Builder
	query.WriteString("select id, author, created, forum, message, case slug is null when true then '' else slug end, title, votes from thread where forum=$1")
	values = append(values, slug)

	if params.Since != "" {
		values = append(values, params.Since)
		i++
		if params.Desc {
			query.WriteString(fmt.Sprintf(" and created <= $%d", i))
		} else {
			query.WriteString(fmt.Sprintf(" and created >= $%d", i))
		}
	}
	query.WriteString(" order by created")
	if params.Desc {
		query.WriteString(" desc")
	}
	query.WriteString(fmt.Sprintf(" limit %d", params.Limit))

	rows, _ := h.repo.Query(context.Background(), query.String(), values...)
	for rows.Next() {
		nt := models.Thread{}
		if err := rows.Scan(&nt.Id, &nt.Author, &nt.Created, &nt.Forum, &nt.Message, &nt.Slug, &nt.Title, &nt.Votes); err != nil {
			log.Fatal("cant scan rows: ", err.Error())
		}
		t = append(t, nt)
	}
	rows.Close()

	if len(t) == 0 {
		var existing int
		if err := h.repo.QueryRow(context.Background(),"select id from forum where slug=$1", slug).Scan(&existing); err != nil {
			c.JSON(http.StatusNotFound, errors.New(fmt.Sprintf("Can't find forum with slug: %s", slug)))
			return
		}
	}

	c.JSON(http.StatusOK, t)
}
