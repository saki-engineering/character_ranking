package handlers

import (
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

	//投票処理が入る
	client := new(http.Client)

	u := &url.URL{}
	u.Scheme = "http"
	u.Host = "vote_api:9090"
	u.Path = "/vote/"
	uStr := u.String()

	values := url.Values{}
	values.Add("character", req.Form.Get("character"))

	_, err := client.Post(uStr, "application/x-www-form-urlencoded", strings.NewReader(values.Encode()))
	if err != nil {
		log.Println("client post err: ", err)
		return
	}

	url := "/characters/" + req.Form["character"][0] + "/voted"
	http.Redirect(w, req, url, 302)
}
