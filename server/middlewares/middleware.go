package middlewares

import (
	"fmt"
	"net/http"
)

// Logging アクセス時にリクエスト内容をロギング
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()

		fmt.Println("form: ", req.Form)
		fmt.Println("method: ", req.Method)
		fmt.Println("path: ", req.URL.Path)
		fmt.Println("scheme: ", req.URL.Scheme)

		next.ServeHTTP(w, req)
	})
}
