package main

import (
	"fmt"
	"net"
)

const defaultListenAddr = ":5000"

type Config struct {
	ListenAddr string
}

type Server struct {
	Config
	ln net.Listener
}

func NewServer(cfg Config) *Server {
	if len(cfg.ListenAddr) == 0 {
		// default to default listen address
		cfg.ListenAddr = defaultListenAddr
	}

	listener, err := net.Listen("tcp", cfg.ListenAddr)

	if err != nil {
		fmt.Printf("Error when attempging to create listener: %s", err)
	}

	return &Server{
		Config: cfg,
		ln:     listener,
	}
}

func main() {

	fmt.Println("Running")
}
