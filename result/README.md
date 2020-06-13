# web_server

## サービス全景
サイトマップは以下の通り。

```
/   # トップページ
├── login(get)  # ログインフォーム
├── login(post userid=X&password=Y)
│               # ログイン処理
├── logout      # ログアウト処理
├── checkid(post userid=X)
│               # 指定ユーザーIDが使用済みかどうかを調べるAPI
├── result
│    ├─/       				 # キャラクターごとの得票数表示
│    ├──/{name} 			 # キャラクターごとの詳細分析ページ予定
│    └─/user				 # (未実装)
│       └─/{gender}/{agemin} # (未実装)
└── admin
     ├─/                # adminユーザーページ
     ├─/userform(get)   # 新規ユーザー作成フォーム
     ├─/userform(post userid=X&password=Y&admin=Z)
     │                  # 新規ユーザー作成処理
     └─/newuser         # 作成したユーザーのID/パスワードを確認・表示する
```

### /login(post)
/loginから遷移。<br>
ログイン成功なら/に、失敗なら/loginにリダイレクトする。

## バックエンド
### 設定する環境変数
サービス中で使用する環境変数一覧は以下の通り。

- DB_ENV : 値が"production"なら本番環境と判定
- DB_USER : 本番環境でDBにアクセスするユーザー名
- DB_PASS : 本番環境でDBにアクセスするときのパスワード
- DB_ADDRESS : 本番環境で使用するDBのエンドポイント
- API_URL : vote_apiのURLを値として格納

### htmlテンプレートに渡す構造体
`http/template`を利用して渡すGolang構造体を以下のように定義している。
```golang
// http/templateに渡すPage構造体
type Page struct {
	Title     string
	UserID    string
	LogIn     bool
	Admin     bool
	Character []VoteResult
	Vote      []Vote
	NewUser   NewAdmin
}

type VoteResult struct {
	ID   int `json:"id,string"`
	Name string
	Vote int `json:"vote"`
}

type NewAdmin struct {
	UserID   string
	Password string
	Auth     bool
}
```
- Title : `<title>`タグに渡す値
- UserID : 現在ログインしているユーザーのID
- Login : 現在ログインしているかどうか
- Admin : 現在ログインしているのがAdminユーザーかどうか
- Character : 各キャラクターごとの得票数一覧を格納
  - ID : エントリーNo
  - Name : キャラクター名
  - Vote : 得票数
- Vote : 票の情報
  - Chara : 投票したキャラクター
  - User : 投票したユーザー
  - Age : ユーザーの年齢
  - Gender : ユーザーの性別
  - Address : ユーザーの居住都道府県
  - CreatedTime : 投票時間
  - IP : 投票元IP
- NewAdmin : 新規作成したユーザー情報を格納
  - UserID : ユーザーID
  - Password : パスワード
  - Auth : 管理者ユーザーかどうか

### jsonから生成するGolang構造体
APIから得られたjsonデータを、Golangの構造体に直してからデータを利用している。

ここでは、以下の構造体を使用している。

#### VoteResult(http/templateでも使用)
```golang
type VoteResult struct {
	ID   int `json:"id,string"`
	Name string
	Vote int `json:"vote"`
}
```
`json.Unmarshal([]byte, *[]VoteResult)`を利用することで、jsonキー"id"が"ID"に、キー"vote"が"Vote"に変換される。

#### Vote(http/templateでも使用)
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
`json.Unmarshal([]byte, *[]Vote)`を利用することで、タグがjsonキーに変換される。

### sessionフィールド一覧
- userid : ログイン中のユーザーID
- auth : 管理者ユーザーでログインしているかどうか
- newuserid : 新規作成したユーザーのID
- newpassword : 新規作成したユーザーのパスワード