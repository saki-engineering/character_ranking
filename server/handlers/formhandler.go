package handlers

import (
	"log"
	"net/http"
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
	//ここにuser登録処理と、投票処理が入る

	req.ParseForm()
	url := "/characters/" + req.Form["character"][0] + "/voted"
	http.Redirect(w, req, url, 302)
}
