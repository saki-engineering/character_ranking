package handlers

import (
	"net/http"
	"net/url"
	"strings"

	"app/apperrors"
	"app/stores"

	"github.com/gorilla/mux"
)

// CharacterHandler キャラクターページのハンドラ
func CharacterHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("characters/index")
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	page := new(Page)
	page.Title = "Characters!"
	page.Character = charas
	err = executeTemplate(w, tmpl, page)
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
}

// CharacterDetailHandler /[name]のハンドラ
func CharacterDetailHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("characters/detail")
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(req)
	page := new(Page)
	page.Title = vars["name"]
	page.Description = desp[vars["name"]]
	err = executeTemplate(w, tmpl, page)
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
}

// CharacterVoteHandler 投票ボタンが押された時に、フォームに行くかVoted画面に行くかを判定する
func CharacterVoteHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	conn, err := stores.ConnectRedis()
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	sessionID, _ := stores.GetSessionID(req)

	user, err := stores.GetSessionValue(sessionID, "user", conn)
	if err != nil {
		stores.SetSessionValue(sessionID, "voting", "true", conn)
		http.Redirect(w, req, "/form", http.StatusSeeOther)
		return
	}

	//投票処理
	client := new(http.Client)

	uStr := apiURLString("/vote/")

	values := url.Values{}
	values.Add("character", req.Form.Get("character"))
	values.Add("user", user)

	_, err = client.Post(uStr, "application/x-www-form-urlencoded", strings.NewReader(values.Encode()))
	if err != nil {
		apperrors.VoteAPIRequestError.Wrap(err, "fail to vote")
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	url := "/characters/" + req.Form.Get("character") + "/voted"
	http.Redirect(w, req, url, http.StatusSeeOther)
}

// CharacterVotedHandler 投票終了後の画面を表示
func CharacterVotedHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("characters/voted")
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	page := new(Page)
	page.Title = "Completed!"
	err = executeTemplate(w, tmpl, page)
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
}
