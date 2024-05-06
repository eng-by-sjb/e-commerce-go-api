package main

import (
	"log"

	"github.com/dev-by-sjb/e-commerce-go-api/cmd/api"
	"github.com/dev-by-sjb/e-commerce-go-api/config"
	"github.com/dev-by-sjb/e-commerce-go-api/db"
)

func main() {
	db, err := db.NewPostgresStorage(db.DBConfig{
		DBHost:     config.Envs.DBHost,
		DBPort:     config.Envs.DBPort,
		DBUser:     config.Envs.DBUser,
		DBPassword: config.Envs.DBPassword,
		DBName:     config.Envs.DBName,
	})
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer("3000", db)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}
