package main

import (
	"fmt"
	"rat/env"
)

func main() {
	fmt.Println("run main.go")

	// server を起動
	// 1. /set:number
	// 2. /attend:number
	// 3. request webex api
	// 4. request google api
	// 5. /help -> コマンド一覧を表示

	// ratbot の主な機能
	// 1. ランダムアサイン
	// 2. meeting
	// 3. work
	// 4. ratportal update
	// 5. 出席管理

	env.GetenvDebug()

}
