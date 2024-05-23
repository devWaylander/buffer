package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Глобальный env конфиг
var GlobalConfig config

type config struct {
	Port        string
	BearerToken string
	GenMock     string
}

func Configure() config {
	// Считываем env файл проекта
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	// Сохраняем env переменные
	var (
		Port        = os.Getenv("PORT")
		BearerToken = os.Getenv("KPI_API_BEARER")
		GenMock     = os.Getenv("GEN_MOCK_REQ")
	)

	if Port == "" || BearerToken == "" || GenMock == "" {
		log.Fatal("failed to read env variables")
	}

	GlobalConfig = config{
		Port:        Port,
		BearerToken: BearerToken,
		GenMock:     GenMock,
	}

	return GlobalConfig
}
