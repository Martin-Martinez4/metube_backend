package config

import (
	"os"

	"github.com/joho/godotenv"
)

var err = godotenv.Load()

type envVars struct {
	DB_URL     string
	JWT_SECRET string

	MAIN_DB_URL string
	TEST_DB_URL string
}

var toReturn *envVars

func GetVars() *envVars {

	if err != nil {
		return nil
	}

	println(os.Getenv("DB_URL"))
	println("here")

	toReturn = &envVars{
		DB_URL:     os.Getenv("DB_URL"),
		JWT_SECRET: os.Getenv("JWT_SECRET"),

		MAIN_DB_URL: os.Getenv("MAIN_DB_URL"),
		TEST_DB_URL: os.Getenv("TEST_DB_URL"),
	}
	return toReturn

}

type Config struct {
	TEST_DB_URL string
	JWT_SECRET  string
}

var config *Config

func ReadEnv(pathToEnv string) *Config {
	currentEnvironment, ok := os.LookupEnv("ENVIRONMENT")

	if config == nil {

		var err error
		if ok {
			err = godotenv.Load(pathToEnv + currentEnvironment)
		} else {
			err = godotenv.Load(pathToEnv)
		}

		if err != nil {
			panic(err)
		}
		TEST_DB_URL, _ := os.LookupEnv("TEST_DB_URL")
		JWT_SECRET, _ := os.LookupEnv("JWT_SECRET")
		config = &Config{
			TEST_DB_URL: TEST_DB_URL,
			JWT_SECRET:  JWT_SECRET,
		}

	}

	// To load the config values
	// TEST_DB_URL, _ := os.LookupEnv("TEST_DB_URL")
	// config := &Config{
	// 	TEST_DB_URL: TEST_DB_URL,
	// }
	return config
}
