package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Vote 投票結果の構造体
type Vote struct {
	Chara string         `json:"character"`
	User  sql.NullString `json:"user"`
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
