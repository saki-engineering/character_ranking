package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"./routers"
)

// /にgetしたとき
func Test_RootServe(t *testing.T) {
	ts := httptest.NewServer(routers.CreateRouter())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("invalid StatusCode: %v", res)
	}
}

// /aboutにgetしたとき
func Test_AboutServe(t *testing.T) {
	ts := httptest.NewServer(routers.CreateRouter())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/about")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("invalid StatusCode: %v", res)
	}
}

// /faqにgetしたとき
func Test_FaqServe(t *testing.T) {
	ts := httptest.NewServer(routers.CreateRouter())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/faq")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("invalid StatusCode: %v", res)
	}
}
