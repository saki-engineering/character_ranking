package models

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

// AdminUser 管理者ユーザーの構造体
type AdminUser struct {
	UserID         string
	HashedPassword string
	// Auth 0だと一般管理者、1だと強管理者
	Auth int
}

// ConnectDB DBと接続してポインタを返す
// 接続時にエラーがあった場合は、それが返り値errorに入る
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

// CreateTable 管理者一覧テーブルがなければ作る
// SQL実行時にエラーがあった場合は、返り値errorに入る
func CreateTable(db *sql.DB) error {
	const createUserTable = `CREATE TABLE IF NOT EXISTS adminusers(
		userid           VARCHAR(50) NOT NULL PRIMARY KEY,
		hashedpassword   VARCHAR(500) NOT NULL,
		auth             INT NOT NULL
	);`
	_, err := db.Exec(createUserTable)
	if err != nil {
		return err
	}
	return nil
}

// UserCreate adminuserのデータをDBに保存する
// パスワードのハッシュ化に失敗、もしくはinsert失敗時に返り値errorが返る
func UserCreate(db *sql.DB, userid, password string, auth int) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	const sqlStr = `INSERT INTO adminusers(userid, hashedpassword, auth) VALUES (?, ?, ?);`
	_, err2 := db.Exec(sqlStr, userid, hash, auth)
	if err2 != nil {
		return err2
	}

	return nil
}

// GetUserData 与えられたuseridの管理者データを探す
// SQL実行、もしくは結果のAdminUser構造体へのパースに失敗した場合は、返り値にerrorが入る
// → その場合、AdminUser{UserID: "", HashedPassword: "", Auth: 0}が返る(構造体初期化時の値)
func GetUserData(db *sql.DB, userid string) (AdminUser, error) {
	const sqlStr = `SELECT * FROM adminusers WHERE userid=?;`
	user := AdminUser{}

	rows, err := db.Query(sqlStr, userid)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	for rows.Next() {
		e := rows.Scan(&user.UserID, &user.HashedPassword, &user.Auth)
		if e != nil {
			return user, e
		}
	}

	return user, nil
}

// CheckIDExist 与えられたuseridの管理者データを探す
// SQL実行、もしくはSQL結果のscanに失敗した場合、エラーを返す
// → その場合、intは1を返す
func CheckIDExist(db *sql.DB, userid string) (int, error) {
	const sqlStr = `SELECT COUNT(*) FROM adminusers WHERE userid=?;`
	cnt := 1

	rows, err := db.Query(sqlStr, userid)
	if err != nil {
		return cnt, err
	}
	defer rows.Close()

	for rows.Next() {
		e := rows.Scan(&cnt)
		if e != nil {
			return cnt, e
		}
	}

	return cnt, nil
}
