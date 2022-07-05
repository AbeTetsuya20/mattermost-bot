package domain

import "github.com/xanzy/go-gitlab"

type GitlabClient struct {
	Token string `yaml:"GITLAB_URL"`
	Url   string `yaml:"GITLAB_TOKEN"`
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
