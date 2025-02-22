package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var DB *sql.DB

func Init(url string) {
	zap.L().Info("start to connecting to postgres db")
	var err error
	DB, err = sql.Open("postgres", url)
	if err != nil {
		panic("error when connection to postgres: " + err.Error())
	}
	if err = DB.Ping(); err != nil {
		panic("error when ping postgres: " + err.Error())
	}
	zap.L().Info("successfully connect ro postgres")
}

func Close() {
	zap.L().Info("Start to closing postgres connection")
	if err := DB.Close(); err != nil {
		panic("Error when closing postgres connection: " + err.Error())
	}
	zap.L().Info("Ending to closing postgres connection")
}
