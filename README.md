# study Go
Multi module sample

## PJ内の命名規約について
- ファイル名はスネークケース(xxx_xxx.go)

## ディレクトリの説明
- data プログラム以外のデータ格納場所
- pkg サブパッケージ

### pkg/firestoresample
Firestoreへの書き込みとテストコード
- firestorebasic.go 読み書き基本コード
- firestoreemulate.go Firestoreエミュレータを使ったサンプル（未完成）

【準備】
1. ディレクトリ内にサービスアカウントファイル(adminsdk.json)を配置

### pkg/disample
DI(依存性の注入)サンプル
- adminsdkjson.go adminsdk.jsonの読み書きを題材にしたインタフェース
- adminsdkjson_test.go DIを使った実装

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

### キャッシュ無しで実行(countオプション)
エミュレータを使う場合等、キャッシュを使うと書き込みが実行されないので毎回このオプションを付けた方が良い
```
go test -v -count=1 ./...
```

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


