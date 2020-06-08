package handlers

import (
	"app/stores"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// ResultRootHandler /resultのハンドラ
func ResultRootHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("result/index")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	client := new(http.Client)
	uStr := apiURLString("/vote/summary")
	res, e := client.Get(uStr)
	if e != nil {
		log.Println("api request err: ", e)
	}
	defer res.Body.Close()

	b, err2 := ioutil.ReadAll(res.Body)
	if err2 != nil {
		log.Println("http response read err: ", err2)
	}

	var data []VoteResult
	if err3 := json.Unmarshal(b, &data); err3 != nil {
		log.Println("json parse err: ", err3)
	}

	for _, votedata := range data {
		for i, chara := range charas {
			if chara.Name == votedata.Name {
				charas[i].Vote = votedata.Vote
				break
			}
		}
	}

	page := new(Page)
	page.Title = "VIew Result!"
	page.Character = charas

	conn, e := stores.ConnectRedis()
	if e != nil {
		log.Fatal("cannot connect redis: ", e)
	}
	defer conn.Close()
	sessionID, _ := stores.GetSessionID(req)

	if userid, _ := stores.GetSessionValue(sessionID, "userid", conn); userid != "" {
		page.UserID = userid
		page.LogIn = true
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute on ResultRootHandler: ", err)
	}
}
