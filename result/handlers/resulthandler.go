package handlers

import (
	"app/apperrors"
	"app/stores"

	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ResultRootHandler /resultのハンドラ
func ResultRootHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("result/index")
	if err != nil {
		apperrors.ErrorHandler(err)
	}

	client := new(http.Client)
	uStr := apiURLString("/vote/summary")
	res, e := client.Get(uStr)
	if e != nil {
		e = apperrors.VoteAPIRequestError.Wrap(e, "cannot get vote data")
		apperrors.ErrorHandler(e)
	}
	defer res.Body.Close()

	b, err2 := ioutil.ReadAll(res.Body)
	if err2 != nil {
		err2 = apperrors.VoteAPIResponseReadFailed.Wrap(err2, "cannot get vote data")
		apperrors.ErrorHandler(err2)
	}

	var data []VoteResult
	if err3 := json.Unmarshal(b, &data); err3 != nil {
		err3 = apperrors.VoteAPIResponseReadFailed.Wrap(err3, "cannot get vote data")
		apperrors.ErrorHandler(err3)
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
		apperrors.ErrorHandler(e)
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
	}
}
