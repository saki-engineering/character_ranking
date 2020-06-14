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

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = apperrors.VoteAPIResponseReadFailed.Wrap(err, "cannot get vote data")
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	var data []VoteResult
	if err = json.Unmarshal(b, &data); err != nil {
		err = apperrors.VoteAPIResponseReadFailed.Wrap(err, "cannot get vote data")
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	for _, votedata := range data {
		charas[votedata.ID-1].Vote = votedata.Vote
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

	if userid, _ := stores.GetSessionValue(sessionID, "userid", conn); userid != "" {
		page.UserID = userid
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

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = apperrors.VoteAPIResponseReadFailed.Wrap(err, "cannot get vote data")
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	var data []Vote
	if err = json.Unmarshal(b, &data); err != nil {
		err = apperrors.VoteAPIResponseReadFailed.Wrap(err, "cannot get vote data")
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	page := new(Page)
	page.Title = vars["name"]
	page.Vote = data

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

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = apperrors.VoteAPIResponseReadFailed.Wrap(err, "cannot get vote data")
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	var data []User
	if err = json.Unmarshal(b, &data); err != nil {
		err = apperrors.VoteAPIResponseReadFailed.Wrap(err, "cannot get vote data")
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	page := new(Page)
	page.Title = "view result"
	page.VoteUser = data

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

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = apperrors.VoteAPIResponseReadFailed.Wrap(err, "cannot get vote data")
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	var data []Vote
	if err = json.Unmarshal(b, &data); err != nil {
		err = apperrors.VoteAPIResponseReadFailed.Wrap(err, "cannot get vote data")
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	page := new(Page)
	page.Title = "view result"
	page.Vote = data

	err = executeTemplate(w, tmpl, page)
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
}
