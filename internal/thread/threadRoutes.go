package thread

import (
	"github.com/gin-gonic/gin"
	"github.com/shysa/TP_DBMS/app/database"
)

func AddThreadRoutes(r *gin.Engine, db *database.DB)  {
	handler := NewHandler(db)

	threadGroup := r.Group("/thread/:slug_or_id")
	{
		threadGroup.POST("/create", handler.CreateThreadPosts)
		threadGroup.GET("/details", handler.GetThreadDetails)
		threadGroup.POST("/details", handler.UpdateThreadDetails)
		threadGroup.GET("/posts", handler.GetThreadPosts)
		threadGroup.POST("/vote", handler.VoteThread)
	}
}

