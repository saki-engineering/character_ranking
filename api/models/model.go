package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type vote struct {
	Chara, User, IP string
	Time            string
}

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

// CreateTable 投票結果を入れるテーブルがなければ作る
func CreateTable(db *sql.DB) error {
	const createTable = `CREATE TABLE IF NOT EXISTS votes(
		chara     VARCHAR(20) NOT NULL,
		user      VARCHAR(100),
		time      DATETIME,
		ip        VARCHAR(50)
	);`

	_, err := db.Exec(createTable)
	return err
}

// InsertVotes 指定キャラの投票データをDBに追加
func InsertVotes(db *sql.DB, chara string) error {
	const sqlStr = `INSERT INTO votes(chara) VALUES (?);`

	_, err := db.Exec(sqlStr, chara)
	if err != nil {
		return err
	}
	return nil
}

// GetAllData votesテーブルの全てのデータを取得
/*
func GetAllData(db *sql.DB) {
	const sqlStr = `SELECT * FROM votes;`

	rows, err := db.Query(sqlStr)
	defer rows.Close()

	if err != nil {
		log.Fatal("GetAllData: ", err)
	}

	for rows.Next() {

	}
}
*/
