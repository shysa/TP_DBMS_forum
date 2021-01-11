package forum

import (
	"github.com/gin-gonic/gin"
	"github.com/shysa/TP_DBMS/app/database"
	"net/http"
	"strings"
)

func AddForumRoutes(r *gin.Engine, db *database.DB)  {
	handler := NewHandler(db)

	forumGroup := r.Group("/forum/:slug")
	{
		forumGroup.POST("", func(c *gin.Context) {
			if strings.HasPrefix(c.Request.RequestURI, "/api/forum/create") {
				handler.CreateForum(c)
				return
			}
			c.AbortWithStatus(http.StatusNotFound)
		})
		forumGroup.POST("/create", handler.CreateForumThread)
		forumGroup.GET("/details", handler.GetForumDetails)
		forumGroup.GET("/threads", handler.GetForumThreads)
		forumGroup.GET("/users", handler.GetForumUsers)
	}
}
