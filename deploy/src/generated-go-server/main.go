// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Tender Management API
 *
 * API для управления тендерами и предложениями.   Основные функции API включают управление тендерами (создание, изменение, получение списка) и управление предложениями (создание, изменение, получение списка).
 *
 * API version: 1.0
 */

package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID/go"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	config := openapi.MustLoad()
	log.Printf("Server started on port %s", config.ServerAddress)
	loggerSlog := setupLogger("local")

	psql, err := openapi.NewStorage(config.PostgresConn, loggerSlog)
	if err != nil {
		log.Fatal(err)
	}
	defer psql.Close()

	ctx := context.Background()
	for {
		err := openapi.InitDataBase(ctx, psql)

		if  err == nil {
			log.Println("Data sucsess  bich!")
			break
		}

		log.Println("Database is not connected. Let's try again in 3 seconds")
		time.Sleep(3 * time.Second)
	}

	DefaultAPIService := openapi.NewDefaultAPIService(psql, loggerSlog)
	DefaultAPIController := openapi.NewDefaultAPIController(DefaultAPIService)
	router := openapi.NewRouter(DefaultAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
