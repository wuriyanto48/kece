package kece

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"sync"
)

// Server struct
type Server struct {
	clients       map[*Client]bool
	listener      net.Listener
	network       string
	port          string
	register      chan *Client
	unregister    chan *Client
	publish       chan []byte
	clientMessage chan *ClientMessage
	commander     Commander
	sync.RWMutex
}

// NewServer function, Server's constructor
func NewServer(network, port string, commander Commander) *Server {
	clients := make(map[*Client]bool)
	register := make(chan *Client)
	unregister := make(chan *Client)
	publish := make(chan []byte)
	clientMessage := make(chan *ClientMessage)
	return &Server{
		network:       network,
		port:          port,
		clients:       clients,
		register:      register,
		unregister:    unregister,
		publish:       publish,
		clientMessage: clientMessage,
		commander:     commander,
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

					server.clientMessage <- &ClientMessage{Client: client, Message: message}
				}
			}()
		case client := <-server.unregister:
			if _, ok := server.clients[client]; ok {
				fmt.Printf("client %s unregister its connection\n", client.ID)
				server.deleteClient(client)
			}
		case clientMessage := <-server.clientMessage:
			fmt.Printf("Received message : %s from %s\n", string(clientMessage.Message), clientMessage.Client.ID)

			go processMessage(clientMessage, server.commander)
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

	// handle concurrent client
	go server.serveClient()

	for {
		c, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		//register to every connected client to DB
		server.register <- &Client{ID: c.RemoteAddr().String(), Conn: c}
	}

}

func processMessage(cm *ClientMessage, commander Commander) {
	for {
		if err := cm.ValidateMessage(); err != nil {
			cm.Client.Conn.Write([]byte(err.Error()))
			return
		}

		messages := bytes.Split(cm.Message, []byte(" "))

		cmd := messages[0]
		key := messages[1]

		switch string(cmd) {
		case "SET":
			value := messages[2]
			fmt.Println(value)
			_, err := commander.Set(cmd, key, value)
			if err != nil {
				reply := replies["ERROR"]
				cm.Client.Conn.Write([]byte(reply))
				return
			}

			reply := replies["OK"]
			cm.Client.Conn.Write([]byte(reply))
			return
		case "GET":
			result, err := commander.Get(cmd, key)
			if err != nil {
				reply := replies["ERROR"]
				cm.Client.Conn.Write([]byte(reply))
				return
			}

			reply := result.Value
			cm.Client.Conn.Write([]byte(reply))
			cm.Client.Conn.Write([]byte("\x0D\x0A"))
			return
		case "DEL":
			_, err := commander.Delete(cmd, key)
			if err != nil {
				reply := replies["ERROR"]
				cm.Client.Conn.Write([]byte(reply))
				return
			}

			reply := replies["OK"]
			cm.Client.Conn.Write([]byte(reply))
			return
		default:
			cm.Client.Conn.Write([]byte(ErrorInvalidCommand))
			return
		}
	}
}
