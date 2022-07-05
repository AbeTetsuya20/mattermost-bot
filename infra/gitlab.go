package infra

import (
	"rat/domain"
	"rat/env"
)

// gitlab API を叩く
// task 番号から task を取得する

// task を取得
func GetTask(taskNumber int) string {

	gitlabClient := env.GetGitlabEnv()

	_, err := domain.MyNewClient(gitlabClient)
	if err != nil {
		panic(err)
	}

	return "task"
}
