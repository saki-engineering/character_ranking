package handlers

//Page ... htmlに渡す値をまとめた構造体
type Page struct {
	Title       string  // <title>タグの部分
	Character   []chara // キャラクター一覧
	Description string  // キャラごとの詳細説明
}

type chara struct {
	ID   int    // エントリーNo.
	Name string // キャラ名
}

var charas = []chara{
	chara{1, "cinnamon"},
	chara{2, "cappuccino"},
	chara{3, "mocha"},
	chara{4, "chiffon"},
	chara{5, "espresso"},
	chara{6, "milk"},
	chara{7, "azuki"},
	chara{8, "coco"},
	chara{9, "nuts"},
	chara{10, "poron"},
	chara{11, "corne"},
	chara{12, "berry"},
	chara{13, "cherry"},
}
