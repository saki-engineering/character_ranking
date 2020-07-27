package handlers

import "app/models"

type mockDB struct{}

func (db mockDB) Begin() {
	return
}

func (db mockDB) Close() error {
	return nil
}

func (db mockDB) InsertVotes(chara, user string) error {
	return nil
}

func (db mockDB) GetAllVoteData() ([]models.Vote, error) {
	createdTime := "2020-07-24 15:18:00"
	voteDataArray := []models.Vote{
		models.Vote{Chara: "cinnamon", User: 1, CreatedTime: createdTime},
		models.Vote{Chara: "cinnamon", User: 2, CreatedTime: createdTime},
		models.Vote{Chara: "cappuccino", User: 1, CreatedTime: createdTime},
	}

	return voteDataArray, nil
}

func (db mockDB) GetCharaVoteData(chara string) ([]models.Vote, error) {
	createdTime := "2020-07-24 15:18:00"
	voteDataArray := []models.Vote{
		models.Vote{User: 1, Age: 22, Gender: 1, Address: 13, CreatedTime: createdTime},
		models.Vote{User: 1, Age: 22, Gender: 1, Address: 13, CreatedTime: createdTime},
		models.Vote{User: 2, Age: 18, Gender: 2, Address: 22, CreatedTime: createdTime},
	}

	return voteDataArray, nil
}

func (db mockDB) GetResultSummary() ([]models.Result, error) {
	resultDataArray := []models.Result{
		models.Result{CharaID: "cinnamon", Vote: 300},
		models.Result{CharaID: "cappuccino", Vote: 50},
		models.Result{CharaID: "mocha", Vote: 50},
		models.Result{CharaID: "chiffon", Vote: 50},
		models.Result{CharaID: "espresso", Vote: 30},
		models.Result{CharaID: "milk", Vote: 20},
		models.Result{CharaID: "azuki", Vote: 10},
		models.Result{CharaID: "coco", Vote: 70},
		models.Result{CharaID: "nuts", Vote: 70},
		models.Result{CharaID: "poron", Vote: 5},
		models.Result{CharaID: "corne", Vote: 90},
		models.Result{CharaID: "berry", Vote: 22},
		models.Result{CharaID: "cherry", Vote: 10},
	}

	return resultDataArray, nil
}

func (db mockDB) GetUserSummary(gender, agemin int) ([]models.Vote, error) {
	createdTime := "2020-07-24 15:18:00"
	voteDataArray := []models.Vote{
		models.Vote{User: 1, Address: 12, Chara: "cinnamon", CreatedTime: createdTime},
		models.Vote{User: 1, Address: 12, Chara: "cinnamon", CreatedTime: createdTime},
		models.Vote{User: 2, Address: 8, Chara: "mocha", CreatedTime: createdTime},
		models.Vote{User: 2, Address: 8, Chara: "coco", CreatedTime: createdTime},
		models.Vote{User: 3, Address: 33, Chara: "nuts", CreatedTime: createdTime},
	}

	return voteDataArray, nil
}

func (db mockDB) GetUserData() ([]models.User, error) {
	userDataArray := []models.User{
		models.User{Num: 20, Age: 0, Gender: 1},
		models.User{Num: 100, Age: 1, Gender: 1},
		models.User{Num: 20, Age: 1, Gender: 2},
		models.User{Num: 10, Age: 1, Gender: 9},
	}

	return userDataArray, nil
}

func (db mockDB) InsertUsers(age, gender, address string) (int64, error) {
	var insertedUserID int64
	insertedUserID = 1

	return insertedUserID, nil
}
