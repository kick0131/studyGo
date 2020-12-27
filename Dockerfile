# 引用
# https://qiita.com/melty_go/items/6ec8b1e423dc1d63da24

# 「Dockerfile」は、プログラムのビルドでよく利用されるmakeツールの「Makefile」ファイルと同様に、
#  Dockerコンテナーの構成内容をまとめて記述するシンプルなテキスト形式のファイルです。
#
# 補足として、このファイルはマルチステージビルドという機能を利用しています。
# 「ビルド依存のライブラリ」と「ランタイム(プログラムの実行)依存のライブラリ」を分離できる機能です。
# これにより、Dockerイメージのサイズを大幅に削減できます。

# ステージ１
# golangのビルド用にgolang-alpine でDockerイメージ(コンテナで利用するファイル・設定)を作成
ARG GO_VERSION=1.15
ARG ALPINE_VERSION=latest
ARG BINFILE=studygo
ARG LOCAL_WORKDIR=localwork/

FROM golang:${GO_VERSION}-alpine AS builder
# [Golang]$GOPATH/go.mod exists but should notを回避する
# https://selfnote.work/20200622/programming/golang-error1/
ENV GOPATH=
ARG BINFILE
ARG LOCAL_WORKDIR

WORKDIR ${LOCAL_WORKDIR}

# カレントディレクトリをコンテナの「/go/src/hello-app」へ追加
# ADD . /go/src/hello-app
# カレントディレクトリをコンテナの${LOCAL_WORKDIR}へ追加
ADD . .
RUN pwd
RUN ls -la

# hello-appの実行ファイルを生成
ARG CGO_ENABLED=0
ARG GOOS="linux" 
RUN go build -o /go/bin/main

# ステージ２
# Alpine Linux(セキュアで軽量な Linux ディストリビューション)ベースでDockerイメージを作成
FROM alpine:${ALPINE_VERSION}
ARG  BINFILE
ARG LOCAL_WORKDIR

WORKDIR /root/

# 「ステージ１」で生成された実行ファイルをこの新しいステージへコピー
# COPY --from=builder /go/bin/${BINFILE} .
# COPY --from=builder ${LOCAL_WORKDIR}${BINFILE} .
COPY --from=0 /go/bin/main .

# ポート番号を指定
ENV PORT 8080

# コンテナ内で実行するコマンドで実行ファイルを実行
# CMD ["./${BINFILE}"]
ENTRYPOINT ["/root/main"]
