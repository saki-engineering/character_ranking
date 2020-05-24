package middlewares

import (
	"fmt"
	"log"
	"net/http"

	"app/models"
	"app/stores"
)

// Logging アクセス時にリクエスト内容をロギング
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		fmt.Printf("method: %s || path: %s || form: %s\n", req.Method, req.URL.Path, req.Form)

		next.ServeHTTP(w, req)
	})
}

// AuthAdmin ユーザーログインしているかどうかをチェック
func AuthAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		session, err := stores.GetSession(req)
		if err != nil {
			log.Fatal("session cannot get: ", err)
		}

		db, e := models.ConnectDB()
		if e != nil {
			log.Fatal("connect DB: ", e)
		}
		defer db.Close()

		var userid string
		if id, ok := session.Values["userid"].(string); ok {
			userid = id
		}

		user, err := models.GetUserData(db, userid)
		if user.UserID != "" {
			next.ServeHTTP(w, req)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
