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

	db, e := sql.Open("mysql", "root:pass@tcp(mysql:3306)/sampledb")
	if e != nil {
		log.Fatal("DB: ", e)
	} else {
		log.Println("Connected to mysql.")
	}
	defer db.Close()

	const createTable = `CREATE TABLE IF NOT EXISTS votes(
			id        INTEGER PRIMARY KEY,
			chara     VARCHAR(20) NOT NULL,
			user      VARCHAR(100),
			time      DATETIME,
			ip        VARCHAR(50)
		);`
	_, err := db.Exec(createTable)
	if err != nil {
		log.Fatal("createTable: ", err)
	}

	r := routers.CreateRouter()

	err2 := http.ListenAndServe(":"+port, r)
	if err2 != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
