package models

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCountCharasTableRow(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// 以下、mockDBに期待する動作を定義
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM charas")).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(2))

	db.Begin()
	cnt, err := countCharasTableRow(db)

	if err != nil {
		t.Error(errors.Unwrap(err))
	}
	if cnt != 2 {
		t.Errorf("cnt is wrong: cnt = %d", cnt)
	}
}

func TestInsertVotes(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO votes(chara, user, created_at) VALUES (?, ?, cast(now() as datetime))`)).
		WithArgs("cinnamon", "1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	db.Begin()

	if err := InsertVotes(db, "cinnamon", "1"); err != nil {
		t.Error(errors.Unwrap(err))
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAllVoteData(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dataLen := 3
	charas := []string{"cinnamon", "cinnamon", "cappuccino"}
	user := []int{1, 2, 1}
	createdTime := "2020-07-24 15:18:00"

	rows := sqlmock.NewRows([]string{"chara", "user", "created_at", "ip"})
	for i := 0; i < dataLen; i++ {
		rows.AddRow(charas[i], user[i], createdTime, "")
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM votes")).WillReturnRows(rows)

	db.Begin()
	vote, err := GetAllVoteData(db)
	if err != nil {
		t.Error(err)
		return
	}

	if len(vote) != dataLen {
		t.Errorf("a number of vote is not match dataLen: len(vote) = %d", len(vote))
		return
	}

	flg := true
	for i := 0; i < dataLen; i++ {
		if vote[i].Chara != charas[i] || vote[i].User != user[i] {
			flg = false
			break
		}
	}

	if !flg {
		t.Errorf("vote is wrong %v", vote)
	}
}

func TestGetCharaVoteData(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dataLen := 4
	userID := []int{1, 2, 3, 1}
	userAge := []int{20, 30, 40, 20}
	userGender := []int{1, 2, 9, 1}
	userAdress := []int{12, 44, 32, 12}
	createdTime := "2020-07-24 15:18:00"
	rows := sqlmock.NewRows([]string{"users.id", "users.age", "users.gender", "users.address", "votes.created_at", "votes.ip"})
	for i := 0; i < dataLen; i++ {
		rows.AddRow(userID[i], userAge[i], userGender[i], userAdress[i], createdTime, "")
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT users.id, users.age, users.gender, users.address, votes.created_at, votes.ip
		FROM votes LEFT JOIN users ON users.id = votes.user
		WHERE votes.chara = ?`)).WithArgs("cinnamon").WillReturnRows(rows)

	db.Begin()
	vote, err := GetCharaVoteData(db, "cinnamon")
	if err != nil {
		t.Error(err)
		return
	}

	if len(vote) != dataLen {
		t.Errorf("a number of vote is not match dataLen: len(vote) = %d", len(vote))
		return
	}

	flg := true
	for i := 0; i < dataLen; i++ {
		if vote[i].User != userID[i] || vote[i].Age != userAge[i] || vote[i].Gender != userGender[i] || vote[i].Address != userAdress[i] {
			flg = false
			break
		}
	}

	if !flg {
		t.Errorf("vote is wrong %v", vote)
	}
}

func TestGetResultSummary(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dataLen := 5
	charaID := []string{"1", "2", "5", "7", "10"}
	count := []int{20, 54, 1, 85, 44}
	rows := sqlmock.NewRows([]string{"charas.id", "count(*)"})
	for i := 0; i < dataLen; i++ {
		rows.AddRow(charaID[i], count[i])
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT charas.id, count(*) FROM charas
	right join votes on charas.chara = votes.chara group by charas.id`)).WillReturnRows(rows)

	db.Begin()
	result, err := GetResultSummary(db)
	if err != nil {
		t.Error(err)
		return
	}

	if len(result) != dataLen {
		t.Errorf("a number of result is not match dataLen: len(result) = %d", len(result))
		return
	}

	flg := true
	for i := 0; i < dataLen; i++ {
		if result[i].CharaID != charaID[i] || result[i].Vote != count[i] {
			flg = false
			break
		}
	}

	if !flg {
		t.Errorf("result is wrong %v", result)
	}
}

func TestGetUserSummary(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dataLen := 4
	userID := []int{1, 2, 3, 1}
	userAdress := []int{12, 44, 32, 12}
	voteChara := []string{"1", "4", "7", "2"}
	createdTime := "2020-07-24 15:18:00"

	rows := sqlmock.NewRows([]string{"users.id", "users.address", "votes.chara", "votes.created_at", "votes.ip"})
	for i := 0; i < dataLen; i++ {
		rows.AddRow(userID[i], userAdress[i], voteChara[i], createdTime, "")
	}

	gender := 1
	age := 10

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT users.id, users.address, votes.chara, votes.created_at, votes.ip
		FROM votes LEFT JOIN users ON users.id = votes.user
		WHERE users.gender = ? AND users.age BETWEEN ? AND ?`)).
		WithArgs(gender, age, age+9).WillReturnRows(rows)

	db.Begin()
	vote, err := GetUserSummary(db, gender, age)
	if err != nil {
		t.Error(err)
		return
	}

	if len(vote) != dataLen {
		t.Errorf("a number of vote is not match dataLen: len(vote) = %d", len(vote))
		return
	}

	flg := true
	for i := 0; i < dataLen; i++ {
		if vote[i].User != userID[i] || vote[i].Address != userAdress[i] || vote[i].Chara != voteChara[i] {
			flg = false
			break
		}
	}

	if !flg {
		t.Errorf("vote is wrong %v", vote)
	}
}

func TestGetUserData(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dataLen := 5
	count := []int{32, 68, 109, 86, 41}
	userAgeLayer := []int{1, 1, 2, 3, 4}
	userGender := []int{1, 2, 1, 1, 1}

	rows := sqlmock.NewRows([]string{"count(*)", "agelayer", "gender"})
	for i := 0; i < dataLen; i++ {
		rows.AddRow(count[i], userAgeLayer[i], userGender[i])
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*), (case when (age between 0 and 9) then 0
				when (age between 10 and 19) then 1
				when (age between 20 and 29) then 2
				when (age between 30 and 39) then 3
				when (age between 40 and 49) then 4
				when (age between 50 and 59) then 5
				when (age between 60 and 69) then 6
				when (age between 70 and 79) then 7
				when (age between 80 and 89) then 8
				else 9 end) as agelayer, gender
			FROM users GROUP BY agelayer, gender`)).WillReturnRows(rows)

	db.Begin()
	user, err := GetUserData(db)
	if err != nil {
		t.Error(err)
		return
	}

	if len(user) != dataLen {
		t.Errorf("a number of vote is not match dataLen: len(vote) = %d", len(user))
		return
	}

	flg := true
	for i := 0; i < dataLen; i++ {
		if user[i].Num != count[i] || user[i].Age != userAgeLayer[i] || user[i].Gender != userGender[i] {
			flg = false
			break
		}
	}

	if !flg {
		t.Errorf("user is wrong %v", user)
	}
}

func TestInsertUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	age, gender, address := "22", "1", "13"

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO users(age, gender, address) VALUES (?, ?, ?)`)).
		WithArgs(age, gender, address).
		WillReturnResult(sqlmock.NewResult(1, 1))

	db.Begin()

	userID, err := InsertUsers(db, age, gender, address)
	if userID != 1 {
		t.Errorf("returnValue is not 1: userID = %d", userID)
	}
	if err != nil {
		t.Error(errors.Unwrap(err))
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
