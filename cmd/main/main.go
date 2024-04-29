package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ismailawa/todo-api/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterUserRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
