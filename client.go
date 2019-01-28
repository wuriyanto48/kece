package kece

import (
	"net"
)

// Client struct
type Client struct {
	ID   string
	Conn net.Conn
}

// Subscribe client method, this function will used by client to subscribe to specific topic
func (client *Client) Subscribe(topic string) {

}

// ClientMessage struct
type ClientMessage struct {
	Client  *Client
	Message []byte
}
