package configs

import (
	"log"
	"os"
	"strconv"

	godotenv "github.com/joho/godotenv"
)

var (
	DB_HOST               string = "mysql-service"
	DB_USER               string = "root"
	DB_PASSWORD           string = "ada"
	DB_NAME               string = "golangTodoList"
	APP_PORT              string = "9091"
	APP_LISTENING_ADDRESS string = "0.0.0.0"
	DB_PORT               string = "3306"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error while loading .env")
	}

	DB_USER = getStringEnv("DB_USER", DB_USER)
	DB_PASSWORD = getStringEnv("DB_PASSWORD", DB_PASSWORD)
	DB_HOST = getStringEnv("DB_HOST", DB_HOST)
	DB_NAME = getStringEnv("DB_NAME", DB_NAME)
	APP_LISTENING_ADDRESS = getStringEnv("APP_LISTENING_ADDRESS", APP_LISTENING_ADDRESS)

}

func getStringEnv(envKey string, defaultValue string) string {
	envValue, isSet := os.LookupEnv(envKey)
	if isSet {
		return envValue
	}
	return defaultValue
}

func getIntEnv(envKey string, defaultValue int) int {
	envValue, isSet := os.LookupEnv(envKey)
	if isSet {
		value, err := strconv.Atoi(envValue)
		if err != nil {
			panic(err.Error())
		}
		return value
	}
	return defaultValue
}
