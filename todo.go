package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Todo struct{
	gorm.Model
	Todo string `json:"todo"`
	Completed bool `json:"completed"`

}

var db *gorm.DB
var err error
var DNS string = "host=localhost user=postgres password=12345 dbname=todo port=5432 sslmode=disable"

func initialMigration(){
	db , err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil{
		fmt.Println(err.Error())
		panic("cannot connect to db")
	}
	db.AutoMigrate(&Todo{})
}





func GetTodos(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var todos []Todo
	db.Find(&todos)
	json.NewEncoder(w).Encode(&todos)
}

func GetTodo(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	param := mux.Vars(r)
	var todo []Todo
	db.First(&todo,param["id"])
	json.NewEncoder(w).Encode(&todo)
}
func CreateTodo(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)
	db.Create(&todo)
	json.NewEncoder(w).Encode(&todo)
}

func UpdateTodo(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	param := mux.Vars(r)
	var todo Todo
	db.First(&todo,param["id"])
	json.NewDecoder(r.Body).Decode(&todo)
	db.Save(&todo)
	json.NewEncoder(w).Encode(&todo)

}

func DeleteTodo(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	param :=mux.Vars(r)
	var todo Todo
	db.Delete(&todo,param["id"])
	json.NewEncoder(w).Encode("todo successfully deleted")
}

func Complete(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	param :=mux.Vars(r)
	var todo Todo
	db.First(&todo,param["id"])
	json.NewDecoder(r.Body).Decode(&todo)
	db.Save(&todo)
	json.NewEncoder(w).Encode(&todo)
}