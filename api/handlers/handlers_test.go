package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

func setDBEnv(env string) {
	os.Setenv("DB_ENV", env)
}

func deleteLineAndTab(s string) string {
	deleteline := strings.ReplaceAll(s, "\n", "")
	deletetab := strings.ReplaceAll(deleteline, "\t", "")

	return deletetab
}

func TestRootHandler(t *testing.T) {
	env := os.Getenv("DB_ENV")
	os.Setenv("DB_ENV", "test")
	defer setDBEnv(env)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	got := httptest.NewRecorder()

	RootHandler(got, req)

	wantBody := "Hello World"
	if got.Code != http.StatusOK {
		t.Errorf("want http.StatusOk, but get %d", got.Code)
	}
	if body := got.Body.String(); body != wantBody {
		t.Errorf("want body %s, but get body %s", wantBody, body)
	}
}

func TestVoteResultHandler(t *testing.T) {
	env := os.Getenv("DB_ENV")
	os.Setenv("DB_ENV", "test")
	defer setDBEnv(env)

	req := httptest.NewRequest(http.MethodGet, "/vote/", nil)
	got := httptest.NewRecorder()

	VoteResultHandler(got, req)

	wantBody := `[
		{"character":"cinnamon","user":1,"age":0,"gender":0,"address":0,"created_at":"2020-07-24 15:18:00","ip":{"String":"","Valid":false}},
		{"character":"cinnamon","user":2,"age":0,"gender":0,"address":0,"created_at":"2020-07-24 15:18:00","ip":{"String":"","Valid":false}},
		{"character":"cappuccino","user":1,"age":0,"gender":0,"address":0,"created_at":"2020-07-24 15:18:00","ip":{"String":"","Valid":false}}
		]`
	if got.Code != http.StatusOK {
		t.Errorf("want http.StatusOk, but get %d", got.Code)
	}
	if body := got.Body.String(); body != deleteLineAndTab(wantBody) {
		t.Errorf("get wrong body %s", body)
	}
}

func TestVoteCharaHandler(t *testing.T) {
	env := os.Getenv("DB_ENV")
	os.Setenv("DB_ENV", "test")
	defer setDBEnv(env)

	values := url.Values{}
	values.Add("character", "cinnamon")
	values.Add("user", "1")

	req := httptest.NewRequest(http.MethodPost, "/vote/", strings.NewReader(values.Encode()))
	got := httptest.NewRecorder()

	VoteCharaHandler(got, req)

	if got.Code != http.StatusOK {
		t.Errorf("want http.StatusOk, but get %d", got.Code)
	}
}

func TestCharaResultHandler(t *testing.T) {
	env := os.Getenv("DB_ENV")
	os.Setenv("DB_ENV", "test")
	defer setDBEnv(env)

	req := httptest.NewRequest(http.MethodGet, "/vote/cinnamon", nil)
	got := httptest.NewRecorder()

	CharaResultHandler(got, req)

	wantBody := `[
		{"character":"","user":1,"age":22,"gender":1,"address":13,"created_at":"2020-07-24 15:18:00","ip":{"String":"","Valid":false}},
		{"character":"","user":1,"age":22,"gender":1,"address":13,"created_at":"2020-07-24 15:18:00","ip":{"String":"","Valid":false}},
		{"character":"","user":2,"age":18,"gender":2,"address":22,"created_at":"2020-07-24 15:18:00","ip":{"String":"","Valid":false}}
		]`
	if got.Code != http.StatusOK {
		t.Errorf("want http.StatusOk, but get %d", got.Code)
	}
	if body := got.Body.String(); body != deleteLineAndTab(wantBody) {
		t.Errorf("get wrong body %s", body)
	}
}

func TestVoteSammaryHandler(t *testing.T) {
	env := os.Getenv("DB_ENV")
	os.Setenv("DB_ENV", "test")
	defer setDBEnv(env)

	req := httptest.NewRequest(http.MethodGet, "/vote/summary", nil)
	got := httptest.NewRecorder()

	VoteSammaryHandler(got, req)

	wantBody := `[
		{"id":"cinnamon","vote":300},
		{"id":"cappuccino","vote":50},
		{"id":"mocha","vote":50},
		{"id":"chiffon","vote":50},
		{"id":"espresso","vote":30},
		{"id":"milk","vote":20},
		{"id":"azuki","vote":10},
		{"id":"coco","vote":70},
		{"id":"nuts","vote":70},
		{"id":"poron","vote":5},
		{"id":"corne","vote":90},
		{"id":"berry","vote":22},
		{"id":"cherry","vote":10}
		]`
	if got.Code != http.StatusOK {
		t.Errorf("want http.StatusOk, but get %d", got.Code)
	}
	if body := got.Body.String(); body != deleteLineAndTab(wantBody) {
		t.Errorf("get wrong body %s", body)
	}
}

func TestUserSummaryHandler(t *testing.T) {
	env := os.Getenv("DB_ENV")
	os.Setenv("DB_ENV", "test")
	defer setDBEnv(env)

	req := httptest.NewRequest(http.MethodGet, "/user/", nil)
	got := httptest.NewRecorder()

	UserSummaryHandler(got, req)

	wantBody := `[
		{"number":20,"age":0,"gender":1},
		{"number":100,"age":1,"gender":1},
		{"number":20,"age":1,"gender":2},
		{"number":10,"age":1,"gender":9}
		]`
	if got.Code != http.StatusOK {
		t.Errorf("want http.StatusOk, but get %d", got.Code)
	}
	if body := got.Body.String(); body != deleteLineAndTab(wantBody) {
		t.Errorf("get wrong body %s", body)
	}
}

func TestCreateUserHandler(t *testing.T) {
	env := os.Getenv("DB_ENV")
	os.Setenv("DB_ENV", "test")
	defer setDBEnv(env)

	values := url.Values{}
	values.Add("age", "1")
	values.Add("gender", "1")
	values.Add("address", "12")

	req := httptest.NewRequest(http.MethodPost, "/user/", strings.NewReader(values.Encode()))
	got := httptest.NewRecorder()

	CreateUserHandler(got, req)

	wantBody := "1"
	if got.Code != http.StatusOK {
		t.Errorf("want http.StatusOk, but get %d", got.Code)
	}
	if body := got.Body.String(); body != deleteLineAndTab(wantBody) {
		t.Errorf("get wrong body %s", body)
	}
}

func TestUserResultHandler(t *testing.T) {
	env := os.Getenv("DB_ENV")
	os.Setenv("DB_ENV", "test")
	defer setDBEnv(env)

	req := httptest.NewRequest(http.MethodGet, "/user/1/1", nil)
	got := httptest.NewRecorder()

	UserResultHandler(got, req)

	wantBody := `[
		{"character":"cinnamon","user":1,"age":0,"gender":0,"address":12,"created_at":"2020-07-24 15:18:00","ip":{"String":"","Valid":false}},
		{"character":"cinnamon","user":1,"age":0,"gender":0,"address":12,"created_at":"2020-07-24 15:18:00","ip":{"String":"","Valid":false}},
		{"character":"mocha","user":2,"age":0,"gender":0,"address":8,"created_at":"2020-07-24 15:18:00","ip":{"String":"","Valid":false}},
		{"character":"coco","user":2,"age":0,"gender":0,"address":8,"created_at":"2020-07-24 15:18:00","ip":{"String":"","Valid":false}},
		{"character":"nuts","user":3,"age":0,"gender":0,"address":33,"created_at":"2020-07-24 15:18:00","ip":{"String":"","Valid":false}}
		]`
	if got.Code != http.StatusOK {
		t.Errorf("want http.StatusOk, but get %d", got.Code)
	}
	if body := got.Body.String(); body != deleteLineAndTab(wantBody) {
		t.Errorf("get wrong body %s", body)
	}
}
