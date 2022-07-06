package infra

import (
	"github.com/xanzy/go-gitlab"
	"rat/env"
)

// gitlab API を叩く

// ### TASK 関連 ###
// GetTask は task 番号から task を取得する
func GetTask(taskNumber int) *gitlab.Issue {
	git := env.GetGitlabEnv()

	// task を取得
	task, _, err := git.Issues.GetIssue("2", taskNumber)
	if err != nil {
		panic(err)
	}

	return task
}

// 新しい issue を作成する
func CreateTask(title string, assignee int) error {
	git := env.GetGitlabEnv()

	var p []int
	i := assignee
	p = append(p, i)

	// issue を作成
	_, _, err := git.Issues.CreateIssue("2", &gitlab.CreateIssueOptions{
		Title:       gitlab.String(title),
		AssigneeIDs: &p,
	})

	if err != nil {
		return err
	}

	return nil
}

// ### ISSUE 関連 ###

// GetUsers は users を取得する
func GetUsers() []*gitlab.User {
	git := env.GetGitlabEnv()

	//git, err := gitlab.NewClient("zo75TMZyEbXx7cSNLGHV", gitlab.WithBaseURL("http://192.168.0.198:800/api/v4"))

	// users を取得
	user, _, err := git.Users.ListUsers(nil)
	if err != nil {
		panic(err)
	}

	return user
}

// Assignee を変更する
func ChangeAssignee(taskNumber int, assignee int) error {
	git := env.GetGitlabEnv()

	var p []int
	i := assignee
	p = append(p, i)

	// issue を作成
	updateIssueOption := &gitlab.UpdateIssueOptions{
		AssigneeIDs: &p,
	}
	_, _, err := git.Issues.UpdateIssue("2", taskNumber, updateIssueOption)

	if err != nil {
		return err
	}

	return nil
}
