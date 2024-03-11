package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"time"

	_ "github.com/lib/pq"

	"github.com/Mohit-Mamtora/gofinlop/app"
	filelogger "github.com/Mohit-Mamtora/gofinlop/app/logger/filelogger"
	"github.com/Mohit-Mamtora/gofinlop/app/repository"
	"github.com/Mohit-Mamtora/gofinlop/app/routes"
	"github.com/Mohit-Mamtora/gofinlop/app/services"
	"github.com/Mohit-Mamtora/gofinlop/config"
)

func main() {
	log, err := filelogger.NewFileLogger("logs", "log.txt", 1, true)

	if err != nil {
		panic(err)
	}

	defer log.Close()

	/* DB connnection */
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.Dbname)

	nativeDbConnection, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = nativeDbConnection.Ping()
	if err != nil {
		panic(err)
	}

	db, err := repository.InitilizeDb(nativeDbConnection, "postgres")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	dependencyHandler := &app.DependencyHandler{
		Logger: log,
	}

	/* Bootstrap Application */
	repo := repository.InitilizeRepository(db, dependencyHandler)
	service := services.InitilizeService(repo, dependencyHandler)
	server := routes.InitilizeRoute(service, dependencyHandler)

	server.RegisterRoutes()
	go func() {
		if err = server.Start(config.ServerPort); err != nil {
			panic(err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = server.Shutdown(ctx); err != nil {
		panic(err)
	}
}
