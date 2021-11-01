package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func Db_Config() string {

	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	var (
		Host     = os.Getenv("APP_DB_HOST")
		User     = os.Getenv("APP_DB_USER")
		Password = os.Getenv("APP_DB_PASSWORD")
		DataBase = os.Getenv("APP_DB_NAME")
		Port     = os.Getenv("APP_DB_PORT")
	)

	var APP_DB_CONFIG = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		Host, User, Password, DataBase, Port,
	)

	return APP_DB_CONFIG
}
