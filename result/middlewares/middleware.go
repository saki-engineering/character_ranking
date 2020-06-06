package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"app/models"
	"app/stores"

	"github.com/google/uuid"
)

// Logging アクセス時にリクエスト内容をロギング
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		fmt.Printf("method: %s || path: %s || form: %s\n", req.Method, req.URL.Path, req.Form)

		next.ServeHTTP(w, req)
	})
}

// CheckSessionID セッションIDがなければ付与する
func CheckSessionID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		id, err := stores.GetSessionID(req)

		if err != nil || id == "" {
			uuid, err2 := uuid.NewRandom()
			if err2 != nil {
				log.Println("cannot make uuid: ", err2)
			}
			cookie := &http.Cookie{
				Name:    stores.SessionName,
				Value:   uuid.String(),
				Expires: time.Now().AddDate(0, 0, 7),
			}
			http.SetCookie(w, cookie)
		}

		next.ServeHTTP(w, req)
	})
}

// AuthAdmin ユーザーログインしているかどうかをチェック
func AuthAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		conn, e := stores.ConnectRedis()
		if e != nil {
			log.Fatal("cannot connect redis: ", e)
		}
		defer conn.Close()
		sessionID, _ := stores.GetSessionID(req)

		db, e2 := models.ConnectDB()
		if e2 != nil {
			log.Fatal("connect DB: ", e2)
		}
		defer db.Close()

		userid, e3 := stores.GetSessionValue(sessionID, "userid", conn)
		if e3 != nil {
			log.Println("cannot get session key userid: ", e3)
		}

		user, _ := models.GetUserData(db, userid)
		if user.UserID != "" {
			next.ServeHTTP(w, req)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

// AuthSuperAdmin 管理者ユーザーログインしているかどうかをチェック
func AuthSuperAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		conn, err := stores.ConnectRedis()
		if err != nil {
			log.Fatal("cannot connect redis: ", err)
		}
		defer conn.Close()
		sessionID, _ := stores.GetSessionID(req)

		db, e := models.ConnectDB()
		if e != nil {
			log.Fatal("connect DB: ", e)
		}
		defer db.Close()

		auth, err2 := stores.GetSessionValue(sessionID, "auth", conn)
		if err2 != nil {
			log.Println("cannot get session key userid: ", err2)
		}

		if auth == "true" {
			next.ServeHTTP(w, req)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
