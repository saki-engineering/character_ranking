package handlers

import (
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

// VoteRootHandler /vote/のハンドラ
func VoteRootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Vote root")
}

// CharaResultHandler /vote/{name}のGETハンドラ
func CharaResultHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	fmt.Fprintf(w, vars["name"])
}

// VoteCharaHandler /vote/{name}のPOSTハンドラ
// テストするためには $curl -X POST -d "character=cinamon" localhost:9090/vote/name
func VoteCharaHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	fmt.Fprintf(w, vars["name"])

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
