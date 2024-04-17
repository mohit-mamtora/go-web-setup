package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/mohit-mamtora/go-web-setup/app/logger"
	filelogger "github.com/mohit-mamtora/go-web-setup/app/logger/filelogger"

	"github.com/mohit-mamtora/go-web-setup/app/repository"
	"github.com/mohit-mamtora/go-web-setup/config"
)

func main() {
	log, err := filelogger.NewFileLogger("logs", "log.txt", 1, logger.DebugLevel, true)

	if err != nil {
		panic("logger initialization panic: " + err.Error())
	}

	defer log.Close()

	/* DB connnection */
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName)

	nativeDbConnection, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("%e", err)
	}
	err = nativeDbConnection.Ping()
	if err != nil {
		log.Fatal("%v", err)
	}

	db, err := repository.InitializeDb(nativeDbConnection, "postgres")
	if err != nil {
		log.Fatal("%v", err)
	}

	defer db.Close()

	/** TODO migration code here */
}
