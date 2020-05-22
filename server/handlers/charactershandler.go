package handlers

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

// CharacterHandler キャラクターページのハンドラ
func CharacterHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("characters/index")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	page := Page{"Characters!", charas}
	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute on viewHandler: ", err)
	}
}

// CharacterDetailHandler /[name]のハンドラ
func CharacterDetailHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("characters/detail")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	vars := mux.Vars(req)
	page := Page{vars["name"], charas}
	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute on viewHandler: ", err)
	}
}

// CharacterVoteHandler 投票ボタンが押された時に、フォームに行くかVoted画面に行くかを判定する
func CharacterVoteHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	if req.Form["isAnswered"][0] == "false" {
		http.Redirect(w, req, "/form", 302)
	} else {
		//投票処理
		client := new(http.Client)

		u := &url.URL{}
		u.Scheme = "http"
		u.Host = "vote_api:9090"
		u.Path = "/vote/"
		uStr := u.String()

		values := url.Values{}
		values.Add("character", req.Form.Get("character"))
		// user情報を付加する

		_, err := client.Post(uStr, "application/x-www-form-urlencoded", strings.NewReader(values.Encode()))
		if err != nil {
			log.Println("client post err: ", err)
			return
		}

		url := "/characters/" + req.Form["character"][0] + "/voted"
		http.Redirect(w, req, url, 302)
	}
}

// CharacterVotedHandler 投票終了後の画面を表示
func CharacterVotedHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("characters/voted")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	page := Page{"form", charas}
	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute on viewHandler: ", err)
	}
}
