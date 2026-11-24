// Package main
package main

import (
	"log"

	"github.com/DiogenesFerminS/devboard/internal/server"
)

func main() {
	port := ":8080"
	srv := server.New(port)

	log.Println("Server running on port " + port)

	if err := srv.Start(); err != nil {
		log.Fatalf("An unexpected error has ocurred %v", err)
	}

}
