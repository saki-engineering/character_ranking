package handlers

import (
	"app/apperrors"
	"app/models"
	"encoding/json"

	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// UserSummaryHandler /user/のGETハンドラ
func UserSummaryHandler(w http.ResponseWriter, req *http.Request) {
	db, err := models.ConnectDB()
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	defer db.Close()

	userStructsData, err := models.GetUserData(db)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	jsonByteString, err := json.Marshal(userStructsData)
	if err != nil {
		apperrors.JSONFormatFailed.Wrap(err, "fail to create json data")
		apperrors.ErrorHandler(w, req, err)
		return
	}
	w.Write([]byte(string(jsonByteString)))
}

// CreateUserHandler /user/のPOSTハンドラ
// テストするためには $curl -X POST -d "age=1&gender=1&address=1" localhost:9090/user/
func CreateUserHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	db, err := models.ConnectDB()
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	defer db.Close()

	insertedUserID, err := models.InsertUsers(db, req.Form.Get("age"), req.Form.Get("gender"), req.Form.Get("address"))
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	idString := strconv.FormatInt(insertedUserID, 10)
	fmt.Fprintf(w, idString)
}

// UserResultHandler /user/{gender}/{agemin}のGETハンドラ
func UserResultHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	db, err := models.ConnectDB()
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	defer db.Close()

	gender, _ := strconv.Atoi(vars["gender"])
	agemin, _ := strconv.Atoi(vars["agemin"])

	voteStructData, err := models.GetUserSummary(db, gender, agemin)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	jsonByteString, err := json.Marshal(voteStructData)
	if err != nil {
		apperrors.JSONFormatFailed.Wrap(err, "fail to create json data")
		apperrors.ErrorHandler(w, req, err)
		return
	}
	w.Write([]byte(string(jsonByteString)))
}
