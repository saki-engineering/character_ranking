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

var desp = map[string]string{
	"cinnamon":   "遠いお空の雲の上で生まれた、白いこいぬの男の子。特技は、大きな耳をパタパタさせて、空を飛ぶこと。",
	"cappuccino": "のんびり屋さんで、食いしん坊",
	"mocha":      "オシャレでおしゃべり。優しくて面倒見がいい、みんなのお姉さん的存在",
	"chiffon":    "いつも元気いっぱい。細かいことは気にしない。ムードメーカー",
	"espresso":   "ワンちゃんコンテストで優勝したこともある、おぼっちゃま",
	"milk":       "シナモンのように、いつか空を飛びたいと思ってる",
	"azuki":      "おっとりしていて、若干テンポがズレているが、ぼそっとするどいツッコミを入れたりする。趣味は、縁側でお茶をすすること。",
	"coco":       "カプチーノの双子の弟。ココがお兄さん。お調子もの。",
	"nuts":       "カプチーノの双子の弟。ナッツが弟。食いしん坊で、寝坊助",
	"poron":      "コルネに乗ってやってきた謎の女の子。雲みたいなふわふわのお耳と空色のリボンがチャームポイント。",
	"corne":      "チョココルネみたいにクルクルねじれたツノがチャームポイント。時空を超えて、過去や未来、絵本の世界へ行くことができる時空のたびびと。",
	"berry":      "大きな角が生えた悪魔の男のコ。シナモンをライバルだと思っている。",
	"cherry":     "ベリーといっしょに行動している悪魔の女のコ。勝ち気でワガママ。エスプレッソのことが、気になっているみたい。",
}
