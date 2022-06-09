package models

import "time"

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
}

// 名前付き戻り値
func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values(?,?,?,?,?)
	)`

	// 同一パッケージ間で変数を共有するためpackage名.変数.メソッド()としない。
	// result, err := Db.Exec(cmd)
}
