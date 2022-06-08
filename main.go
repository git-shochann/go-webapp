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
}
