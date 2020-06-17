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
	port := "7070"
	fmt.Printf("Result Server Listening on port %s\n", port)

	r := routers.CreateRouter()

	db, err := models.ConnectDB()
	if err != nil {
		log.Fatal("fatal err: ", err)
	}

	if err := models.CreateTable(db); err != nil {
		log.Fatal("fatal err: ", err)
	} else {
		log.Println("success to create adminuser table")
	}

	user, err := models.GetUserData(db, "admin")
	if err != nil {
		log.Fatal("fatal err: ", err)
	}
	if user.UserID != "admin" {
		err = models.UserCreate(db, "admin", "admin", 1)
		if err != nil {
			log.Fatal("fatal err: ", err)
		} else {
			log.Println("success to create super admin user")
		}
	}

	db.Close()

	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		err = apperrors.HTTPServerPortListenFailed.Wrap(err, "server cannot listen port")
		log.Fatal("fatal err: ", err)
	}
}
