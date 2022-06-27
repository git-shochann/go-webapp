package main

import (
	"fmt"
	"go-webapp/app/controllers"
	"go-webapp/app/models"
	"log"
)

func main() {

	// fmt.Println("##########################")

	// fmt.Println("Configを確認します")

	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// fmt.Println("##########################")

	// fmt.Println("Userを確認します")

	// // init関数を呼び出すためだけに書いたコード
	// fmt.Println(models.Db)

	// // 構造体の初期化

	// // u := &models.User{
	// // 	Name:     "Sho",
	// // 	Email:    "aaa@gmail.com",
	// // 	PassWord: "test000",
	// // }
	// // fmt.Println(u)
	// // u.CreateUser()

	// // Userを取得する
	// u, err := models.GetUserByID(1)
	// fmt.Println(u, err)

	// // Userを更新する
	// // 以下のようにしてUser型のuを上書きしてから、レシーバ.メソッドで実行する
	// u.Name = "Yua"
	// u.Email = "bbb@gmail.com"
	// u.UpdateUser()

	// // 再度表示
	// u, _ = models.GetUserByID(1)
	// fmt.Println(u)

	// // Userの削除
	// u.DeleteUser()

	// // 再度表示 -> 初期値のUser構造体が返ってくる
	// u, _ := models.GetUserByID(2)
	// fmt.Println(u)

	// fmt.Println("##########################")

	// fmt.Println("Todoを確認します")

	// // Todoの作成
	// fmt.Println(u)
	// u.CreateTodo("ラーメンを食べる")

	// fmt.Println("##########################")

	// Todoの取得
	// todo, _ := models.GetTodo(2)
	// fmt.Println(todo)

	// 全部のTodoの取得
	// todo, _ := models.GetAllTodo()

	// // sliceを1つずつ出力する
	// for i, v := range todo {
	// 	fmt.Println(i, v)
	// }

	// 特定ユーザーのTodoを全部取得
	// u, _ := models.GetUserByID(2)
	// // fmt.Println(u)
	// todos, _ := u.GetMultipleTodo()
	// // fmt.Println(todos)
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// Todoの更新
	// todo, _ := models.GetTodo(5)
	// fmt.Println(todo)
	// todo.Content = "家系食べる"
	// todo.UpdateTodo()
	// fmt.Println(todo)

	// Todoの削除
	// t, _ := models.GetTodo(5)
	// fmt.Println(t)
	// t.DeleteTodo()
	// fmt.Println(t)

	controllers.StartMainServer()

	user, _ := models.GetUserByEmail("test@gmail.com")
	// fmt.Println(user)

	sessions, err := user.CreateSession()
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println(sessions)
	// -> {1 7772e68e-f5f3-11ec-a2e4-e68f59fb0199 test@gmail.com 1 2022-06-27 17:30:57.152111 +0900 +0900}
	valid, _ := sessions.CheckSession()
	fmt.Println(valid)

}
