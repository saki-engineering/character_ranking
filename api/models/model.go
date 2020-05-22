package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Vote 投票結果の構造体
type Vote struct {
	Chara string         `json:"character"`
	User  sql.NullInt64  `json:"user"`
	Time  sql.NullString `json:"time"`
	IP    sql.NullString `json:"ip"`
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
	const createUserTable = `CREATE TABLE IF NOT EXISTS users(
		id        INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
		age       INT NOT NULL,
		gender    INT NOT NULL,
		address   INT NOT NULL
	);`
	_, err := db.Exec(createUserTable)
	if err != nil {
		return err
	}

	const createVoteTable = `CREATE TABLE IF NOT EXISTS votes(
		chara     VARCHAR(20) NOT NULL,
		user      INT UNSIGNED,
		time      DATETIME,
		ip        VARCHAR(50),
		FOREIGN KEY (user) REFERENCES users (id)
	);`

	_, err2 := db.Exec(createVoteTable)
	if err2 != nil {
		return err2
	}
	return nil
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

// GetAllVoteData votesテーブルの全てのデータを取得
func GetAllVoteData(db *sql.DB) ([]Vote, error) {
	const sqlStr = `SELECT * FROM votes;`

	rows, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dataArray := make([]Vote, 0)
	for rows.Next() {
		var data Vote
		err := rows.Scan(&data.Chara, &data.User, &data.Time, &data.IP)
		if err != nil {
			return nil, err
		}
		dataArray = append(dataArray, data)
	}
	return dataArray, nil
}

// GetCharaVoteData votesテーブルの全てのデータを取得
func GetCharaVoteData(db *sql.DB, chara string) ([]Vote, error) {
	const sqlStr = `SELECT * FROM votes WHERE chara=?;`

	rows, err := db.Query(sqlStr, chara)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dataArray := make([]Vote, 0)
	for rows.Next() {
		var data Vote
		err := rows.Scan(&data.Chara, &data.User, &data.Time, &data.IP)
		if err != nil {
			return nil, err
		}
		dataArray = append(dataArray, data)
	}
	return dataArray, nil
}

// InsertUsers 指定キャラの投票データをDBに追加
func InsertUsers(db *sql.DB, age, gender, address string) error {
	const sqlStr = `INSERT INTO users(age, gender, address) VALUES (?, ?, ?);`

	_, err := db.Exec(sqlStr, age, gender, address)
	if err != nil {
		return err
	}
	return nil
}
