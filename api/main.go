package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"app/routers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	port := "9090"
	fmt.Printf("API Server Listening on port %s\n", port)

	db, e := sql.Open("mysql", "root@/sampledb")
	if e != nil {
		log.Fatal("DB: ", e)
	} else {
		log.Println("Connected to mysql.")
	}
	defer db.Close()

	r := routers.CreateRouter()

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
