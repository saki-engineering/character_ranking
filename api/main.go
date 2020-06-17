package main

import (
	"fmt"
	"log"
	"net/http"

	"app/apperrors"
	"app/models"
	"app/routers"
)

func main() {
	port := "9090"
	fmt.Printf("API Server Listening on port %s\n", port)

	db, err := models.ConnectDB()
	if err != nil {
		log.Fatal(err.Code, err.Unwrap())
	}

	if err := models.CreateTable(db); err != nil {
		log.Fatal(err.Code, err.Unwrap())
	} else {
		log.Println("success to create votes & users")
	}

	db.Close()

	r := routers.CreateRouter()

	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		err = apperrors.HTTPServerPortListenFailed.Wrap(err, "server cannot listen port")
		log.Fatal(err.Code, err.Unwrap())
	}
}
