package kece

import (
	"net"
	"sync"
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
