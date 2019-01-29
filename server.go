package kece

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"net"
	"sync"
)

// Server struct
type Server struct {
	clients       map[*Client]bool
	listener      net.Listener
	args          *Arguments
	register      chan *Client
	unregister    chan *Client
	publish       chan []byte
	clientMessage chan *ClientMessage
	commander     Commander
	sync.RWMutex
}

// NewServer function, Server's constructor
func NewServer(args *Arguments, commander Commander) *Server {
	clients := make(map[*Client]bool)
	register := make(chan *Client)
	unregister := make(chan *Client)
	publish := make(chan []byte)
	clientMessage := make(chan *ClientMessage)
	return &Server{
		args:          args,
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

			go processMessage(clientMessage, server.commander, server.args.Auth)
		}
	}

}

// Start function, start Kece server
func (server *Server) Start() error {
	listener, err := net.Listen(server.args.Network, fmt.Sprintf(":%s", server.args.Port))
	if err != nil {
		return err
	}

	fmt.Println("log -> kece server listen on port :", server.args.Port)

	defer listener.Close()

	done := make(chan bool)
	// handle concurrent client
	go server.serveClient()

	// handle concurrent incoming client
	go func() {
		for {
			c, err := listener.Accept()
			if err != nil {
				panic(err)
			}

			//register to every connected client to DB
			server.register <- &Client{ID: c.RemoteAddr().String(), Conn: c}
		}
	}()

	<-done

	return nil

}

func validateAuth(cm *ClientMessage, commander Commander, auth string) error {
	if err := cm.ValidateMessage(); err != nil {
		return err
	}

	clientID := cm.Client.ID

	result, err := commander.Get([]byte(commands["GET"]), []byte(clientID))
	if err != nil {
		return errors.New(ErrorInvalidAuth)
	}

	if !bytes.Equal([]byte(auth), result.Value) {
		return errors.New(ErrorInvalidAuth)
	}

	return nil
}

func processMessage(cm *ClientMessage, commander Commander, auth string) {
	for {
		if err := cm.ValidateMessage(); err != nil {
			cm.Client.Conn.Write([]byte(err.Error()))
			return
		}

		messages := bytes.Split(cm.Message, []byte(" "))

		cmd := messages[0]
		key := messages[1]

		switch string(cmd) {
		case commands["AUTH"]:
			value := messages[1]
			value = bytes.Trim(value, crlf)
			if len(auth) <= 0 {
				reply := replies["ERROR"]
				cm.Client.Conn.Write([]byte(reply))
				return
			}

			if !bytes.Equal([]byte(auth), value) {
				cm.Client.Conn.Write([]byte(ErrorInvalidAuth))
				return
			}

			key = []byte(cm.Client.ID)

			err := commander.Auth(cmd, key, value)
			if err != nil {
				reply := replies["ERROR"]
				cm.Client.Conn.Write([]byte(reply))
				return
			}

			reply := replies["OK"]
			cm.Client.Conn.Write([]byte(reply))
			return
		case commands["SET"]:
			if len(auth) > 0 {
				if err := validateAuth(cm, commander, auth); err != nil {
					cm.Client.Conn.Write([]byte(err.Error()))
					return
				}
			}

			value := messages[2]
			_, err := commander.Set(cmd, key, value)
			if err != nil {
				reply := replies["ERROR"]
				cm.Client.Conn.Write([]byte(reply))
				return
			}

			reply := replies["OK"]
			cm.Client.Conn.Write([]byte(reply))
			return
		case commands["GET"]:
			if len(auth) > 0 {
				if err := validateAuth(cm, commander, auth); err != nil {
					cm.Client.Conn.Write([]byte(err.Error()))
					return
				}
			}

			result, err := commander.Get(cmd, key)
			if err != nil {
				reply := replies["ERROR"]
				cm.Client.Conn.Write([]byte(reply))
				return
			}

			reply := result.Value
			cm.Client.Conn.Write([]byte(reply))
			cm.Client.Conn.Write([]byte(crlf))
			return
		case commands["DEL"]:
			if len(auth) > 0 {
				if err := validateAuth(cm, commander, auth); err != nil {
					cm.Client.Conn.Write([]byte(err.Error()))
					return
				}
			}

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
