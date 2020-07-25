package handlers

import (
	"app/models"
	"database/sql"
)

type DB interface {
	Close() error
}

type RealDB struct {
	DB *sql.DB
}

func (db RealDB) Close() error {
	return db.DB.Close()
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
