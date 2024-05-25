package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DatabaseUrl string
}

func GetEnv() (e Env, err error) {
	err = godotenv.Load()

	if err != nil {
		return
	}

	e.DatabaseUrl = fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_DB"))
	return
}
