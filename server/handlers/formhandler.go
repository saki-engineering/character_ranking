package handlers

import (
	"app/stores"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// FormHandler フォームを表示
func FormHandler(w http.ResponseWriter, req *http.Request) {
	session, e := stores.GetSession(req)
	if e != nil {
		log.Fatal("session cannot get: ", e)
	}

	voting, _ := session.Values["voting"].(bool)
	if !voting {
		http.Redirect(w, req, "/", 302)
		return
	}

	tmpl, err := loadTemplate("form")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	age := make([]int, 99)
	for i := 0; i < 99; i++ {
		age[i] = i + 1
	}

	page := new(Page)
	page.Title = "form"
	page.Prefecture = prefecture
	page.Age = age
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

	uStr1 := apiURLString("/user/")

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

	// ユーザー情報をsessionに付与
	session, e := stores.GetSession(req)
	if e != nil {
		log.Fatal("session cannot get: ", e)
	}
	session.Values["user"] = userID

	delete(session.Values, "voting")
	session.Save(req, w)

	//投票処理が入る
	uStr := apiURLString("/vote/")

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
