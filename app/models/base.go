package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"go-webapp/config"
	"log"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

// Dbという変数でsqlパッケージの*DB型(構造体)として宣言する -> メソッドを使用するため
var Db *sql.DB

var err error

const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)

func init() {

	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		fmt.Printf("type is %T\n", err) // DEBUG errは構造体
		log.Fatalln(err)
	}

	// usersテーブルの作成
	// S -> 文字列を返す f-> フォーマットする
	createUsersTableCmd := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

	_, err := Db.Exec(createUsersTableCmd)
	if err != nil {
		log.Fatalln(err)
	}

	// todosテーブルの作成
	createTodosTableCmd := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME)`, tableNameTodo)

	_, err = Db.Exec(createTodosTableCmd)
	if err != nil {
		log.Fatalln(err) // log.Fatalln() -> logging.goの標準出力設定を変更しているのでそこに書き出して、その後プログラムをストップする
	}

	createSessionsTablecmd := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING,
		user_id STRING,
		created_at DATETIME
	)`, tableNameSession)

	_, err = Db.Exec(createSessionsTablecmd)
	if err != nil {
		log.Fatalln(err)
	}

}

// uuidを作成する関数 名前付き戻り値
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

// Passwordを作成する
func Encrypt(plaintext string) (cryptext string) {
	// sha1を用いて、ハッシュ化する
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
