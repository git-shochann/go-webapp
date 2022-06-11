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

	// 構造体を初期化する
	u := &models.User{
		Name:     "Sho",
		Email:    "tarochann@gmail.com",
		PassWord: "tarotest000",
	}
	fmt.Println(u)
	u.CreateUser()

	// Userを取得する
	returnValue, err := models.GetUser(1)
	fmt.Println(returnValue, err)
}
