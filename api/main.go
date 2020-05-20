package main

import (
	"fmt"
	"log"
	"net/http"

	"app/models"
	"app/routers"
)

func main() {
	port := "9090"
	fmt.Printf("API Server Listening on port %s\n", port)

	db := models.ConnectDB()
	defer db.Close()

	models.CreateTable(db)

	r := routers.CreateRouter()

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
