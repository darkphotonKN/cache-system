package server

import (
	"fmt"
	"net"
)

const (
	defaultListenAddr = ":5000"
	readSizeLimit     = 2048
)

type Config struct {
	ListenAddr string
}

type Server struct {
	Config
	ln    net.Listener
	stuff string
}

func NewServer(cfg Config) *Server {

	if len(cfg.ListenAddr) == 0 {
		// default to default listen address
		cfg.ListenAddr = defaultListenAddr
	}

	return &Server{
		Config: cfg,
	}
}

func (s *Server) StartServer() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	s.ln = ln

	if err != nil {
		fmt.Printf("Error when attempging to create listener: %s", err)
		return err
	}

	go s.connectionLoop()

	return nil
}

func (s *Server) connectionLoop() {
	for {
		conn, err := s.ln.Accept()

		if err != nil {
			fmt.Println("Error when attempting to connect.")

			continue
		}

		// start read-loop goroutines for each connection

		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {

	for {
		buf := make([]byte, readSizeLimit)

		_, err := conn.Read(buf)

		if err != nil {
			fmt.Println("Error when reading message:", err)
			continue
		}

	}
}
