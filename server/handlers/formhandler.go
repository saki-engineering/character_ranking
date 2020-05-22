package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// FormHandler フォームを表示
func FormHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("form")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	page := Page{"form", charas}
	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute on viewHandler: ", err)
	}
}

// FormVoteHandler アンケートフォームから投票した時のハンドラ
func FormVoteHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	//ここにuser登録処理
	client := new(http.Client)

	u1 := &url.URL{}
	u1.Scheme = "http"
	u1.Host = "vote_api:9090"
	u1.Path = "/user/"
	uStr1 := u1.String()

	values1 := url.Values{}
	values1.Add("age", req.Form.Get("age"))
	values1.Add("gender", req.Form.Get("gender"))
	values1.Add("address", req.Form.Get("address"))

	res, err1 := client.Post(uStr1, "application/x-www-form-urlencoded", strings.NewReader(values1.Encode()))
	if err1 != nil {
		log.Println("client post err to user create: ", err1)
		return
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	userID := string(b)

	// ユーザー情報をcookieに付与
	cookie := &http.Cookie{
		Name:  "user",
		Value: userID,
		Path:  "/",
	}
	http.SetCookie(w, cookie)

	//投票処理が入る
	u := &url.URL{}
	u.Scheme = "http"
	u.Host = "vote_api:9090"
	u.Path = "/vote/"
	uStr := u.String()

	values := url.Values{}
	values.Add("character", req.Form.Get("character"))
	values.Add("user", userID)

	_, err = client.Post(uStr, "application/x-www-form-urlencoded", strings.NewReader(values.Encode()))
	if err != nil {
		log.Println("client post err: ", err)
		return
	}

	url := "/characters/" + req.Form.Get("character") + "/voted"
	http.Redirect(w, req, url, 302)
}
