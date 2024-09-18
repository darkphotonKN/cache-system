package peermanager

import (
	"fmt"
	"net"
)

const (
	readSizeLimit = 2048
)

type Peer struct {
	Conn net.Conn
}
type PeerList map[*Peer]bool

type PeerManager struct {
	peers PeerList
	// for notifying
	AddPeerChan chan *Peer
}

func NewPeerManager() *PeerManager {
	peers := make(PeerList)
	addPeerChan := make(chan *Peer)

	return &PeerManager{
		peers:       peers,
		AddPeerChan: addPeerChan,
	}
}

func (pm *PeerManager) AddPeer(conn net.Conn) {
	peer := Peer{
		Conn: conn,
	}
	pm.peers[&peer] = true
}

func (pm *PeerManager) AcceptLoop() {
	for {
		select {
		case peer := <-pm.AddPeerChan:
			fmt.Println("Reading in new peer.")
			pm.peers[peer] = true
		default:
		}
	}
}

func (pm *PeerManager) GetPeers() PeerList {
	return pm.peers
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
