package app

import (
	"context"
	"github.com/Vladislav557/catalog/internal/resources"
	"github.com/Vladislav557/catalog/internal/resources/postgres"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"os"
	"time"
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
	r := resources.RouterInit()
	s := resources.New("8081", r)
	go func() {
		err := s.Start()
		if err != nil {
			panic("failed to start server")
		}
	}()
	quitSigCh := make(chan os.Signal, 1)
	sig := <-quitSigCh
	zap.S().Info("Service stopped by signal ", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown: " + err.Error())
	}
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
