package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/teten777/crud_book_api/api/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (s *Server) Initialize(Dbdriver, Dbuser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", Dbuser, DbPassword, DbHost, DbPort, DbName)
		s.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Connect connect to %s database", Dbdriver)
			log.Fatal("this is a error : ", err)
		}
	}

	s.DB.Debug().AutoMigrate(&models.Kategori{}, &models.User{})

	s.Router = mux.NewRouter()

	s.initializeRoutes()
}

func (s *Server) Run(addr string) {
	fmt.Println("Listening port 8080")
	log.Fatal(http.ListenAndServe(addr, s.Router))
}
