package handlers

import "database/sql"

//Page ... htmlに渡す値をまとめた構造体
type Page struct {
	Title     string
	UserID    string
	LogIn     bool
	Admin     bool
	Character []VoteResult
	Vote      []Vote
	VoteUser  []User
	NewUser   NewAdmin
}

// VoteResult キャラクターごとの得票数をまとめた構造体
type VoteResult struct {
	ID   int `json:"id,string"`
	Name string
	Vote int `json:"vote"`
}

// Vote ユーザーが投票した票を表す構造体
type Vote struct {
	Chara       string         `json:"character"`
	User        int            `json:"user"`
	Age         int            `json:"age"`
	Gender      int            `json:"gender"`
	Address     int            `json:"address"`
	CreatedTime string         `json:"created_at"`
	IP          sql.NullString `json:"ip"`
}

// User 投票に参加した人の構造体
type User struct {
	Num    int `json:"number"`
	Age    int `json:"age"`
	Gender int `json:"gender"`
}

// NewAdmin 新規作成したユーザーの情報
type NewAdmin struct {
	UserID   string
	Password string
	Auth     bool
}

var charas = []VoteResult{
	VoteResult{1, "cinnamon", 0},
	VoteResult{2, "cappuccino", 0},
	VoteResult{3, "mocha", 0},
	VoteResult{4, "chiffon", 0},
	VoteResult{5, "espresso", 0},
	VoteResult{6, "milk", 0},
	VoteResult{7, "azuki", 0},
	VoteResult{8, "coco", 0},
	VoteResult{9, "nuts", 0},
	VoteResult{10, "poron", 0},
	VoteResult{11, "corne", 0},
	VoteResult{12, "berry", 0},
	VoteResult{13, "cherry", 0},
}
