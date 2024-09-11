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
	"log"
	"net/http"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID/go"
)

func main() {
	db := openapi.SetupDB()
	defer db.Close()	

	log.Printf("Server started")

	DefaultAPIService := openapi.NewDefaultAPIService(db)
	DefaultAPIController := openapi.NewDefaultAPIController(DefaultAPIService)

	router := openapi.NewRouter(DefaultAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
