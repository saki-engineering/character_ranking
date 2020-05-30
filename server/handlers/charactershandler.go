package handlers

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"app/stores"

	"github.com/gorilla/mux"
)

// CharacterHandler キャラクターページのハンドラ
func CharacterHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("characters/index")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	page := new(Page)
	page.Title = "Characters!"
	page.Character = charas
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
	page := new(Page)
	page.Title = vars["name"]
	page.Description = desp[vars["name"]]
	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute on viewHandler: ", err)
	}
}

// CharacterVoteHandler 投票ボタンが押された時に、フォームに行くかVoted画面に行くかを判定する
func CharacterVoteHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	session, e := stores.GetSession(req)
	if e != nil {
		log.Fatal("session cannot get: ", e)
	}

	user, _ := session.Values["user"].(string)
	if user == "" {
		session.Values["voting"] = true
		session.Save(req, w)

		http.Redirect(w, req, "/form", 302)
		return
	}

	//投票処理
	client := new(http.Client)

	u := &url.URL{}
	u.Scheme = "http"
	u.Host = "vote_api:9090"
	if apiURL := os.Getenv("API_URL"); apiURL != "" {
		u.Host = apiURL + ":9090"
	}
	u.Path = "/vote/"
	uStr := u.String()

	values := url.Values{}
	values.Add("character", req.Form.Get("character"))
	values.Add("user", user)

	_, err1 := client.Post(uStr, "application/x-www-form-urlencoded", strings.NewReader(values.Encode()))
	if err1 != nil {
		log.Println("client post err: ", err1)
		return
	}

	url := "/characters/" + req.Form.Get("character") + "/voted"
	http.Redirect(w, req, url, 302)
}

// CharacterVotedHandler 投票終了後の画面を表示
func CharacterVotedHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("characters/voted")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	page := new(Page)
	page.Title = "Completed!"
	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute on viewHandler: ", err)
	}
}
