package handlers

import (
	"app/models"
	"os"
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

func connectDB() (DB, error) {
	if os.Getenv("DB_ENV") == "test" {
		db := mockDB{}
		return db, nil
	}
	return models.ConnectDB()
}
