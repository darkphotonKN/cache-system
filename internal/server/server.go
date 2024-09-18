package server

import (
	"fmt"
	"net"
	"time"

	peermanager "github.com/darkphotonKN/cache-system/internal/peer_manager"
)

const (
	defaultListenAddr = ":5000"
)

type Config struct {
	ListenAddr string
}

type Server struct {
	Config
	ln net.Listener
	*peermanager.PeerManager
}

func NewServer(cfg Config) *Server {
	if len(cfg.ListenAddr) == 0 {
		// default to default listen address
		cfg.ListenAddr = defaultListenAddr
	}

	pm := peermanager.NewPeerManager()

	return &Server{
		Config:      cfg,
		PeerManager: pm,
	}
}

func (s *Server) StartServer() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	s.ln = ln

	if err != nil {
		fmt.Printf("Error when attempging to create listener: %s", err)
		return err
	}

	// CHECK PEER LIST - TODO: Remove after testing
	go func() {
		ticker := time.NewTicker(time.Second * 10)

		defer ticker.Stop()

		for range ticker.C {
			fmt.Println("Peer list:", s.GetPeers())
		}
	}()

	// start accept loop for peers
	go s.AcceptLoop()
	// start go routine to listen for peer channel setups
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

		// pass new peer connection to add them to the peers map
		peer := peermanager.Peer{
			Conn: conn,
		}

		// send to add peer channel
		s.AddPeerChan <- &peer

		// start read-loop goroutines for each connection
		go s.ReadLoop(conn)
	}
}
