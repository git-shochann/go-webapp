package models

import (
	"database/sql"
	"fmt"
	"go-webapp/config"
	"log"
	// _ "github.com/mattn/go-sqlite3"
)

// Userテーブルの作成

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
	cmdUser := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTO INCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

	Db.Exec(cmdUser)
}
