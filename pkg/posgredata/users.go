package posgredata

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// Users はユーザ情報テーブルに対応する型です
type Users struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

// Db はDB操作を行うグローバル変数です
var Db *sql.DB

func init() {
	var err error
	// DB名やユーザー、テーブルは適宜適切な値を入れてください
	Db, err = sql.Open("postgres", "user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
}

// クエリ結果の標準出力を行います
func dispSQLRows(items *sql.Rows) {
	for items.Next() {
		m := Users{}
		items.Scan(&m.ID, &m.UUID, &m.Name, &m.Email, &m.Password, &m.CreatedAt)
		fmt.Println(m)
	}
}

// SampleQuery は一連のクエリ操作確認を行います
func SampleQuery() {

	fmt.Println("----- Query 1 -----")
	rows, err := Db.Query("SELECT * FROM users ORDER BY ID")
	if err != nil {
		return
	}
	dispSQLRows(rows)
	fmt.Println()

	fmt.Println("----- Query 2 -----")
	rows, err = Db.Query("SELECT * FROM users WHERE ID = $1 ORDER BY ID", 1)
	if err != nil {
		return
	}
	dispSQLRows(rows)
	fmt.Println()

	fmt.Println("----- QueryRow -----")
	m := Users{}
	err = Db.QueryRow("SELECT ID,Name,Email FROM users WHERE ID = $1 ORDER BY ID", 1).Scan(&m.ID, &m.Name, &m.Email)
	if err != nil {
		fmt.Printf("== err occured : %s\n", err)
		return
	}
	fmt.Println(m)
	fmt.Println()

	fmt.Println("----- Prepare -----")
	statement := "SELECT * FROM users WHERE ID = $1 ORDER BY ID"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Printf("== err occured : %s\n", err)
		return
	}
	defer stmt.Close()

	rows, err = stmt.Query(1)
	dispSQLRows(rows)
	fmt.Println()

	fmt.Println("----- Exec -----")
	_, err = Db.Exec("INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6)", 100, "1000000000", "user2", "user2@google.com", "pass", "2020/01/01 12:34:56")
	if err != nil {
		fmt.Printf("== err occured : %s\n", err)
		return
	}
	rows, err = Db.Query("SELECT * FROM users ORDER BY ID")
	if err != nil {
		fmt.Printf("== err occured : %s\n", err)
		return
	}
	dispSQLRows(rows)
	_, err = Db.Exec("delete from users where ID=$1", 100)
	if err != nil {
		fmt.Printf("== err occured : %s\n", err)
		return
	}
	fmt.Println()

}
