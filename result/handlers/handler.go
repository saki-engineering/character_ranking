package handlers

import (
	"log"
	"net/http"
	"text/template"

	"app/models"
	"app/stores"

	"golang.org/x/crypto/bcrypt"
)

//Page ... htmlに渡す値をまとめた構造体
type Page struct {
	Title  string
	UserID string
	LogIn  bool
}

// RootHandler /のハンドラ
func RootHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("index")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	page := Page{"View Result!", "", false}
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
		log.Fatal("Execute on RootHandler: ", err)
	}
}

// LoginPageHandler /loginのGETハンドラ
func LoginPageHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("login")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	page := Page{"View Result!", "", false}
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
	session.Save(req, w)

	log.Println("login success: userid=", user.UserID)

	http.Redirect(w, req, "/", http.StatusSeeOther)
}

// SignupPageHandler /signupのGETハンドラ
func SignupPageHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("signup")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	page := Page{"View Result!", "", false}
	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute on RootHandler: ", err)
	}
}

// SignupHandler /signupのPOSTハンドラ
func SignupHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	db, e := models.ConnectDB()
	if e != nil {
		log.Fatal("connect DB: ", e)
	}
	defer db.Close()

	err := models.UserCreate(db, req.Form.Get("userid"), req.Form.Get("password"))
	if err != nil {
		log.Println("create admin user: ", err)
	} else {
		log.Println("success to create admin user")
	}
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

// LogoutHandler /logoutのハンドラ
func LogoutHandler(w http.ResponseWriter, req *http.Request) {
	session, err := stores.GetSession(req)
	if err != nil {
		log.Fatal("session cannot get: ", err)
	}
	delete(session.Values, "userid")
	session.Save(req, w)

	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func loadTemplate(name string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(
		"templates/"+name+".html",
		"templates/partials/_header.html",
		"templates/partials/_footer.html",
	)
	return tmpl, err
}
