package bootstrap

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Env struct {
	Database struct {
		User     string
		Password string
		Database string
		Port     int
	}
}

func NewEnv() Env {
	env := Env{}
	if _, err := toml.DecodeFile("config/config.toml", &env); err != nil {
		log.Fatal(err)
	}

	return env
}
