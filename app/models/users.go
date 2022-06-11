package models

import (
	"fmt"
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
}

// Userの作成
// 名前付き戻り値
func (u *User) CreateUser() (err error) {
	createUserCmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values(?,?,?,?,?)`

	// 同一パッケージ間で変数を共有するためpackage名.変数.メソッド()としないでOK
	// ここで実際にUserテーブルに登録する
	_, err = Db.Exec(createUserCmd, createUUID(), u.Name, u.Email, Encrypto(u.PassWord), time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// Userの取得
func GetUser(id int) (user User, err error) {
	user = User{}
	getUserCommand := `select id, uuid, name, email, password, created_at from users where id = ?`
	// idを渡して１行検索するクエリを投げる
	data := Db.QueryRow(getUserCommand, id)
	fmt.Println("---")
	fmt.Println(data)
	fmt.Println("---")
	//　作成したuserに埋め込む
	err = data.Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.PassWord, &user.CreatedAt)
	return user, err
}

// Userの更新
func (u User) UpdateUser() (err error) {
	// idを元にnameとemailを変更
	updateUserCommand := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(updateUserCommand, u.Name, u.Email, u.ID)
	// DEBUG: わざとerrを出力すると...?
	fmt.Println("---")
	fmt.Printf("type is %T, value is %v\n", err, err)
	fmt.Println("---")
	if err != nil {
		log.Fatalln(err) // そのときのエラー構造体を出力する?
	}
	return err

}

// Userの削除
func (u *User) DeleteUser() (err error) {
	deleteUserCommand := `delete from users where id = ?`
	_, err = Db.Exec(deleteUserCommand, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err

}
