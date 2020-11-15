package firestoresample

import (
	"fmt"
	"time"

	"google.golang.org/api/iterator"
)

// NoticeInfo は通知情報テーブルです
type NoticeInfo struct {
	subCollection []NoticeInfoSub
}

// NoticeInfoSub は言語別通知情報テーブルです
type NoticeInfoSub struct {
	// 配置先のサブコレクション
	collectionName string
	// 親ドキュメントのID
	ParentID   string    `firestore:"parentID"`
	CreateAt   time.Time `firestore:"createAt"`
	Title      string    `firestore:"title"`
	Data       string    `firestore:"notificationData"`
	Dispflag   bool      `firestore:"dispflag"`
	Deleteflag bool      `firestore:"deleteflag,omitempty"`
	DeletedAt  time.Time `firestore:"deletedAt,omitempty"`
}

// Set はドキュメントに自身のデータを追加します
func (docData *NoticeInfo) Set(client *FireBaseClient, collection string) error {
	// サブコレクション用の親ドキュメントを生成
	newdoc, _, _ := client.fb.Collection(collection).Add(client.ctx, docData)
	fmt.Println("newdoc.ID : " + newdoc.ID)
	for _, sub := range docData.subCollection {
		// 親ドキュメントIDを格納
		sub.ParentID = newdoc.ID
		// サブコレクション内にドキュメントを生成
		_, _, err := client.fb.Collection(collection).Doc(newdoc.ID).Collection(sub.collectionName).Add(client.ctx, sub)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetSubcollectionDoc は指定したサブコレクションからドキュメント一覧をマッピングして返します
func (docData *NoticeInfo) GetSubcollectionDoc(client *FireBaseClient, subcollection string) error {

	// 初期化
	docData.subCollection = make([]NoticeInfoSub, 0)

	// ドキュメント一覧取得
	docs := client.fb.CollectionGroup(subcollection).
		Documents(client.ctx)
	var noticeData NoticeInfoSub
	for {
		doc, err := docs.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(doc.Data())

		// ドキュメント内容を取得し、戻り値の情報通知構造体に追記
		err = doc.DataTo(&noticeData)
		if err != nil {
			return err
		}
		fmt.Println(noticeData.Title)
		docData.subCollection = append(docData.subCollection, noticeData)
	}
	fmt.Println("GetSubcollectionDoc success")
	return nil
}
