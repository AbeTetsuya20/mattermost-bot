package domain

import "github.com/xanzy/go-gitlab"

type GitlabInfo struct {
	Info GitlabClient `yaml:"Gitlab"`
}

type GitlabClient struct {
	Url   string `yaml:"GITLAB_URL"`
	Token string `yaml:"GITLAB_TOKEN"`
}

type Issue struct {
	Id       int
	Title    string
	assignee string
}

func NewGitlabClient(token, url string) *GitlabClient {
	return &GitlabClient{
		Token: token,
		Url:   url,
	}
}

func MyNewClient(client *GitlabClient) (*gitlab.Client, error) {
	return gitlab.NewClient(client.Token, gitlab.WithBaseURL(client.Url))
}
