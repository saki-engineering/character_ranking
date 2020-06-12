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
		apperrors.ErrorHandler(err)
		log.Fatal(apperrors.GetType(err), "||", apperrors.GetMessage(err), "||", err)
	}

	if err := models.CreateTable(db); err != nil {
		apperrors.ErrorHandler(err)
		log.Fatal(apperrors.GetType(err), "||", apperrors.GetMessage(err), "||", err)
	} else {
		log.Println("success to create adminuser table")
	}

	user, err := models.GetUserData(db, "admin")
	if err != nil {
		apperrors.ErrorHandler(err)
		log.Fatal(apperrors.GetType(err), "||", apperrors.GetMessage(err), "||", err)
	}
	if user.UserID != "admin" {
		err = models.UserCreate(db, "admin", "admin", 1)
		if err != nil {
			apperrors.ErrorHandler(err)
			log.Fatal(apperrors.GetType(err), "||", apperrors.GetMessage(err), "||", err)
		} else {
			log.Println("success to create super admin user")
		}
	}

	db.Close()

	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		err = apperrors.HTTPServerPortListenFailed.Wrap(err, "server cannot listen port")
		apperrors.ErrorHandler(err)
		log.Fatal(apperrors.GetType(err), "||", apperrors.GetMessage(err), "||", err)
	}
}
