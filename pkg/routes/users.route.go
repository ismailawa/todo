package routes

import (
	"github.com/gorilla/mux"
	"github.com/ismailawa/todo-api/pkg/controllers"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
}
