package handlers

import (
	"app/apperrors"
	"app/stores"

	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// ResultRootHandler /resultのハンドラ
func ResultRootHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("result/index")
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	client := new(http.Client)
	uStr := apiURLString("/vote/summary")
	res, err := client.Get(uStr)
	if err != nil {
		err = apperrors.VoteAPIRequestError.Wrap(err, "cannot get vote data")
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	resBodyByteString, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = apperrors.VoteAPIResponseReadFailed.Wrap(err, "cannot get vote data")
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	var jsonParsedData []VoteResult
	if err = json.Unmarshal(resBodyByteString, &jsonParsedData); err != nil {
		err = apperrors.VoteAPIResponseReadFailed.Wrap(err, "cannot get vote data")
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	for _, voteResultData := range jsonParsedData {
		charas[voteResultData.ID-1].Vote = voteResultData.Vote
	}

	page := new(Page)
	page.Title = "VIew Result!"
	page.Character = charas

	conn, err := stores.ConnectRedis()
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	sessionID, _ := stores.GetSessionID(req)

	if nowLoginUserID, _ := stores.GetSessionValue(sessionID, "userid", conn); nowLoginUserID != "" {
		page.UserID = nowLoginUserID
		page.LogIn = true
	}

	err = executeTemplate(w, tmpl, page)
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
}

// CharacterResultHandler /result/{name}のハンドラ
func CharacterResultHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("result/detail")
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(req)

	client := new(http.Client)
	uStr := apiURLString("/vote/" + vars["name"])
	res, err := client.Get(uStr)
	if err != nil {
		err = apperrors.VoteAPIRequestError.Wrap(err, "cannot get vote data")
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	resBodyByteString, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = apperrors.VoteAPIResponseReadFailed.Wrap(err, "cannot get vote data")
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	var jsonParsedData []Vote
	if err = json.Unmarshal(resBodyByteString, &jsonParsedData); err != nil {
		err = apperrors.VoteAPIResponseReadFailed.Wrap(err, "cannot get vote data")
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	page := new(Page)
	page.Title = vars["name"]
	page.Vote = jsonParsedData

	err = executeTemplate(w, tmpl, page)
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
}

// UserSummaryHandler /result/userのハンドラ
func UserSummaryHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("result/user")
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	client := new(http.Client)
	uStr := apiURLString("/user/")
	res, err := client.Get(uStr)
	if err != nil {
		err = apperrors.VoteAPIRequestError.Wrap(err, "cannot get vote data")
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	resBodyByteString, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = apperrors.VoteAPIResponseReadFailed.Wrap(err, "cannot get vote data")
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	var jsonParsedData []User
	if err = json.Unmarshal(resBodyByteString, &jsonParsedData); err != nil {
		err = apperrors.VoteAPIResponseReadFailed.Wrap(err, "cannot get vote data")
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	page := new(Page)
	page.Title = "view result"
	page.VoteUser = jsonParsedData

	err = executeTemplate(w, tmpl, page)
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
}

// UserDetailHandler /result/user/{gender}/{agemin}のハンドラ
func UserDetailHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("result/userdetail")
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(req)

	client := new(http.Client)
	uStr := apiURLString("/user/" + vars["gender"] + "/" + vars["agemin"])
	res, err := client.Get(uStr)
	if err != nil {
		err = apperrors.VoteAPIRequestError.Wrap(err, "cannot get vote data")
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	resBodyByteString, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = apperrors.VoteAPIResponseReadFailed.Wrap(err, "cannot get vote data")
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	var jsonParsedData []Vote
	if err = json.Unmarshal(resBodyByteString, &jsonParsedData); err != nil {
		err = apperrors.VoteAPIResponseReadFailed.Wrap(err, "cannot get vote data")
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	page := new(Page)
	page.Title = "view result"
	page.Vote = jsonParsedData

	err = executeTemplate(w, tmpl, page)
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
}
