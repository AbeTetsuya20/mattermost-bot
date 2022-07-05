package env

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
