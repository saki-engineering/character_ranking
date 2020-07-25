package handlers

import (
	"app/apperrors"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// VoteResultHandler /vote/のGETハンドラ
func VoteResultHandler(w http.ResponseWriter, req *http.Request) {
	db, err := connectDB()
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	defer db.Close()

	voteStructsData, err := db.GetAllVoteData()
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	jsonByteString, err := json.Marshal(voteStructsData)
	if err != nil {
		apperrors.JSONFormatFailed.Wrap(err, "fail to create json data")
		apperrors.ErrorHandler(w, req, err)
		return
	}
	w.Write([]byte(string(jsonByteString)))
}

// VoteCharaHandler /vote/のPOSTハンドラ
// テストするためには $curl -X POST -d "character=cinamon&user=1" localhost:9090/vote/
func VoteCharaHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	db, err := connectDB()
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	defer db.Close()

	err = db.InsertVotes(req.Form.Get("character"), req.Form.Get("user"))
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	log.Println("vote insert success")
}

// CharaResultHandler /vote/{name}のGETハンドラ
func CharaResultHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	db, err := connectDB()
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	defer db.Close()

	voteStructsData, err := db.GetCharaVoteData(vars["name"])
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	jsonByteString, err := json.Marshal(voteStructsData)
	if err != nil {
		apperrors.JSONFormatFailed.Wrap(err, "fail to create json data")
		apperrors.ErrorHandler(w, req, err)
		return
	}
	w.Write([]byte(string(jsonByteString)))
}

// VoteSammaryHandler /vote/summaryのGETハンドラ
func VoteSammaryHandler(w http.ResponseWriter, req *http.Request) {
	db, err := connectDB()
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	defer db.Close()

	resultStructsData, err := db.GetResultSummary()
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	jsonByteString, err := json.Marshal(resultStructsData)
	if err != nil {
		apperrors.JSONFormatFailed.Wrap(err, "fail to create json data")
		apperrors.ErrorHandler(w, req, err)
		return
	}
	w.Write([]byte(string(jsonByteString)))
}
