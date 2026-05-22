package main

import (
	"TestTaskAPI/configs"
	"TestTaskAPI/internal/departament"
	_ "TestTaskAPI/migrations"
	"TestTaskAPI/pkg/db"
	"fmt"
	"net/http"
)

func main() {

	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	DepartRepository := departament.NewDepartRepository(db)

	departament.NewDepartHandler(router, departament.DepartHandlerDeps{
		DepartRepository: DepartRepository,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
