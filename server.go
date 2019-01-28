package kece

import (
	"fmt"
	"net"
	"sync"
)

var (
	commands = map[string]string{
		"SET":     "SET",
		"GET":     "GET",
		"PUBLISH": "PUBLISH",
	}
)

// Server struct
type Server struct {
	clients    map[*Client]bool
	listener   net.Listener
	network    string
	port       string
	register   chan *Client
	unregister chan *Client
	publish    chan []byte
	sync.RWMutex
}

// NewServer function
func NewServer(network, port string) *Server {
	clients := make(map[*Client]bool)
	register := make(chan *Client)
	unregister := make(chan *Client)
	publish := make(chan []byte)
	return &Server{
		network:    network,
		port:       port,
		clients:    clients,
		register:   register,
		unregister: unregister,
		publish:    publish,
	}
}

//AddClient function will push new client to the map clients
func (server *Server) AddClient(key *Client, b bool) {
	server.Lock()
	fmt.Printf("log -> new client connected %s\n", key.ID)
	server.clients[key] = b
	server.Unlock()
}

//DeleteClient function will delete client by specific key from map clients
func (server *Server) DeleteClient(key *Client) {
	server.Lock()
	delete(server.clients, key)
	server.Unlock()
}
