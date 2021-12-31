package helper

import (
	"github.com/joho/godotenv"
)

func loadConfig() error {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		err = godotenv.Load("../.env")
		if err != nil {
			return err
		}
	}
	return nil
}
