package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ismailawa/todo-api/pkg/models"
	"github.com/ismailawa/todo-api/pkg/utils"
)

var newUser models.User

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	utils.BodyParser(r, user)
	createdUser := models.CreateUser(user)
	res, _ := json.Marshal(createdUser)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

}

func GetUser(w http.ResponseWriter, r *http.Request) {

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
