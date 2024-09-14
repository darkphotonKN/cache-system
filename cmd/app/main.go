package main

import (
	"fmt"
	"github.com/darkphotonKN/cache-system/internal/server"
	"log"
)

func main() {
	port := ":5100"

	cfg := server.Config{ListenAddr: port}
	s := server.NewServer(cfg)

	err := s.StartServer()

	if err != nil {
		log.Fatal("Error when attmepting to startup server.")
	}

	fmt.Printf("Server is running on port %s\n", port)
}
