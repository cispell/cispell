package config

import (
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	githubOAuth "golang.org/x/oauth2/github"
	"log"
)

var (
	Env map[string]string
)

func LoadEnvData() {
	env, err := godotenv.Read()
	if err != nil {
		log.Fatal("Unable to read environment variables from .env")
	}
	Env = env
}

func Oauth() oauth2.Config {
	oauthConf := oauth2.Config{
		ClientID:     Env["GITHUB_CLIENT_ID"],
		ClientSecret: Env["GITHUB_CLIENT_SECRET"],
		Scopes:       []string{"user:email", "repo"},
		RedirectURL:  Env["WEBSITE_URL"] + "/login/github",
		Endpoint:     githubOAuth.Endpoint,
	}
	return oauthConf
}

// Random string for oauth2 API calls to protect against CSRF
func OauthStateString() string {
	return "thisShouldBeARandomString"
}
