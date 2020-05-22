package main

import (
	"fmt"
	"log"
	"net/http"

	"app/middlewares"
	"app/models"
	"app/routers"
)

func main() {
	port := "9090"
	fmt.Printf("API Server Listening on port %s\n", port)

	db, e := models.ConnectDB()
	if e != nil {
		log.Fatal("connect DB: ", e)
	}

	if err := models.CreateTable(db); err != nil {
		log.Println("create table: ", err)
	} else {
		log.Println("success to create votes & users")
	}

	db.Close()

	r := routers.CreateRouter()
	r.Use(middlewares.Logging)

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
