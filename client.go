package kece

import (
	"net"
)

// Client struct
type Client struct {
	ID   string
	Conn net.Conn
}

func (client *Client) Read() {

}

func (client *Client) Subscribe(topic string) {

}
