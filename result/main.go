package main

import (
	"fmt"
	"log"
	"net/http"

	"app/handlers"
	"app/middlewares"
	"app/models"

	"github.com/gorilla/mux"
)

func main() {
	port := "7070"
	fmt.Printf("Result Server Listening on port %s\n", port)

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.RootHandler)
	r.HandleFunc("/login", handlers.LoginPageHandler).Methods("GET")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/signup", handlers.SignupPageHandler).Methods("GET")
	r.HandleFunc("/signup", handlers.SignupHandler).Methods("POST")

	fs := http.FileServer(http.Dir("./resources"))
	r.PathPrefix("/resources/").Handler(http.StripPrefix("/resources/", fs))

	r.Use(middlewares.Logging)

	db, e := models.ConnectDB()
	if e != nil {
		log.Fatal("connect DB: ", e)
	}

	if err := models.CreateTable(db); err != nil {
		log.Println("create table: ", err)
	} else {
		log.Println("success to create adminuser table")
	}

	db.Close()

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
