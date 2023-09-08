package utils

import (
	"fmt"
	"os"
)

type envVars struct {
	ENV             string
	PORT            string
	GH_PAT          string
	JWT_SIGNING_KEY []byte
}

var Env envVars

func LoadEnvVars() {
	Env.ENV = os.Getenv("ENV")
	Env.PORT = fmt.Sprintf(":%s", os.Getenv("PORT"))
	Env.GH_PAT = os.Getenv("GH_PAT")
	Env.JWT_SIGNING_KEY = []byte(os.Getenv("JWT_SIGNING_KEY"))
}
