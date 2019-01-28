package kece

import (
	"bufio"
	"fmt"
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

//addClient function will push new client to the map clients
func (server *Server) addClient(key *Client, b bool) {
	server.Lock()
	fmt.Printf("log -> new client connected %s\n", key.ID)
	server.clients[key] = b
	server.Unlock()
}

//deleteClient function will delete client by specific key from map clients
func (server *Server) deleteClient(key *Client) {
	server.Lock()
	delete(server.clients, key)
	server.Unlock()
}

func (server *Server) serveClient() {

	for {
		select {
		case client := <-server.register:
			// register client to client collection
			server.addClient(client, true)

			// handle message from client
			go func() {
				defer func() {
					client.Conn.Close()
					server.unregister <- client
				}()

				for {
					message, err := bufio.NewReader(client.Conn).ReadBytes('\n')
					if err != nil {
						server.unregister <- client
						break
					}
					fmt.Printf("Received message : %s", string(message))

					m := "from server\n"
					client.Conn.Write([]byte(m))
				}
			}()
		case client := <-server.unregister:
			if _, ok := server.clients[client]; ok {
				fmt.Printf("client %s unregister\n", client.ID)
				server.deleteClient(client)
			}
		}
	}

}

// Start function, start Kece server
func (server *Server) Start() error {
	listener, err := net.Listen(server.network, fmt.Sprintf(":%s", server.port))
	if err != nil {
		return err
	}

	fmt.Println("log -> kece server listen on port :", server.port)

	defer listener.Close()

	// handle client concurrently
	go server.serveClient()

	for {
		c, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		server.register <- &Client{ID: c.RemoteAddr().String(), Conn: c}
	}

}
