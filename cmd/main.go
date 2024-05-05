package main

import (
	"log"

	"github.com/dev-by-sjb/e-commerce-go-api/cmd/api"
)

func main() {
	server := api.NewAPIServer("3000", nil)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}
