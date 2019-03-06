package kece

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"
	"time"
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
	Exp     time.Duration
}

func processingValue(val string) (value string, expiredValue int, err error) {
	var pair = map[string]string{"{": "}", `"`: `"`, "'": "'"}
	if _, ok := pair[val]; ok {
		err = errors.New(ErrorInvalidArgument)
		return
	}

	// calculate if arguments has expired value (suffix is integer)
	decimal := 0
	for i := len(val) - 1; i >= 0; i-- {
		k := string(val[i])
		a, errConv := strconv.Atoi(k)
		if errConv != nil {
			break
		}
		expiredValue += a * int(math.Pow10(decimal))
		decimal++
	}

	if expiredValue != 0 && fmt.Sprint(expiredValue) != val {
		val = strings.TrimRight(val, fmt.Sprint(expiredValue)) // trim argument with expired value
		val = strings.TrimSpace(val)
	}

	lastChar := string(val[len(val)-1])
	value = val
	if res, ok := pair[string(val[0])]; ok {
		if res == lastChar {
			if res == `"` || res == "'" {
				value = val[1 : len(val)-1] // remove prefix & suffix string => ex: "test" -> test
			}
			return
		}
		err = errors.New(ErrorInvalidArgument)
		return
	}

	if len(strings.Fields(val)) <= 1 { // single string without space and list from pair
		return
	}

	err = errors.New(ErrorInvalidArgument)
	return
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

		mess := strings.TrimLeft(string(message), command)
		idx := strings.Index(mess, messages[1])
		if idx < 0 {
			return errors.New(ErrorInvalidArgument)
		}

		value := strings.TrimSpace(mess[idx+len(messages[1]):])
		val, expired, err := processingValue(value)
		if err != nil {
			return err
		}

		c.Value = []byte(val)
		if expired != 0 {
			c.Exp = time.Second * time.Duration(expired)
		}
	}

	c.Message = nil // garbage
	return nil
}
