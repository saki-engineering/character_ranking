package handlers

import (
	"net/http"
	"net/url"
	"os"
	"text/template"

	"app/apperrors"
)

func loadTemplate(name string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(
		"templates/"+name+".html",
		"templates/partials/_header.html",
		"templates/partials/_footer.html",
	)

	if err != nil {
		err = apperrors.HTMLTemplateLoadFailed.Wrap(err, "cannot load html")
	}
	return tmpl, err
}

func executeTemplate(w http.ResponseWriter, tmpl *template.Template, page *Page) error {
	err := tmpl.Execute(w, page)

	if err != nil {
		err = apperrors.HTMLTemplateExecFailed.Wrap(err, "cannot load html")
		return err
	}
	return nil
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
