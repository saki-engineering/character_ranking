package handlers

import (
	"net/url"
	"os"
	"text/template"
)

func loadTemplate(name string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(
		"templates/"+name+".html",
		"templates/partials/_header.html",
		"templates/partials/_footer.html",
	)
	return tmpl, err
}

func apiURLString(path string) string {
	url := &url.URL{}
	url.Scheme = "http"
	url.Host = "vote_api:9090"
	if apiURL := os.Getenv("API_URL"); apiURL != "" {
		url.Host = apiURL + ":9090"
	}
	url.Path = path
	urlStr := url.String()

	return urlStr
}
