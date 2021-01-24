package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	userId int
	Name   string
	Email  string
}

var db *gorm.DB
var err error

var users = []User{
	{Name: "Hina", Email: "abc@gmail.com", userId: 1},
	{Name: "Sudha", Email: "tyi@gmail.com", userId: 2},
	{Name: "kiran", Email: "kol@gmail.com", userId: 3},
	{Name: "vbhu", Email: "uio@gmail.com", userId: 3},
	{Name: "qwe", Email: "zxc@gmail.com", userId: 4},
}

func main() {

	myRouter := mux.NewRouter()

	db, err := gorm.Open("postgres", "user=postgres password= postgres dbname=postgres  sslmode=disable ")

	if err != nil {
		panic("Failed to connect to database")
	}

	defer db.Close()

	db.AutoMigrate(&User{})

	for index := range users {

		db.Create(&users[index])

	}

	myRouter.HandleFunc("/users", GetUsers).Methods("GET")

	myRouter.HandleFunc("/users/{name}/{email}", GetUser).Methods("GET")

	myRouter.HandleFunc("/users/{name}", DeleteUser).Methods("DELETE")

	handler := cors.Default().Handler(myRouter)

	log.Fatal(http.ListenAndServe(":8080", handler))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	var user User

	db.First(&user, vars["id"])
	json.NewEncoder(w).Encode(&user)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	var user User
	db.First(&user, vars["id"])

	db.Delete(&user)

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(&users)
}
