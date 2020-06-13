package handlers

import (
	"app/apperrors"
	"app/models"
	"encoding/json"

	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateUserHandler /user/のPOSTハンドラ
// テストするためには $curl -X POST -d "age=1&gender=1&address=1" localhost:9090/user/
func CreateUserHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	db, err := models.ConnectDB()
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	id, err := models.InsertUsers(db, req.Form.Get("age"), req.Form.Get("gender"), req.Form.Get("address"))
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
	log.Println("user insert success, id:", id)

	printid := strconv.FormatInt(id, 10)
	fmt.Fprintf(w, printid)
}

// UserResultHandler /user/{gender}/{agemin}のGETハンドラ
func UserResultHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	db, err := models.ConnectDB()
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	gender, _ := strconv.Atoi(vars["gender"])
	agemin, _ := strconv.Atoi(vars["agemin"])

	data, err := models.GetUserSummary(db, gender, agemin)
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		apperrors.JSONFormatFailed.Wrap(err, "fail to create json data")
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
	}
	w.Write([]byte(string(bytes)))
}
