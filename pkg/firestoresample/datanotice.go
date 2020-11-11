package firestoresample

import (
	"time"
)

// NoticeInfo は通知情報テーブルです
type NoticeInfo struct {
	CreateAt   time.Time `firestore:"createAt"`
	Title      string    `firestore:"title"`
	Data       string    `firestore:"notificationData"`
	Dispflag   bool      `firestore:"dispflag"`
	Deleteflag bool      `firestore:"deleteflag,omitempty"`
	DeletedAt  time.Time `firestore:"deletedAt,omitempty"`
}

// Set はドキュメントに自身のデータをセットします
// func (n NoticeInfo) Set(client *FireBaseClient, path string) error {
// 	client.fb.Collection("noticeInfo").Add(client.ctx,
// 	return client.Set(path, n)
// }
