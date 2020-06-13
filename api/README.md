# vote_api

## サービス全景
サイトマップは以下の通り。

```
/   # 稼動確認用にhello,worldを表示
├── vote
│    ├─/(get)   # 全投票データをjsonで取得
│    ├─/(post character=X&user=Y) 
│    │          # ユーザーYのキャラクターXへの投票処理
│    ├─/summary # 各キャラの得票数一覧をjsonで取得
│    └─/{name}  # {name}キャラの投票データをjsonで取得
└── user
     ├─/(post age=X&gender=Y&address=Z)
     │                    # ユーザーの作成→主キーのIDをbodyに入れて返却
     └─/{gender}/{agemin} # ある性別・年齢層の投票データをjsonで取得
```

## バックエンド
### 設定する環境変数
バックエンドで使用する環境変数一覧は以下の通り。

- DB_ENV : 値が"production"なら本番環境と判定
- DB_USER : 本番環境でDBにアクセスするユーザー名
- DB_PASS : 本番環境でDBにアクセスするときのパスワード
- DB_ADDRESS : 本番環境で使用するDBのエンドポイント

### json生成のために定義された構造体
jsonレスポンス作成のためには、以下のような処理を行っている。

SQLクエリの結果→Golang構造体→json

このサービスでは、以下のようなGolang構造体を利用している。

#### Vote
```golang
type Vote struct {
	Chara       string         `json:"character"`
	User        int            `json:"user"`
	Age         int            `json:"age"`
	Gender      int            `json:"gender"`
	Address     int            `json:"address"`
	CreatedTime string         `json:"created_at"`
	IP          sql.NullString `json:"ip"`
}
```
「誰が、どのキャラクターに投票したのか」という票を表す構造体。`json.Marshal([]Vote)`実行時には、タグに記載されているjsonキーに変換される。

#### Result
```golang
type Result struct {
	CharaID string `json:"id"`
	Vote    int    `json:"vote"`
}
```
「エントリーNoのキャラクターが、何票集めたか」ということを表す構造体。`json.Marshal([]Result)`実行時には、タグに記載されているjsonキーに変換される。`/vote/sammary`でのみ使用。