package handlers

import (
	"app/models"
)

type DB interface {
	Begin()
	Close() error
	InsertVotes(string, string) error
	GetAllVoteData() ([]models.Vote, error)
	GetCharaVoteData(string) ([]models.Vote, error)
	GetResultSummary() ([]models.Result, error)
	GetUserSummary(int, int) ([]models.Vote, error)
	GetUserData() ([]models.User, error)
	InsertUsers(string, string, string) (int64, error)
}

type mockDB struct{}

func (db mockDB) Close() error {
	return nil
}

func connectDB() (DB, error) {
	// ReadDBの方
	return models.ConnectDB()

	// mockDBの方
	//db := mockDB{}
	//return db, nil
}
