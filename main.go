package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initiateRoute(){
	r := mux.NewRouter()

	r.HandleFunc("/todo",GetTodos).Methods("GET")
	r.HandleFunc("/todo/{id}",GetTodo).Methods("GET")
	r.HandleFunc("/todo",CreateTodo).Methods("POST")
	r.HandleFunc("/todo/{id}",UpdateTodo).Methods("PUT")
	r.HandleFunc("/todo/{id}",DeleteTodo).Methods("DELETE")
	r.HandleFunc("/todo/{id}",Complete).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000",r))
}

func main(){
	//initiateRoute()
	initialMigration()
	initiateRoute()

}
