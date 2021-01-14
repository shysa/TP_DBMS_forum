package user

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
)

type Handler struct {
	repo *database.DB
}

func NewHandler(db *database.DB) *Handler {
	return &Handler{
		repo: db,
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	// TODO: sanitize validate?
	nickname := c.Param("nickname")

	u := models.User{}
	if err := c.BindJSON(&u); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Error{Error: "[BindJSON]: " + err.Error()})
		return
	}
	u.Nickname = nickname

	row, err := h.repo.Exec(context.Background(),"insert into users(nickname, about, email, fullname) values($1, $2, $3, $4)", nickname, u.About, u.Email, u.Fullname)
	if err != nil {
		log.Print(err.Error())
	}

	if row == nil {
		existings := models.Users{}
		rows, _ := h.repo.Query(context.Background(),"select nickname, about, email, fullname from users where nickname=$1 or email=$2", u.Nickname, u.Email)
		defer rows.Close()
		for rows.Next() {
			eu := models.User{}
			if err := rows.Scan(&eu.Nickname, &eu.About, &eu.Email, &eu.Fullname); err != nil {
				log.Fatal("cant scan rows: ", err.Error())
			}
			existings = append(existings, eu)
		}
		c.JSON(http.StatusConflict, existings)
		return
	}

	c.JSON(http.StatusCreated, u)
}

func (h *Handler) GetUser(c *gin.Context)  {
	nickname := c.Param("nickname")

	u := models.User{}
	if err := h.repo.QueryRow(context.Background(),"select nickname, about, email, fullname from users where nickname=$1", nickname).Scan(&u.Nickname, &u.About, &u.Email, &u.Fullname); err != nil {
		c.JSON(http.StatusNotFound, errors.New(fmt.Sprintf("Can't find user with nickname: %s", nickname)))
		return
	}

	c.JSON(http.StatusOK, u)
}

func (h *Handler) UpdateUser(c *gin.Context)  {
	nickname := c.Param("nickname")

	// 404
	var exists int
	if err := h.repo.QueryRow(context.Background(),"select id from users where nickname=$1", nickname).Scan(&exists); err != nil {
		c.JSON(http.StatusNotFound, errors.New(fmt.Sprintf("Can't find user with nickname: %s", nickname)))
		return
	}

	u := models.User{
		Nickname: nickname,
	}
	upd := &models.UserUpdate{}
	if err := c.BindJSON(&upd); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Error{Error: "[BindJSON]: " + err.Error()})
		return
	}

	i := 0
	query := "update users set "
	var queryParams []string
	var values []interface{}

	if upd.About != "" {
		i++
		queryParams = append(queryParams, "about=$"+strconv.Itoa(i))
		values = append(values, upd.About)
	}
	if upd.Email != "" {
		i++
		queryParams = append(queryParams, "email=$"+strconv.Itoa(i))
		values = append(values, upd.Email)
	}
	if upd.Fullname != "" {
		i++
		queryParams = append(queryParams, "fullname=$"+strconv.Itoa(i))
		values = append(values, upd.Fullname)
	}

	query += strings.Join(queryParams, ",")
	i++
	query += " where nickname=$" + strconv.Itoa(i)
	query += " returning nickname, about, email, fullname"
	values = append(values, nickname)

	// {} in query returns full profile
	if i == 1 {
		if err := h.repo.QueryRow(context.Background(),"select about, email, fullname from users where nickname=$1", nickname).Scan(&u.About, &u.Email, &u.Fullname); err != nil {
			fmt.Println(err.Error())
			return
		}
	} else {
		if err := h.repo.QueryRow(context.Background(), query, values...).Scan(&u.Nickname, &u.About, &u.Email, &u.Fullname); err != nil {
			c.JSON(http.StatusConflict, errors.New("new user profile data conflicted with already existing users"))
			return
		}
	}

	// 200
	c.JSON(http.StatusOK, u)
}