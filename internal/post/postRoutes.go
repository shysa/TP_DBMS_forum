package post

import (
	"github.com/gin-gonic/gin"
	"github.com/shysa/TP_DBMS/app/database"
)

func AddPostRoutes(r *gin.Engine, db *database.DB)  {
	handler := NewHandler(db)

	postGroup := r.Group("/post/:id")
	{
		postGroup.GET("/details", handler.GetPostDetails)
		postGroup.POST("/details", handler.UpdatePostDetails)
	}
}
