package utils

import (
	"fmt"
	"os"
)

type envVars struct {
	ENV             string
	PORT            string
	JWT_SIGNING_KEY []byte
}

var Env envVars

func LoadEnvVars() {
	Env.ENV = os.Getenv("ENV")
	Env.PORT = fmt.Sprintf(":%s", os.Getenv("PORT"))
	Env.JWT_SIGNING_KEY = []byte(os.Getenv("JWT_SIGNING_KEY"))
}
