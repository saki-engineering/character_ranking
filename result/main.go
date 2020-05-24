package main

import (
	"fmt"
	"log"
	"net/http"

	"app/models"
	"app/routers"
	"app/stores"
)

func main() {
	port := "7070"
	fmt.Printf("Result Server Listening on port %s\n", port)

	r := routers.CreateRouter()

	fs := http.FileServer(http.Dir("./resources"))
	r.PathPrefix("/resources/").Handler(http.StripPrefix("/resources/", fs))

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

	stores.SessionInit()

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
