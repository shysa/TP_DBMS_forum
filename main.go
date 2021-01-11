package main

import (
	"fmt"
	"github.com/shysa/TP_DBMS/app/database"
	"github.com/shysa/TP_DBMS/app/server"
	"github.com/shysa/TP_DBMS/config"
	"log"
)

func main() {
	config.Cfg = config.Init()

	dbConn := database.NewDB(&config.Cfg.DB)
	if err := dbConn.Open(); err != nil {
		log.Fatal("connection refused: ", err)
		return
	}
	defer dbConn.Close()
	fmt.Println("connected to DB")

	srv := server.New(config.Cfg, dbConn)
	fmt.Println("listening on ", srv.Addr)
	srv.ListenAndServe()
}
