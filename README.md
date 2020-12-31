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
- ゴルーチンの勉強(`https://note.crohaco.net/2019/golang-goroutine/`)

### pkg/posgredata
PostgreSQLへの書き込み
- users.go usersテーブルへの読み書き基本コード

### pkg/httpserver
goでHTTPサーバ
- http/templateを使ってHTMLファイルを出力させる
- VuetifyライブラリをCDN経由で使用

## 方針
動かすだけであればmain()が無くてもテストコードを用意してgo testを実行すれば良いので、実行モジュールを作る用途が無ければ、テストコードをエントリポイントとする。

<b>main()が必要なケース</b>

- CloudRunにデプロイして実行させる場合
- 開発環境以外の場所で実行させる場合

## トラブルシュート
- .gcloudignoreがないと、Dockerで.gitignoreのファイルが対象外になる

- CloudRunへのデプロイは専用のRoleが必要

    `https://phpnews.io/feeditem/google-cloud-build-google-cloud-run-fixing-error-gcloud-run-deploy-permission-denied-the-caller-does-not-have-permission`

    ```
    run.services.create and run.services.update on the project level. Typically assigned through the roles/run.admin role.
    iam.serviceAccounts.actAs for the Cloud Run runtime service account. By default, this is PROJECT_NUMBER-compute@developer.gserviceaccount.com. The permission is typically assigned through the roles/iam.serviceAccountUser role.
    ```

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
# Deploy to Cloud Run
prepaire
- make cloudbuild.yaml at root directory

```
gcloud builds submit
```

---
# Another Go Commands
- list go package
`go list -m all`


