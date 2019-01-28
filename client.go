package kece

import (
	"bytes"
	"errors"
	"net"
)

// Client struct
type Client struct {
	ID   string
	Conn net.Conn
}

// Subscribe client method, this function will used by client to subscribe to specific topic
//TODO
func (client *Client) Subscribe(topic string) {

}

// ClientMessage struct
type ClientMessage struct {
	Client  *Client
	Message []byte
}

// ValidateMessage function
func (c *ClientMessage) ValidateMessage() error {
	message := bytes.Trim(c.Message, " ")
	messages := bytes.Split(message, []byte(" "))
	if len(messages) < 3 || len(messages) > 3 {
		return errors.New(ErrorInvalidOperation)
	}

	c.Message = message
	return nil
}
