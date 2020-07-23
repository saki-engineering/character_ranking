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
