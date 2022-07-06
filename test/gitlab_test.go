package test

import (
	"fmt"
	"rat/infra"
	"testing"
)

func TestGitlab(t *testing.T) {

	// ユーザーを取得する
	users := infra.GetUsers()
	for _, user := range users {
		fmt.Printf(" user: %#v \n", user)
	}

	// タスクを作成する
	err := infra.CreateTask("Golang でテスト", 5)
	if err != nil {
		panic(err)
	}

	// タスクを取得する
	task := infra.GetTask(1)
	fmt.Printf("task: %#v ", task)

	// Assignee を変更する
	err = infra.ChangeAssignee(1, 5)
	if err != nil {
		panic(err)
	}
}
