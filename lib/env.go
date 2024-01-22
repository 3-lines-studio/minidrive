package lib

import (
	"os"
)

type EnvType struct {
	Watch              bool
	SessionSecret      string
	GithubClientID     string
	GithubClientSecret string
}

var Env EnvType = EnvType{
	Watch:              os.Getenv("WATCH") == "true",
	SessionSecret:      os.Getenv("SESSION_SECRET"),
	GithubClientID:     os.Getenv("GITHUB_CLIENT_ID"),
	GithubClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
}
