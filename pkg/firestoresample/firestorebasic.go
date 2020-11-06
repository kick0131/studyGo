package firestoresample

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// FireBaseClient はFirebase接続用クライアントです
type FireBaseClient struct {
	fb  *firestore.Client
	ctx context.Context
}

// BuildClientWithProjectID はプロジェクト名を使用してFireBaseClientを作成します（GCP環境用）
func BuildClientWithProjectID(projectID string) *FireBaseClient {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: projectID}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return &FireBaseClient{
		fb:  client,
		ctx: ctx,
	}
}

// BuildClientWithCredential は認証ファイルを使用してFireBaseClientを作成します（ローカル用）
func BuildClientWithCredential(credentialJSONPath string) *FireBaseClient {
	ctx := context.Background()
	sa := option.WithCredentialsFile(credentialJSONPath)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return &FireBaseClient{
		fb:  client,
		ctx: ctx,
	}
}
