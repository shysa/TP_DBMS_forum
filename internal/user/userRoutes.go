package user

import (
	"github.com/gin-gonic/gin"
	"github.com/shysa/TP_DBMS/app/database"
)

func AddUserRoutes(r *gin.Engine, db *database.DB)  {
	handler := NewHandler(db)

	userGroup := r.Group("/user")
	{
		userGroup.POST("/:nickname/create", handler.CreateUser)
		userGroup.GET("/:nickname/profile", handler.GetUser)
		userGroup.POST("/:nickname/profile", handler.UpdateUser)
	}
}
