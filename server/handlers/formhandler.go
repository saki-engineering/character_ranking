package handlers

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"app/apperrors"
	"app/stores"
)

// FormHandler フォームを表示
func FormHandler(w http.ResponseWriter, req *http.Request) {
	conn, err := stores.ConnectRedis()
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	sessionID, _ := stores.GetSessionID(req)

	voting, _ := stores.GetSessionValue(sessionID, "voting", conn)
	if voting != "true" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tmpl, err := loadTemplate("form")
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	age := make([]int, 99)
	for i := 0; i < 99; i++ {
		age[i] = i + 1
	}

	page := new(Page)
	page.Title = "form"
	page.Prefecture = prefecture
	page.Age = age
	err = executeTemplate(w, tmpl, page)
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
}

// FormVoteHandler アンケートフォームから投票した時のハンドラ
func FormVoteHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	//ここにuser登録処理
	client := new(http.Client)

	uStr1 := apiURLString("/user/")

	values1 := url.Values{}
	values1.Add("age", req.Form.Get("age"))
	values1.Add("gender", req.Form.Get("gender"))
	values1.Add("address", req.Form.Get("address"))

	res, err := client.Post(uStr1, "application/x-www-form-urlencoded", strings.NewReader(values1.Encode()))
	if err != nil {
		apperrors.VoteAPIRequestError.Wrap(err, "fail to vote")
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	userID := string(b)

	// ユーザー情報をsessionに付与
	conn, err := stores.ConnectRedis()
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	sessionID, _ := stores.GetSessionID(req)

	stores.SetSessionValue(sessionID, "user", userID, conn)
	stores.DeleteSessionValue(sessionID, "voting", conn)

	//投票処理が入る
	uStr := apiURLString("/vote/")

	values := url.Values{}
	values.Add("character", req.Form.Get("character"))
	values.Add("user", userID)

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
