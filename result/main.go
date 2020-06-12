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

	db, e := models.ConnectDB()
	if e != nil {
		log.Fatal(apperrors.GetType(e), "||", apperrors.GetMessage(e), "||", e)
	}

	if err := models.CreateTable(db); err != nil {
		apperrors.ErrorHandler(err)
		log.Fatal(apperrors.GetType(err), "||", apperrors.GetMessage(err), "||", err)
	} else {
		log.Println("success to create adminuser table")
	}

	user, err2 := models.GetUserData(db, "admin")
	if err2 != nil {
		log.Fatal(apperrors.GetType(err2), "||", apperrors.GetMessage(err2), "||", err2)
	}
	if user.UserID != "admin" {
		err3 := models.UserCreate(db, "admin", "admin", 1)
		if err3 != nil {
			log.Fatal(apperrors.GetType(err3), "||", apperrors.GetMessage(err3), "||", err3)
		} else {
			log.Println("success to create super admin user")
		}
	}

	db.Close()

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		err = apperrors.HTTPServerPortListenFailed.Wrap(err, "server cannot listen port")
		log.Fatal(apperrors.GetType(err), "||", apperrors.GetMessage(err), "||", err)
	}
}
