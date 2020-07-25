package models

import (
	"database/sql"
	"os"

	"app/apperrors"

	_ "github.com/go-sql-driver/mysql"
)

// Vote 投じられた票の構造体
type Vote struct {
	Chara       string         `json:"character"`
	User        int            `json:"user"`
	Age         int            `json:"age"`
	Gender      int            `json:"gender"`
	Address     int            `json:"address"`
	CreatedTime string         `json:"created_at"`
	IP          sql.NullString `json:"ip"`
}

// Result キャラクターごとの得票数をまとめた構造体
type Result struct {
	CharaID string `json:"id"`
	Vote    int    `json:"vote"`
}

// User 投票に参加した人の構造体
type User struct {
	Num    int `json:"number"`
	Age    int `json:"age"`
	Gender int `json:"gender"`
}

// エントリーNoとキャラ名をセットにした構造体
type chara struct {
	ID   int    // エントリーNo.
	Name string // キャラ名
}

type RealDB struct {
	DB *sql.DB
}

func (db RealDB) Begin() {
	db.DB.Begin()
}

func (db RealDB) Close() error {
	return db.DB.Close()
}

// ConnectDB DBと接続してポインタを返す
func ConnectDB() (RealDB, error) {
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
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbAddress+":3306)/"+dbName)
	if err != nil {
		apperrors.DBConnectionFailed.Wrap(err, "cannot connect to DB")
		return RealDB{nil}, err
	}
	return RealDB{db}, nil
}

func countCharasTableRow(db *sql.DB) (int64, error) {
	const sqlStr = `SELECT COUNT(*) FROM charas;`
	var cnt int64

	rows, err := db.Query(sqlStr)
	if err != nil {
		err = apperrors.MySQLQueryError.Wrap(err, "cannot get data from DB")
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&cnt)
		if err != nil {
			err = apperrors.MySQLDataFormatFailed.Wrap(err, "cannot get data from DB")
			return 0, err
		}
	}
	return cnt, nil
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
		apperrors.MySQLSetUpError.Wrap(err, "fail to set up DB")
		return err
	}

	const createVoteTable = `CREATE TABLE IF NOT EXISTS votes(
		chara       VARCHAR(20) NOT NULL,
		user        INT UNSIGNED NOT NULL,
		created_at  DATETIME NOT NULL,
		ip          VARCHAR(50),
		FOREIGN KEY (user) REFERENCES users (id)
	);`

	_, err = db.Exec(createVoteTable)
	if err != nil {
		apperrors.MySQLSetUpError.Wrap(err, "fail to set up DB")
		return err
	}

	const createCharaTable = `CREATE TABLE IF NOT EXISTS charas(
		id          INT UNSIGNED NOT NULL,
		chara       VARCHAR(20) NOT NULL
	);`

	_, err = db.Exec(createCharaTable)
	if err != nil {
		apperrors.MySQLSetUpError.Wrap(err, "fail to set up DB")
		return err
	}

	cnt, err := countCharasTableRow(db)
	if err != nil {
		apperrors.MySQLSetUpError.Wrap(err, "fail to set up DB")
		return err
	}
	if cnt == 0 {
		const sqlStr = `INSERT INTO charas(id, chara) VALUES (?, ?);`
		for _, chara := range charas {
			_, err = db.Exec(sqlStr, chara.ID, chara.Name)
			if err != nil {
				apperrors.MySQLExecError.Wrap(err, "fail to save data")
				return err
			}
		}
	}

	return nil
}

// InsertVotes 指定キャラの投票データをDBに追加
func (sqldb RealDB) InsertVotes(chara, user string) error {
	db := sqldb.DB
	const sqlStr = `INSERT INTO votes(chara, user, created_at) VALUES (?, ?, cast(now() as datetime));`

	_, err := db.Exec(sqlStr, chara, user)
	if err != nil {
		apperrors.MySQLExecError.Wrap(err, "fail to save data")
		return err
	}
	return nil
}

// GetAllVoteData votesテーブルの全てのデータを取得
func (sqldb RealDB) GetAllVoteData() ([]Vote, error) {
	db := sqldb.DB
	const sqlStr = `SELECT * FROM votes;`

	rows, err := db.Query(sqlStr)
	if err != nil {
		apperrors.MySQLQueryError.Wrap(err, "cannot get data")
		return nil, err
	}
	defer rows.Close()

	voteDataArray := make([]Vote, 0)
	for rows.Next() {
		var voteData Vote
		err := rows.Scan(&voteData.Chara, &voteData.User, &voteData.CreatedTime, &voteData.IP)
		if err != nil {
			apperrors.MySQLDataFormatFailed.Wrap(err, "cannot get data from DB")
			return nil, err
		}
		voteDataArray = append(voteDataArray, voteData)
	}
	return voteDataArray, nil
}

// GetCharaVoteData 指定キャラクターの投票データを取得
func (sqldb RealDB) GetCharaVoteData(chara string) ([]Vote, error) {
	db := sqldb.DB
	const sqlStr = `SELECT users.id, users.age, users.gender, users.address, votes.created_at, votes.ip
					FROM votes LEFT JOIN users ON users.id = votes.user
					WHERE votes.chara = ?;`

	rows, err := db.Query(sqlStr, chara)
	if err != nil {
		apperrors.MySQLQueryError.Wrap(err, "cannot get data")
		return nil, err
	}
	defer rows.Close()

	voteDataArray := make([]Vote, 0)
	for rows.Next() {
		var voteData Vote
		err := rows.Scan(&voteData.User, &voteData.Age, &voteData.Gender, &voteData.Address, &voteData.CreatedTime, &voteData.IP)
		if err != nil {
			apperrors.MySQLDataFormatFailed.Wrap(err, "cannot get data from DB")
			return nil, err
		}
		voteDataArray = append(voteDataArray, voteData)
	}
	return voteDataArray, nil
}

// GetResultSummary 各キャラとその得票数のデータを取得
func (sqldb RealDB) GetResultSummary() ([]Result, error) {
	db := sqldb.DB
	const sqlStr = `SELECT charas.id, count(*)
					FROM charas
					right join votes on charas.chara = votes.chara
					group by charas.id;`

	rows, err := db.Query(sqlStr)
	if err != nil {
		apperrors.MySQLQueryError.Wrap(err, "cannot get data")
		return nil, err
	}
	defer rows.Close()

	resultDataArray := make([]Result, 0)
	for rows.Next() {
		var resultData Result
		err := rows.Scan(&resultData.CharaID, &resultData.Vote)
		if err != nil {
			apperrors.MySQLDataFormatFailed.Wrap(err, "cannot get data from DB")
			return nil, err
		}
		resultDataArray = append(resultDataArray, resultData)
	}
	return resultDataArray, nil
}

// GetUserSummary 性別:gender・年齢:agemin~agemin+9の人たちの投票をみる
func (sqldb RealDB) GetUserSummary(gender, agemin int) ([]Vote, error) {
	db := sqldb.DB
	const sqlStr = `SELECT users.id, users.address, votes.chara, votes.created_at, votes.ip
					FROM votes LEFT JOIN users ON users.id = votes.user
					WHERE users.gender = ? AND users.age BETWEEN ? AND ?;`

	rows, err := db.Query(sqlStr, gender, agemin, agemin+9)
	if err != nil {
		apperrors.MySQLQueryError.Wrap(err, "cannot get data")
		return nil, err
	}
	defer rows.Close()

	voteDataArray := make([]Vote, 0)
	for rows.Next() {
		var voteData Vote
		err := rows.Scan(&voteData.User, &voteData.Address, &voteData.Chara, &voteData.CreatedTime, &voteData.IP)
		if err != nil {
			apperrors.MySQLDataFormatFailed.Wrap(err, "cannot get data from DB")
			return nil, err
		}
		voteDataArray = append(voteDataArray, voteData)
	}
	return voteDataArray, nil
}

// GetUserData 投票に参加した人の一覧データを取得
func (sqldb RealDB) GetUserData() ([]User, error) {
	db := sqldb.DB
	const sqlStr = `SELECT count(*), (case when (age between 0 and 9) then 0
										   when (age between 10 and 19) then 1
										   when (age between 20 and 29) then 2
										   when (age between 30 and 39) then 3
										   when (age between 40 and 49) then 4
										   when (age between 50 and 59) then 5
										   when (age between 60 and 69) then 6
										   when (age between 70 and 79) then 7
										   when (age between 80 and 89) then 8
										   else 9 end) as agelayer, gender
					FROM users GROUP BY agelayer, gender;`

	rows, err := db.Query(sqlStr)
	if err != nil {
		apperrors.MySQLQueryError.Wrap(err, "cannot get data")
		return nil, err
	}
	defer rows.Close()

	userDataArray := make([]User, 0)
	for rows.Next() {
		var userData User
		err := rows.Scan(&userData.Num, &userData.Age, &userData.Gender)
		if err != nil {
			apperrors.MySQLDataFormatFailed.Wrap(err, "cannot get data from DB")
			return nil, err
		}
		userDataArray = append(userDataArray, userData)
	}
	return userDataArray, nil
}

// InsertUsers 投票に参加したユーザーのデータをDBに追加
func (sqldb RealDB) InsertUsers(age, gender, address string) (int64, error) {
	db := sqldb.DB
	const sqlStr = `INSERT INTO users(age, gender, address) VALUES (?, ?, ?);`

	result, err := db.Exec(sqlStr, age, gender, address)
	if err != nil {
		apperrors.MySQLExecError.Wrap(err, "fail to save data")
		return 0, err
	}

	insertedUserID, err := result.LastInsertId()
	if err != nil {
		apperrors.MySQLExecError.Wrap(err, "fail to save data")
		return 0, err
	}
	return insertedUserID, nil
}
