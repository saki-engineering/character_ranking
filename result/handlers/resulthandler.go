package handlers

import (
	"app/stores"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Vote 投票結果の構造体
type Vote struct {
	Chara       string         `json:"character"`
	User        int            `json:"user"`
	CreatedTime string         `json:"created_at"`
	IP          sql.NullString `json:"ip"`
}

// ResultRootHandler /resultのハンドラ
func ResultRootHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("result/index")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	client := new(http.Client)

	u := &url.URL{}
	u.Scheme = "http"
	u.Host = "vote_api:9090"
	u.Path = "/vote/"
	for i, chara := range charas {
		u.Path = "/vote/" + chara.Name
		uStr := u.String()

		res, e := client.Get(uStr)
		if e != nil {
			log.Println("api request err: ", e)
		}
		defer res.Body.Close()

		b, err2 := ioutil.ReadAll(res.Body)
		if err2 != nil {
			log.Println("http response read err: ", err2)
		}

		var data []Vote
		if err3 := json.Unmarshal(b, &data); err3 != nil {
			log.Println("json parse err: ", err3)
		}

		charas[i].Vote = len(data)
	}

	page := new(Page)
	page.Title = "VIew Result!"
	page.Character = charas
	session, e := stores.GetSession(req)
	if e != nil {
		log.Fatal("session cannot get: ", e)
	}

	if userid, ok := session.Values["userid"].(string); ok {
		page.UserID = userid
		page.LogIn = true
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute on ResultRootHandler: ", err)
	}
}
