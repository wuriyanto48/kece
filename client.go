package kece

import (
	"net"
)

// Client struct
type Client struct {
	ID    string
	Conn  net.Conn
	Send  chan []byte
	Topic string
}

func (client *Client) Read() {

}
