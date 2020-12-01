package controllers

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			DbUser,
			DbPassword,
			DbHost,
			DbPort,
			DbName,
		)
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
	pCrt, _ := filepath.Abs("./server.crt")
	pKey, _ := filepath.Abs("./server.key")
	// https://github.com/denji/golang-tls
	log.Fatal(http.ListenAndServeTLS(addr, pCrt, pKey, server.Router))
}
