package middlewares

import (
	"fmt"
	"net/http"

	"app/apperrors"
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

// CheckSessionID セッションIDがなければ付与する
func CheckSessionID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		id, err := stores.GetSessionID(req)

		if err != nil || id == "" {
			stores.SetSessionID(w)
		}

		next.ServeHTTP(w, req)
	})
}

// AuthAdmin ユーザーログインしているかどうかをチェック
func AuthAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		conn, e := stores.ConnectRedis()
		if e != nil {
			apperrors.ErrorHandler(e)
			http.Error(w, apperrors.GetMessage(e), http.StatusInternalServerError)
			return
		}
		defer conn.Close()
		sessionID, _ := stores.GetSessionID(req)

		userid, e3 := stores.GetSessionValue(sessionID, "userid", conn)
		if e3 != nil {
			apperrors.ErrorHandler(e3)
			http.Error(w, apperrors.GetMessage(e3), http.StatusInternalServerError)
			return
		}

		if userid != "" {
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
			apperrors.ErrorHandler(err)
			http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
			return
		}
		defer conn.Close()
		sessionID, _ := stores.GetSessionID(req)

		auth, err2 := stores.GetSessionValue(sessionID, "auth", conn)
		if err2 != nil {
			apperrors.ErrorHandler(err2)
			http.Error(w, apperrors.GetMessage(err2), http.StatusInternalServerError)
			return
		}

		if auth == "true" {
			next.ServeHTTP(w, req)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
