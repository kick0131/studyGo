package firestoresample

import (
	"fmt"
	"time"
)

// NoticeInfo は通知情報テーブルです
type NoticeInfo struct {
	subCollection []NoticeInfoSub
}

// NoticeInfoSub は言語別通知情報テーブルです
type NoticeInfoSub struct {
	collectionName string
	CreateAt       time.Time `firestore:"createAt"`
	Title          string    `firestore:"title"`
	Data           string    `firestore:"notificationData"`
	Dispflag       bool      `firestore:"dispflag"`
	Deleteflag     bool      `firestore:"deleteflag,omitempty"`
	DeletedAt      time.Time `firestore:"deletedAt,omitempty"`
}

// Set はドキュメントに自身のデータをセットします
func (docData NoticeInfo) Set(client *FireBaseClient, collection string) error {
	// サブコレクション用の親ドキュメントを生成
	newdoc, _, _ := client.fb.Collection(collection).Add(client.ctx, docData)
	fmt.Println("newdoc.ID : " + newdoc.ID)
	for _, sub := range docData.subCollection {
		// サブコレクション内にドキュメントを生成
		_, _, err := client.fb.Collection(collection).Doc(newdoc.ID).Collection(sub.collectionName).Add(client.ctx, sub)
		if err != nil {
			return err
		}
	}

	return nil
}
