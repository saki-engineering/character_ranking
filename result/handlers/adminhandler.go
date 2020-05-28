package handlers

import (
	"app/models"
	"app/stores"
	"log"
	"net/http"
)

// AdminRootHandler /adminのハンドラ
func AdminRootHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("admin/index")
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

// CreateUserFormHandler /admin/userformのGETハンドラ
func CreateUserFormHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("admin/form")
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

// CreateUserHandler /admin/userformのPOSTハンドラ
func CreateUserHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	session, err2 := stores.GetSession(req)
	if err2 != nil {
		log.Fatal("session cannot get: ", err2)
	}
	session.Values["newuserid"] = req.Form.Get("userid")
	session.Values["newpassword"] = req.Form.Get("password")
	session.Save(req, w)

	db, e := models.ConnectDB()
	if e != nil {
		log.Fatal("connect DB: ", e)
	}
	defer db.Close()

	auth := 0
	if req.Form.Get("admin") == "on" {
		auth = 1
	}

	err := models.UserCreate(db, req.Form.Get("userid"), req.Form.Get("password"), auth)
	if err != nil {
		log.Println("create admin user: ", err)
	} else {
		log.Println("success to create admin user")
	}
	http.Redirect(w, req, "/admin/newuser", http.StatusSeeOther)
}

// CheckUserHandler /admin/newuserのハンドラ
func CheckUserHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("admin/newuser")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	session, err2 := stores.GetSession(req)
	if err2 != nil {
		log.Fatal("session cannot get: ", err2)
	}

	page := new(Page)
	page.Title = "View Result!"
	page.NewUser.UserID, _ = session.Values["newuserid"].(string)
	page.NewUser.Password, _ = session.Values["newpassword"].(string)

	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute on RootHandler: ", err)
	}

	delete(session.Values, "newuserid")
	delete(session.Values, "newpassword")
	session.Save(req, w)
}
