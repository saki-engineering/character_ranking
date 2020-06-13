# web_server

## サービス全景
サイトマップは以下の通り。

```
/   # トップページ
├── characters
│    ├─/       # キャラクター一覧&投票
│    └─/{name}
│       ├─/       # {name}キャラの詳細&投票
│       ├─/vote(post character=X)
│       │         # キャラXに投票(2回目以降の投票)
│       └─/voted  # {name}キャラへの投票完了画面
├── form
│    ├─/    # アンケートフォームのページ
│    └─/vote(post age=X&gender=Y&address=Z&character=W)
│           # フォーム経由での投票処理
├── about   # aboutページ
└── faq     # FAQページ
```

### /characters/{name}/vote
/characters/と/characters/{name}/から遷移。<br>
一度でもformに回答したことがあるのならば/characters/{name}/voteに、ないのならば/formに遷移する。

### /form
/characters/{name}/voteからのリダイレクト以外ではアクセス不可能。<br>
入力したら/form/voteに遷移する。

### /form/vote
/formから遷移。ここでの処理が終わったら/characters/{name}/votedにリダイレクトされる。

## フロントエンド
### ローカルストレージに設定されている変数
ローカルストレージには、キー`ranking`が保存されるようになっている。

`ranking`の値には以下が設定される。
```js
var ranking = {
            date: countDay(),
            age: 0,
            area: 99,
            fromCharaId: [],
            gender: 9,
            idAnswered: false,
            votingHistory: [],
            votingTodayHistory: []
        };
```

- date : `countDay()`関数により、2020/6/1からの差分[日]<br>
votinTodayHistoryのクリア判定に使用
- age : formに入力されたユーザーの年齢
- area : formに入力された居住都道府県コード
- fromCharaId : 直前に投票を行ったキャラクターの名前<br>
voting(Today)Historyへの追加処理に使用
- gender : formに入力されたユーザーの性別
- isAnswered : formに回答したことがあるかどうか
- votingHistory : 今までの投票履歴
- votingTodayHistory : 当日の投票履歴<br>
当日中の同キャラクターへの投票ボタン無効化処理に使用

## バックエンド
### 設定する環境変数
サービス中で使用する環境変数一覧は以下の通り。

- API_URL : vote_apiのURLを値として格納

### htmlテンプレートに渡す構造体
`http/template`を利用して渡すGolang構造体を以下のように定義している。
```golang
// http/templateに渡すPage構造体
type Page struct {
	Title       string
	Character   []chara
	Description string
	Prefecture  []pref
	Age         []int
}

type chara struct {
	ID   int
	Name string
}

type pref struct {
	Code int
	Name string
}
```
- Title : `<title>`タグに渡す値
- Character : エントリーキャラ一覧の配列
  - ID : エントリーNo
  - Name : キャラクター名
- Description : 1キャラクターのアピール文
- Prefecture : フォームに渡す都道府県の一覧
  - Code : 都道府県コード
  - Name : 都道府県名
- Age : フォームに渡す年齢一覧

### sessionフィールド一覧
- user : DBのuserテーブルのID
- voting : 投票動作を行っている最中かどうか。/formへの不正遷移防止に利用する。