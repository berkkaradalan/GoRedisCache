package bootstrap

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct{
	SERVER_ADDRESS							string	`mapstructure:"SERVER_ADDRESS"`
	PORT 									string	`mapstructure:"PORT"`
	CONTEXT_TIMEOUT 						int		`mapstructure:"CONTEXT_TIMEOUT"`
	FOOTBALL_DATA_AREA_API_URL				string	`mapstructure:"FOOTBALL_DATA_AREA_API_URL"`
	FOOTBALL_DATA_TEAM_API_URL				string	`mapstructure:"FOOTBALL_DATA_TEAM_API_URL"`
	FOOTBALL_DATA_PLAYER_API_URL			string	`mapstructure:"FOOTBALL_DATA_PLAYER_API_URL"`
	FOOTBALL_DATA_MATCH_API_URL				string	`mapstructure:"FOOTBALL_DATA_MATCH_API_URL"`
	FOOTBALL_DATA_API_KEY					string	`mapstructure:"FOOTBALL_DATA_API_KEY"`
	MONGODB_URL								string	`mapstructure:"MONGODB_URL"`
	MONGODB_PORT							string	`mapstructure:"MONGODB_PORT"`
	MONGODB_DB_NAME							string	`mapstructure:"MONGODB_DB_NAME"`
	MONGODB_USERNAME 						string	`mapstructure:"MONGODB_USERNAME"`
	MONGODB_PASSWORD 						string	`mapstructure:"MONGODB_PASSWORD"`
	REDIS_URL								string	`mapstructure:"REDIS_URL"`
	REDIS_PORT								string	`mapstructure:"REDIS_PORT"`
	REDIS_PASSWORD							string	`mapstructure:"REDIS_PASSWORD"`
}

func NewEnv() (*Env){
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	contextTimeoutStr := os.Getenv("CONTEXT_TIMEOUT")
	contextTimeout, err := strconv.Atoi(contextTimeoutStr)
	if err != nil {
		log.Fatalf("Invalid CONTEXT_TIMEOUT value in .env file: %v", err)
	}


	return &Env{
		SERVER_ADDRESS: os.Getenv("SERVER_ADDRESS"),
		PORT: os.Getenv("PORT"),
		FOOTBALL_DATA_AREA_API_URL: os.Getenv("FOOTBALL_DATA_AREA_API_URL"),
		FOOTBALL_DATA_TEAM_API_URL: os.Getenv("FOOTBALL_DATA_TEAM_API_URL"),
		FOOTBALL_DATA_PLAYER_API_URL: os.Getenv("FOOTBALL_DATA_PLAYER_API_URL"),
		FOOTBALL_DATA_MATCH_API_URL: os.Getenv("FOOTBALL_DATA_MATCH_API_URL"),
		FOOTBALL_DATA_API_KEY: os.Getenv("FOOTBALL_DATA_API_KEY"),
		MONGODB_DB_NAME: os.Getenv("MONGODB_DB_NAME"),
		MONGODB_USERNAME: os.Getenv("MONGODB_USERNAME"),
		MONGODB_PASSWORD: os.Getenv("MONGODB_PASSWORD"),
		MONGODB_URL:  os.Getenv("MONGODB_URL"),
		MONGODB_PORT: os.Getenv("MONGODB_PORT"),
		REDIS_URL: os.Getenv("REDIS_URL"),
		REDIS_PORT: os.Getenv("REDIS_PORT"),
		REDIS_PASSWORD: os.Getenv("REDIS_PASSWORD"),
		CONTEXT_TIMEOUT:  contextTimeout,
	}
}