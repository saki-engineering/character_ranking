package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDB DBと接続してポインタを返す
func ConnectDB() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "pass"
	dbName := "sampledb"

	//db, e := sql.Open("mysql", "root:pass@tcp(mysql:3306)/sampledb")
	db, e := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(mysql:3306)/"+dbName)
	return db, e
}

// CreateTable 管理者一覧テーブルがなければ作る
func CreateTable(db *sql.DB) error {
	const createUserTable = `CREATE TABLE IF NOT EXISTS adminusers(
		userid           VARCHAR(50) NOT NULL PRIMARY KEY,
		hashedpassword   VARCHAR(50) NOT NULL,
		auth             INT NOT NULL
	);`
	_, err := db.Exec(createUserTable)
	if err != nil {
		return err
	}
	return nil
}
