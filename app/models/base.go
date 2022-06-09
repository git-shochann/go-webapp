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

const tableNameUser string = "users"

func init() {
	Db, err := sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		fmt.Printf("type is %T\n", err) // DEBUG errは構造体
		log.Fatalln(err)
	}
	// S -> 文字列を返す f-> フォーマットする
	cmd := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTO INCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

	Db.Exec(cmd)
}

// uuidを作成する関数 名前付き戻り値
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

// Passwordを作成する
func Encrypto(plaintext string) (cryptext string) {
	// sha1を用いて、ハッシュ化する
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
