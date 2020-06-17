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
		apperrors.ErrorHandler(w, req, err)
		return
	}
	defer conn.Close()
	sessionID, _ := stores.GetSessionID(req)

	nowVoting, _ := stores.GetSessionValue(sessionID, "voting", conn)
	if nowVoting != "true" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tmpl, err := loadTemplate("form")
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	ageList := make([]int, 99)
	for i := 0; i < 99; i++ {
		ageList[i] = i + 1
	}

	page := new(Page)
	page.Title = "form"
	page.Prefecture = prefecture
	page.Age = ageList
	err = executeTemplate(w, tmpl, page)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
}

// FormVoteHandler アンケートフォームから投票した時のハンドラ
func FormVoteHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	userAge := req.Form.Get("age")
	userGender := req.Form.Get("gender")
	userAddress := req.Form.Get("address")

	//ここにuser登録処理
	client := new(http.Client)
	uStr := apiURLString("/user/")

	values := url.Values{}
	values.Add("age", userAge)
	values.Add("gender", userGender)
	values.Add("address", userAddress)

	res, err := client.Post(uStr, "application/x-www-form-urlencoded", strings.NewReader(values.Encode()))
	if err != nil {
		apperrors.VoteAPIRequestError.Wrap(err, "fail to vote")
		apperrors.ErrorHandler(w, req, err)
		return
	}
	defer res.Body.Close()

	resBodyByteString, err := ioutil.ReadAll(res.Body)
	userID := string(resBodyByteString)

	// ユーザー情報をsessionに付与
	conn, err := stores.ConnectRedis()
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	defer conn.Close()
	sessionID, _ := stores.GetSessionID(req)

	stores.SetSessionValue(sessionID, "user", userID, conn)
	stores.DeleteSessionValue(sessionID, "voting", conn)

	//投票処理が入る
	uStr = apiURLString("/vote/")

	values = url.Values{}
	values.Add("character", req.Form.Get("character"))
	values.Add("user", userID)

	_, err = client.Post(uStr, "application/x-www-form-urlencoded", strings.NewReader(values.Encode()))
	if err != nil {
		apperrors.VoteAPIRequestError.Wrap(err, "fail to vote")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	redirectURL := "/characters/" + req.Form.Get("character") + "/voted"
	http.Redirect(w, req, redirectURL, http.StatusSeeOther)
}
