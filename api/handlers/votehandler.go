package handlers

import (
	"app/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// VoteResultHandler /vote/のGETハンドラ
func VoteResultHandler(w http.ResponseWriter, req *http.Request) {
	db, e := models.ConnectDB()
	if e != nil {
		log.Fatal("connect DB: ", e)
	}
	defer db.Close()

	data, err := models.GetAllVoteData(db)
	if err != nil {
		log.Println("fail GetAllVoteData: ", err)
	}

	bytes, err2 := json.Marshal(data)
	if err2 != nil {
		log.Println("fail json Marshal: ", err2)
	}
	w.Write([]byte(string(bytes)))
}

// VoteCharaHandler /vote/のPOSTハンドラ
// テストするためには $curl -X POST -d "character=cinamon&user=1" localhost:9090/vote/
func VoteCharaHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	db, e := models.ConnectDB()
	if e != nil {
		log.Fatal("connect DB: ", e)
	}
	defer db.Close()

	err := models.InsertVotes(db, req.Form.Get("character"), req.Form.Get("user"))
	if err != nil {
		log.Println("insert: ", err)
	} else {
		log.Println("vote insert success")
	}
}

// CharaResultHandler /vote/{name}のGETハンドラ
func CharaResultHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	db, e := models.ConnectDB()
	if e != nil {
		log.Fatal("connect DB: ", e)
	}
	defer db.Close()

	data, err := models.GetCharaVoteData(db, vars["name"])
	if err != nil {
		log.Println("fail GetCharaVoteData: ", err)
	}

	bytes, err2 := json.Marshal(data)
	if err2 != nil {
		log.Println("fail json Marshal: ", err2)
	}
	w.Write([]byte(string(bytes)))
}

// VoteSammaryHandler /vote/summaryのGETハンドラ
func VoteSammaryHandler(w http.ResponseWriter, req *http.Request) {
	db, e := models.ConnectDB()
	if e != nil {
		log.Fatal("connect DB: ", e)
	}
	defer db.Close()

	data, err := models.GetResultSummary(db)
	if err != nil {
		log.Println("fail GetResultSummary: ", err)
	}

	bytes, err2 := json.Marshal(data)
	if err2 != nil {
		log.Println("fail json Marshal: ", err2)
	}
	w.Write([]byte(string(bytes)))
}
