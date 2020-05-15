package middlewares

import (
	"fmt"
	"net/http"
	"strings"
)

// Logging アクセス時にリクエスト内容をロギング
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()

		fmt.Println("form: ", req.Form)
		fmt.Println("path: ", req.URL.Path)
		fmt.Println("scheme: ", req.URL.Scheme)
		fmt.Println(req.Form["url_long"])
		for k, v := range req.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
		}

		next.ServeHTTP(w, req)
	})
}
