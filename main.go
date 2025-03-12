package main

import (
	"fmt"
	"os"
)

// TODOリストを格納するスライス
var todos []string

func main() {
	// コマンドライン引数を取得
	args := os.Args

	if len(args) < 2 {
		fmt.Println("使い方: add [タスク名] または list")
		return
	}

	command := args[1]

	switch command {
	case "add":
		if len(args) < 3 {
			fmt.Println("タスク名を入力してください")
			return
		}
		task := args[2]
		todos = append(todos, task)
		fmt.Printf("追加しました: %s\n", task)

	case "list":
		fmt.Println("TODOリスト:")
		for i, task := range todos {
			fmt.Printf("%d: %s\n", i+1, task)
		}

	default:
		fmt.Println("不明なコマンドです: add または list を使ってください")
	}
}
