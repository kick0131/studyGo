# study Go
Multi module sample

## ディレクトリの説明
- data プログラム以外のデータ格納場所
- pkg サブパッケージ

### pkg/firestoresample
Firestoreへの書き込みとテストコード
- firestorebasic.go 読み書き基本コード
- firestoreemulate.go Firestoreエミュレータを使ったサンプル（未完成）

【準備】
1. ディレクトリ内にサービスアカウントファイル(adminsdk.json)を配置

### pkg/go-home
お遊びプログラム

### pkg/gosample
サブパッケージの動作確認用
- zodiac.go ランダムで干支を返すプログラム

### pkg/posgredata
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

### Rule
- filename must be `xxx_test.go`
- funcformat must be `Testxxxx(t *testing.T)`

### 全パッケージ実行
```
go test -v ./...
```

### 特定のパッケージ実行
```
go test -v studygo/pkg/firestoresample
```

### 指定したテストメソッド実行
TestXXXXのXXXX部分を正規表現で抽出
```
go test -v -run "ReadFs"
```

# Another Go Commands

### list go package
`go list -m all`

### Go Doc
`godoc -http=:(port)`


