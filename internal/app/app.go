package app

import (
	"github.com/Vladislav557/catalog/internal/resources/postgres"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
)

func Run() {
	defer postgres.Close()
	configInit()
	loggerInit()
	databaseInit()
	serverInit()
}

func databaseInit() {
	postgres.Init(os.Getenv("DATABASE_URL"))
}

func serverInit() {
	zap.L().Info("starting server")
}

func loggerInit() {
	var err error
	var logger *zap.Logger
	if os.Getenv("APP_ENV") == "" {
		panic("can`t find APP_ENV")
	}
	switch os.Getenv("APP_ENV") {
	case "prod":
		logger, err = zap.NewProduction()
	case "dev":
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		panic("can`t create logger: " + err.Error())
	}
	zap.ReplaceGlobals(logger)
}

func configInit() {
	if env := os.Getenv("APP_ENV"); env == "" {
		if err := godotenv.Load(); err != nil {
			panic("can`t configure envs")
		}
	}
}
