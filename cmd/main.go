package main

import (
	"TestTaskAPI/configs"
	"TestTaskAPI/db"
	_ "TestTaskAPI/migrations"
	"fmt"
	"net/http"
)

func main() {

	conf := configs.LoadConfig()
	db := db.NewDatabase(conf)
	router := http.NewServeMux()

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
