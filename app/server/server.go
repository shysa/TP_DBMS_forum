package server

import (
	"fmt"
	_ "github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/shysa/TP_DBMS/app/database"
	"github.com/shysa/TP_DBMS/config"
	forum "github.com/shysa/TP_DBMS/internal/forum"
	post "github.com/shysa/TP_DBMS/internal/post"
	service "github.com/shysa/TP_DBMS/internal/service"
	thread "github.com/shysa/TP_DBMS/internal/thread"
	user "github.com/shysa/TP_DBMS/internal/user"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
)

func New(cfg *config.Config, db *database.DB) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	router := gin.Default()

	//pprof.Register(router)

	router.RouterGroup = *router.Group("/api")
	{
		forum.AddForumRoutes(router, db)
		user.AddUserRoutes(router, db)
		thread.AddThreadRoutes(router, db)
		post.AddPostRoutes(router, db)
		service.AddServiceRoutes(router, db)
	}

	return &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Server.Address, cfg.Server.Port),
		Handler: router,
	}
}
