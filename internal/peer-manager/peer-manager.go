package peermanager

import (
	"fmt"
	"net"
)

const (
	readSizeLimit = 2048
)

type Peer struct {
	conn net.Conn
}

type PeerManager struct {
	peers       map[*Peer]bool
	addPeerChan chan *Peer
}

func NewPeerManager() *PeerManager {
	peers := make(map[*Peer]bool)
	addPeerChan := make(chan *Peer)

	return &PeerManager{
		peers:       peers,
		addPeerChan: addPeerChan,
	}
}

func (pm *PeerManager) AddPeer(conn net.Conn) {
	peer := Peer{
		conn: conn,
	}
	pm.peers[&peer] = true
}

func (pm *PeerManager) ReadLoop(conn net.Conn) {
	defer func() {
		fmt.Println("Closing connection for peer:", conn)
		conn.Close()
	}()

	fmt.Println("New connection...")

	for {
		conn.Write([]byte("Connected to server."))

		buf := make([]byte, readSizeLimit)

		_, err := conn.Read(buf)

		if err != nil {
			fmt.Println("Error when reading message:", err)
			continue
		}

		fmt.Printf("Message received %s", buf)
	}
}
