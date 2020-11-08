# study Go
Multi module sample

## ディレクトリの説明
- pkg サブパッケージ
- data プログラム以外のデータ格納場所

### firestoresample
Firestoreへの書き込みとテストコード
- firestorebasic.go 読み書き基本コード
- firestoreemulate.go Firestoreエミュレータを使ったサンプル

【準備】
1. ディレクトリ内にサービスアカウントファイル(adminsdk.json)を配置

### pkg/gosample
サブパッケージの動作確認用
- zodiac.go ランダムで干支を返すプログラム

### posgredata
PostgreSQLへの書き込み

- users.go usersテーブルへの読み書き基本コード


## 方針
動かすだけであればmain()が無くてもテストコードを用意してgo testを実行すれば良いので、実行モジュールを作る用途が無ければ、テストコードをエントリポイントとする。

<b>main()が必要なケース</b>

- CloudRunにデプロイして実行させる場合
- 開発環境以外の場所で実行させる場合

### execute program
```
go run ./main.go
```
### stop program
```
ctrl + c
```

---
# Go Test
Rules
- filename must be `xxx_test.go`
- funcformat must be `Testxxxx(t *testing.T)`

## Run test
### パッケージ指定
```
go test -v studygo/pkg/(pkgname)
```

### 全パッケージ指定
```
go test -v ./...
```

### テスト名指定

正規表現を使用可能

```
例）末尾がsampleのテスト
go test -run "Sample$"
```

---
# Go Doc
`godoc -http=:(port)`


---
# Another Go Commands
- list go package
`go list -m all`


