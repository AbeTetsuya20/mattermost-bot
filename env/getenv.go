package env

import (
	"github.com/xanzy/go-gitlab"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"rat/domain"
)

// 環境変数から値を取得する
func Getenv(key string) string {
	return ""
}

// GetenvDebug 環境変数から値を取得して、デバッグ用に出力する
func GetenvDebug() string {
	buf, err := ioutil.ReadFile("/Users/abetetsuya/rat/env/env.yml")
	if err != nil {
		panic(err)
	}

	var tmpUsers domain.Users
	err = yaml.Unmarshal(buf, &tmpUsers)
	if err != nil {
		panic(err)
	}

	domain.PrintUsers(&tmpUsers)

	return ""
}

// TOKEN を取得し、Gitlab Client を作成する
func GetGitlabEnv() *gitlab.Client {
	buf, err := ioutil.ReadFile("/Users/abetetsuya/rat/env/gitlab.yml")
	if err != nil {
		panic(err)
	}

	expaned := os.ExpandEnv(string(buf))

	var tmpGitlab domain.GitlabInfo
	err = yaml.Unmarshal([]byte(expaned), &tmpGitlab)
	if err != nil {
		panic(err)
	}

	tmpGitlabClient := domain.NewGitlabClient(tmpGitlab.Info.Token, tmpGitlab.Info.Url)

	// Gitlab Client を作成
	git, err := domain.MyNewClient(tmpGitlabClient)

	if err != nil {
		panic(err)
	}

	return git
}
