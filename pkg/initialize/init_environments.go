package initialize

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var isEnvLoaded = false

func LoadEnv() {
	if isEnvLoaded {
		return
	}
	err := godotenv.Load(".env")
	if err != nil {
		log.Warn().Msg("Error loading .env file")
	} else {
		isEnvLoaded = true
		log.Debug().Msg(".env file loaded successfully")
	}
}
