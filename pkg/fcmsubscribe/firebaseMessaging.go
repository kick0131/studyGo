package fcmsubscribe

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

// JSONPATH はGoogleSDKのサービスアカウントファイルです
const JSONPATH = "./data/adminsdk.json"

// FireBaseMessagingClient はFirebaseMessaging用クライアントです
type FireBaseMessagingClient struct {
	client *messaging.Client
	ctx    context.Context
}

// BuildClientWithProjectID はプロジェクト名を使用してFireBaseClientを作成します（GCP環境用）
func BuildClientWithProjectID(projectID string) *FireBaseMessagingClient {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: projectID}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return &FireBaseMessagingClient{
		client: client,
		ctx:    ctx,
	}
}

// BuildClientWithCredential は認証ファイルを使用してFireBaseClientを作成します（ローカル用）
func BuildClientWithCredential(credentialJSONPath string) *FireBaseMessagingClient {
	ctx := context.Background()
	sa := option.WithCredentialsFile(credentialJSONPath)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return &FireBaseMessagingClient{
		client: client,
		ctx:    ctx,
	}
}

// SubscribeToTopic トピックの登録を行う
func (messaging *FireBaseMessagingClient) SubscribeToTopic(registrationTokens []string, topic string) (*messaging.TopicManagementResponse, error) {
	return messaging.client.SubscribeToTopic(messaging.ctx, registrationTokens, topic)
}

// UnsubscribeFromTopic トピックの解除を行う
func (messaging *FireBaseMessagingClient) UnsubscribeFromTopic(registrationTokens []string, topic string) (*messaging.TopicManagementResponse, error) {
	return messaging.client.UnsubscribeFromTopic(messaging.ctx, registrationTokens, topic)
}
