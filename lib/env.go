package lib

import (
	"os"
)

type EnvType struct {
	Watch bool
}

var Env EnvType = EnvType{
	Watch: os.Getenv("WATCH") == "true",
}
