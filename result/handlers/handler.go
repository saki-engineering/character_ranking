package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"app/models"
	"app/stores"

	"golang.org/x/crypto/bcrypt"
)

// RootHandler /のハンドラ
func RootHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("index")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	page := new(Page)
	page.Title = "View Result!"
	page.UserID = ""
	page.LogIn = false
	page.Admin = false
	session, e := stores.GetSession(req)
	if e != nil {
		log.Fatal("session cannot get: ", e)
	}

	if userid, ok := session.Values["userid"].(string); ok {
		page.UserID = userid
		page.LogIn = true
	}
	if auth, ok2 := session.Values["auth"].(bool); ok2 {
		page.Admin = auth
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute on RootHandler: ", err)
	}
}

// LoginPageHandler /loginのGETハンドラ
func LoginPageHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("login")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	page := new(Page)
	page.Title = "View Result!"
	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute on RootHandler: ", err)
	}
}

// LoginHandler /loginのPOSTハンドラ
func LoginHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	db, e := models.ConnectDB()
	if e != nil {
		log.Fatal("connect DB: ", e)
	}
	defer db.Close()

	user, err := models.GetUserData(db, req.Form.Get("userid"))
	if err != nil {
		log.Println("cannot get adminuser data: ", err)
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	err2 := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.Form.Get("password")))
	if err2 != nil {
		log.Println("password is not correct: ", err2)
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	session, err3 := stores.GetSession(req)
	if err3 != nil {
		log.Fatal("session cannot get: ", err3)
	}
	session.Values["userid"] = user.UserID
	if user.Auth == 1 {
		session.Values["auth"] = true
	} else {
		session.Values["auth"] = false
	}
	session.Save(req, w)

	log.Println("login success: userid=", user.UserID)

	http.Redirect(w, req, "/", http.StatusSeeOther)
}

// LogoutHandler /logoutのハンドラ
func LogoutHandler(w http.ResponseWriter, req *http.Request) {
	session, err := stores.GetSession(req)
	if err != nil {
		log.Fatal("session cannot get: ", err)
	}
	delete(session.Values, "userid")
	delete(session.Values, "auth")
	session.Save(req, w)

	http.Redirect(w, req, "/", http.StatusSeeOther)
}

// CheckIDHandler /checkidのPOSTハンドラ
func CheckIDHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	db, e := models.ConnectDB()
	if e != nil {
		log.Fatal("connect DB: ", e)
	}
	defer db.Close()

	cnt, err := models.CheckIDExist(db, req.Form.Get("userid"))
	if err != nil {
		log.Println("checkIDExxist: ", err)
	}

	printnum := strconv.FormatInt(int64(cnt), 10)

	fmt.Fprintf(w, printnum)
}

func loadTemplate(name string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(
		"templates/"+name+".html",
		"templates/partials/_header.html",
		"templates/partials/_footer.html",
	)
	return tmpl, err
}
