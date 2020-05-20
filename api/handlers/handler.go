package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"app/models"

	"github.com/gorilla/mux"
)

// RootHandler /のハンドラ
func RootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

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
// テストするためには $curl -X POST -d "character=cinamon" localhost:9090/vote/
func VoteCharaHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	db, e := models.ConnectDB()
	if e != nil {
		log.Fatal("connect DB: ", e)
	}
	defer db.Close()

	err := models.InsertVotes(db, req.Form.Get("character"))
	if err != nil {
		log.Println("insert: ", err)
	} else {
		log.Println("insert success")
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
