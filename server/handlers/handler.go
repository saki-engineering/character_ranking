package handlers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

//Page ... htmlに渡す値をまとめた構造体
type Page struct {
	Title     string
	Character []string
}

var charas = []string{
	"cinnamon",
	"cappuccino",
	"mocha",
	"chiffon",
	"espresso",
	"milk",
	"azuki",
	"coco",
	"nuts",
	"poron",
	"corne",
	"berry",
	"cherry",
}

// ViewTopHandler /のハンドラ
func ViewTopHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("index")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	page := Page{"Character rankinig!", charas}
	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute on viewHandler: ", err)
	}
}

// ViewAboutHandler Aboutページのハンドラ
func ViewAboutHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("about")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	page := Page{"About", charas}
	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute on viewHandler: ", err)
	}
}

// ViewFaqHandler FAQページのハンドラ
func ViewFaqHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("faq")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	page := Page{"FAQ", charas}
	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute on viewHandler: ", err)
	}
}

// ViewCharacterHandler キャラクターページのハンドラ
func ViewCharacterHandler(w http.ResponseWriter, req *http.Request) {
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

// ViewFormHandler フォームを表示
func ViewFormHandler(w http.ResponseWriter, req *http.Request) {
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

// VoteHandler 投票ボタンが押された時に、フォームに行くかVoted画面に行くかを判定する
func VoteHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	//まずはformに飛ばすかどうかの前段階リクエストかどうか
	if len(req.Form["isAnswered"]) > 0 {
		if req.Form["isAnswered"][0] == "false" {
			http.Redirect(w, req, "/form", 302)
		} else {
			http.Redirect(w, req, "/characters/name/voted", 302)
		}
	} else { //formに回答した後のpostならば
		http.Redirect(w, req, "/characters/name/voted", 302)
	}
}

// VotedHandler 投票終了後の画面を表示
func VotedHandler(w http.ResponseWriter, req *http.Request) {
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

func loadTemplate(name string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(
		"templates/"+name+".html",
		"templates/partials/_header.html",
		"templates/partials/_footer.html",
	)
	return tmpl, err
}
