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

	user, err2 := models.GetUserData(db, "admin")
	if err2 != nil {
		log.Println("fail to search user 'admin'")
	}
	if user.UserID != "admin" {
		err3 := models.UserCreate(db, "admin", "admin", 1)
		if err3 != nil {
			log.Println("fail to create super admin user: ", err3)
		} else {
			log.Println("success to create super admin user")
		}
	}

	db.Close()

	stores.SessionInit()

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
