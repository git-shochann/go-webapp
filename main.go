package main

import (
	"fmt"
	"go-webapp/app/models"
	"go-webapp/config"
	"log"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	log.Println("aaa")

	// init関数を呼び出すためだけに書いたコード
	fmt.Println(models.Db)

	// 構造体の初期化

	// u := &models.User{
	// 	Name:     "Sho",
	// 	Email:    "aaa@gmail.com",
	// 	PassWord: "test000",
	// }
	// fmt.Println(u)
	// u.CreateUser()

	// Userを取得する
	u, err := models.GetUser(1)
	fmt.Println(u, err)

	// Userを更新する
	// 以下のようにしてUser型のuを上書きしてから、レシーバ.メソッドで実行する
	u.Name = "Yua"
	u.Email = "bbb@gmail.com"
	u.UpdateUser()

	// 再度表示
	u, _ = models.GetUser(1)
	fmt.Println(u)

	// Userの削除
	u.DeleteUser()

	// 再度表示
	u, _ = models.GetUser(1)
	fmt.Println(u)

}
