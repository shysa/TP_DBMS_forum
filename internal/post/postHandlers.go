package post

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shysa/TP_DBMS/app/database"
	"github.com/shysa/TP_DBMS/internal/models"
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

func (h *Handler) GetPostDetails(c *gin.Context) {
	id := c.Param("id")
	params := models.Params{}
	_ = c.Bind(&params)

	query := "select * from post where id=$1"
	fields := models.PostDetails{Author: nil, Forum: nil, Post: &models.Post{}, Thread: nil}
	if err := h.repo.QueryRow(context.Background(), query, id).Scan(&fields.Post.Id, &fields.Post.Created, &fields.Post.IsEdited, &fields.Post.Message, &fields.Post.Parent,  &fields.Post.Tree, &fields.Post.Thread, &fields.Post.Author, &fields.Post.Forum); err != nil {
		c.JSON(http.StatusNotFound, errors.New(fmt.Sprintf("Can't find post with id: %s", id)))
		return
	}

	if strings.Contains(params.Related, "user") {
		fields.Author = &models.User{}
		query = "select * from users where nickname=$1"
		var uId int
		h.repo.QueryRow(context.Background(), query, fields.Post.Author).Scan(&uId, &fields.Author.Nickname, &fields.Author.About, &fields.Author.Email, &fields.Author.Fullname)
	}
	if strings.Contains(params.Related, "forum") {
		fields.Forum = &models.Forum{}
		query = "select * from forum where slug=$1"
		var fId int
		h.repo.QueryRow(context.Background(), query, fields.Post.Forum).Scan(&fId, &fields.Forum.Posts, &fields.Forum.Slug, &fields.Forum.Threads, &fields.Forum.Title, &fields.Forum.User)
	}
	if strings.Contains(params.Related, "thread") {
		fields.Thread = &models.Thread{}
		query = "select * from thread where id=$1"
		h.repo.QueryRow(context.Background(), query, fields.Post.Thread).Scan(&fields.Thread.Id, &fields.Thread.Created, &fields.Thread.Message, &fields.Thread.Slug, &fields.Thread.Title, &fields.Thread.Votes, &fields.Thread.Author, &fields.Thread.Forum)
	}

	c.JSON(http.StatusOK, fields)
}

func (h *Handler) UpdatePostDetails(c *gin.Context) {
	id := c.Param("id")
	p := models.Post{}
	upd := models.PostUpdate{}
	if err := c.BindJSON(&upd); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Error{Error: "[BindJSON]: " + err.Error()})
		return
	}

	// 404
	query := "select id, message, author, created, isedited, parent, forum, thread, tree from post where id=$1"
	if err := h.repo.QueryRow(context.Background(), query, id).Scan(&p.Id, &p.Message, &p.Author, &p.Created, &p.IsEdited, &p.Parent, &p.Forum, &p.Thread, &p.Tree); err != nil {
		c.JSON(http.StatusNotFound, errors.New("can't find post with this id"))
		return
	}

	if upd.Message != "" && upd.Message != p.Message {
		if err := h.repo.QueryRow(context.Background(), "update post set message=$1, isedited=$2 where id=$3 returning message", upd.Message, true, id).Scan(&p.Message); err != nil {
			c.JSON(http.StatusNotFound, errors.New(fmt.Sprintf("can't find post with id %s", id) ))
			return
		}
	}

	c.JSON(http.StatusOK, p)
}