package models

import (
	"github.com/ismailawa/todo-api/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Password  string `json:"password"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func CreateUser(user *User) *User {
	db.NewRecord(user)
	db.Create(&user)
	return user
}

func GetUsers() []User {
	var users []User
	db.Find(&users)
	return users
}

func GetUser(id int64) (*User, *gorm.DB) {
	var user User
	db := db.Where("ID=?", id).Find(&user)
	return &user, db
}

func DeleteUser(id int64) User {
	var user User
	db.Where("ID=?", id).Delete(user)
	return user
}
