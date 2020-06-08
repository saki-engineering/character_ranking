package models

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Vote 投じられた票の構造体
type Vote struct {
	Chara       string         `json:"character"`
	User        int            `json:"user"`
	CreatedTime string         `json:"created_at"`
	IP          sql.NullString `json:"ip"`
}

// Result キャラクターごとの得票数をまとめた構造体
type Result struct {
	Chara string `json:"name"`
	Vote  int    `json:"vote"`
}

// ConnectDB DBと接続してポインタを返す
func ConnectDB() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "pass"
	dbAddress := "mysql"
	dbName := "sampledb"

	if os.Getenv("DB_ENV") == "production" {
		dbUser = os.Getenv("DB_USER")
		dbPass = os.Getenv("DB_PASS")
		dbAddress = os.Getenv("DB_ADDRESS")
	}

	//db, e := sql.Open("mysql", "root:pass@tcp(mysql:3306)/sampledb")
	db, e := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbAddress+":3306)/"+dbName)
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
		chara       VARCHAR(20) NOT NULL,
		user        INT UNSIGNED NOT NULL,
		created_at  DATETIME NOT NULL,
		ip          VARCHAR(50),
		FOREIGN KEY (user) REFERENCES users (id)
	);`

	_, err2 := db.Exec(createVoteTable)
	if err2 != nil {
		return err2
	}
	return nil
}

// InsertVotes 指定キャラの投票データをDBに追加
func InsertVotes(db *sql.DB, chara, user string) error {
	const sqlStr = `INSERT INTO votes(chara, user, created_at) VALUES (?, ?, cast(now() as datetime));`

	_, err := db.Exec(sqlStr, chara, user)
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
		err := rows.Scan(&data.Chara, &data.User, &data.CreatedTime, &data.IP)
		if err != nil {
			return nil, err
		}
		dataArray = append(dataArray, data)
	}
	return dataArray, nil
}

// GetCharaVoteData 指定キャラクターの投票データを取得
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
		err := rows.Scan(&data.Chara, &data.User, &data.CreatedTime, &data.IP)
		if err != nil {
			return nil, err
		}
		dataArray = append(dataArray, data)
	}
	return dataArray, nil
}

// GetResultSummary 各キャラとその得票数のデータを取得
func GetResultSummary(db *sql.DB) ([]Result, error) {
	const sqlStr = `SELECT chara, count(*) FROM votes group by chara;`

	rows, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dataArray := make([]Result, 0)
	for rows.Next() {
		var data Result
		err := rows.Scan(&data.Chara, &data.Vote)
		if err != nil {
			return nil, err
		}
		dataArray = append(dataArray, data)
	}
	return dataArray, nil
}

// InsertUsers 投票に参加したユーザーのデータをDBに追加
func InsertUsers(db *sql.DB, age, gender, address string) (int64, error) {
	const sqlStr = `INSERT INTO users(age, gender, address) VALUES (?, ?, ?);`

	result, err := db.Exec(sqlStr, age, gender, address)
	if err != nil {
		return 0, err
	}

	id, err2 := result.LastInsertId()
	if err2 != nil {
		return 0, err2
	}
	return id, nil
}
