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
		apperrors.ErrorHandler(err)
		log.Fatal(apperrors.GetType(err), "||", apperrors.GetMessage(err), "||", err)
	}

	if err := models.CreateTable(db); err != nil {
		apperrors.ErrorHandler(err)
		log.Fatal(apperrors.GetType(err), "||", apperrors.GetMessage(err), "||", err)
	} else {
		log.Println("success to create votes & users")
	}

	db.Close()

	r := routers.CreateRouter()

	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		err = apperrors.HTTPServerPortListenFailed.Wrap(err, "server cannot listen port")
		apperrors.ErrorHandler(err)
		log.Fatal(apperrors.GetType(err), "||", apperrors.GetMessage(err), "||", err)
	}
}
