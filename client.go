package kece

import (
	"bytes"
	"errors"
	"net"
	"strings"
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
	Cmd     []byte
	Key     []byte
	Value   []byte
}

func isValidValue(val string) (bool, string) {
	var pair = map[string]string{"{": "}", `"`: `"`, "'": "'"}
	if _, ok := pair[val]; ok {
		return false, val
	}

	lastChar := string(val[len(val)-1])

	if res, ok := pair[string(val[0])]; ok {
		if res == lastChar {
			if res == `"` || res == "'" {
				val = val[1 : len(val)-1] // remove prefix & suffix string => ex: "test" -> test
			}
			return true, val
		}
		return false, val
	}

	if len(strings.Fields(val)) <= 1 { // single string without space and list from pair
		return true, val
	}

	return false, val
}

// ValidateMessage function
func (c *ClientMessage) ValidateMessage() error {
	message := bytes.TrimSpace(c.Message)

	messages := strings.Fields(string(message))

	command, ok := commands[messages[0]]
	if !ok {
		return errors.New(ErrorInvalidCommand)
	}

	c.Cmd = []byte(messages[0])
	c.Key = []byte(messages[1])

	if command == "GET" || command == "DEL" || command == "AUTH" {
		if len(messages) < 2 || len(messages) > 2 {
			return errors.New(ErrorInvalidOperation)
		}
	}

	if command == "SET" {
		if len(messages) < 3 {
			return errors.New(ErrorInvalidOperation)
		}

		messValue := strings.TrimSpace(strings.TrimLeft(string(message), strings.Join(messages[:2], " ")))
		isValidVal, value := isValidValue(messValue)
		if !isValidVal {
			return errors.New(ErrorInvalidArgument)
		}
		c.Value = []byte(value)
	}

	c.Message = nil // garbage
	return nil
}
