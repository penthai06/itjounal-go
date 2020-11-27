package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// var (
// 	bundle    *i18n.Bundle
// 	localizer *i18n.Localizer
// )

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to the %s database", Dbdriver)
			log.Fatal("This is ther error:", err)
		} else {
			fmt.Printf("Connected to the %s database", Dbdriver)
		}
	}

	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(DbPort, addr string) {
	fmt.Println("Listening to port" + DbPort)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
