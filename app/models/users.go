package models

import (
	"database/sql"
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

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    string
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
	_, err = Db.Exec(createUserCmd, createUUID(), u.Name, u.Email, Encrypt(u.PassWord), time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// Userの取得
func GetUserByID(id int) (user User, err error) {
	user = User{}
	getUserCommand := `select id, uuid, name, email, password, created_at from users where id = ?`
	// idを渡して１行検索するクエリを投げる
	data := Db.QueryRow(getUserCommand, id)
	// fmt.Println("---")
	// fmt.Println(data)
	// fmt.Println("---")
	// 作成したuserに埋め込む
	err = data.Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.PassWord, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("type is %T, value is %v\n", err, err) // ここの型とメッセージを調べてみた
		} else {
			log.Fatalln(err)
		}
	}
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

// Emailを元にDBから検索をかける
func GetUserByEmail(email string) (user User, err error) {
	getUserCommand := `select id, uuid, name, email, password, created_at
	from users where email = ?`
	user = User{} // マッピング用のstructを初期化
	err = Db.QueryRow(getUserCommand, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt)
	if err != nil {
		log.Fatalln(err)
	}
	// err != nil でエラーハンドリングをするので、errもしっかりここで戻り値として返す。
	return user, err

}

// セッションを作成し、そのまま取得する関数
func (u *User) CreateSession() (sessions Session, err error) {
	createSessionCommand := `insert into sessions(
		uuid,
		email,
		user_id,
		created_at) values (?, ?, ?, ?)`
	_, err = Db.Exec(createSessionCommand, createUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	// 上記で作成したセッションをそのまま取得する
	getSessionCommand := `select id, uuid, email, user_id, created_at from sessions where user_id= ? and email = ?`
	sessions = Session{}
	err = Db.QueryRow(getSessionCommand, u.ID, u.Email).Scan(
		&sessions.ID,
		&sessions.UUID,
		&sessions.Email,
		&sessions.UserID,
		&sessions.CreatedAt,
	)

	return sessions, err
}

// セッションがDBに存在するか判定するメソッド
func (sess *Session) CheckSession() (valid bool, err error) {
	CheckSessionCommand := `select id, uuid, email, user_id, created_at
	from sessions where uuid = ?`

	err = Db.QueryRow(CheckSessionCommand, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.UserID,
		&sess.CreatedAt,
	)
	if err != nil {
		valid = false
		return
	}
	if sess.ID != 0 {
		valid = true
	}
	return valid, err
}

// sessionを破棄する(ログアウトにする)
func (sess *Session) DeleteSessionByUUID() (err error) {
	deleteSessionCommand := `delete from sessions where uuid = ?`
	_, err = Db.Exec(deleteSessionCommand, sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
