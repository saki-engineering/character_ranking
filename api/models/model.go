package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDB DBと接続してポインタを返す
func ConnectDB() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "pass"
	dbName := "sampledb"

	//db, e := sql.Open("mysql", "root:pass@tcp(mysql:3306)/sampledb")
	db, e := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(mysql:3306)/"+dbName)
	if e != nil {
		log.Fatal("DB: ", e)
	} else {
		log.Println("Connected to mysql.")
	}
	return db
}

// CreateTable 投票結果を入れるテーブルがなければ作る
func CreateTable(db *sql.DB) {
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
	} else {
		log.Println("createTable: success to create votes")
	}
}
