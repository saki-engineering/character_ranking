package handlers

//Page ... htmlに渡す値をまとめた構造体
type Page struct {
	Title       string  // <title>タグの部分
	Character   []chara // キャラクター一覧
	Description string  // キャラごとの詳細説明
	Prefecture  []pref
	Age         []int
}

type chara struct {
	ID   int    // エントリーNo.
	Name string // キャラ名
}

type pref struct {
	Code int
	Name string
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

var prefecture = []pref{
	pref{1, "北海道"},
	pref{2, "青森県"},
	pref{3, "岩手県"},
	pref{4, "宮城県"},
	pref{5, "秋田県"},
	pref{6, "山形県"},
	pref{7, "福島県"},
	pref{8, "茨城県"},
	pref{9, "栃木県"},
	pref{10, "群馬県"},
	pref{11, "埼玉県"},
	pref{12, "千葉県"},
	pref{13, "東京都"},
	pref{14, "神奈川県"},
	pref{15, "新潟県"},
	pref{16, "富山県"},
	pref{17, "石川県"},
	pref{18, "福井県"},
	pref{19, "山梨県"},
	pref{20, "長野県"},
	pref{21, "岐阜県"},
	pref{22, "静岡県"},
	pref{23, "愛知県"},
	pref{24, "三重県"},
	pref{25, "滋賀県"},
	pref{26, "京都府"},
	pref{27, "大阪府"},
	pref{28, "兵庫県"},
	pref{29, "奈良県"},
	pref{30, "和歌山県"},
	pref{31, "鳥取県"},
	pref{32, "島根県"},
	pref{33, "岡山県"},
	pref{34, "広島県"},
	pref{35, "山口県"},
	pref{36, "徳島県"},
	pref{37, "香川県"},
	pref{38, "愛媛県"},
	pref{39, "高知県"},
	pref{40, "福岡県"},
	pref{41, "佐賀県"},
	pref{42, "長崎県"},
	pref{43, "熊本県"},
	pref{44, "大分県"},
	pref{45, "宮崎県"},
	pref{46, "鹿児島県"},
	pref{47, "沖縄県"},
}
