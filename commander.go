package kece

import (
	"bytes"
	"errors"
	"sync"
)

var (
	commands = map[string]string{
		"AUTH":    "\x41\x55\x54\x48",
		"SET":     "\x53\x45\x54",
		"GET":     "\x47\x45\x54",
		"DEL":     "\x44\x45\x4C",
		"PUBLISH": "\x50\x55\x42\x4C\x49\x53\x48",
	}

	replies = map[string]string{
		"OK":    "+OK\x0D\x0A",
		"ERROR": "-ERROR\x0D\x0A",
	}

	crlf = "\x0D\x0A"
	cr   = "\x0D"
	lf   = "\x0A"

	lock = &sync.Mutex{}
)

// Commander interface
type Commander interface {
	Auth(command, key, value []byte) error
	Set(command, key, value []byte) (*Schema, error)
	Get(command, key []byte) (*Schema, error)
	Delete(command, key []byte) error
	Publish(topic string, command, value []byte) ([]byte, error)
}

// NewCommander function, Commander's constructor
func NewCommander(dataStorage DataStructure) Commander {
	return &commander{ds: dataStorage}
}

type commander struct {
	ds DataStructure
}

// Auth will set auth to kece server
// TODO
func (c *commander) Auth(command, key, value []byte) error {
	lock.Lock()
	defer lock.Unlock()

	_, ok := commands[string(command)]
	if !ok {
		return errors.New(ErrorInvalidCommand)
	}

	// remove line feed and carriage return (13/10)/ CF/LF
	key = bytes.Trim(key, crlf)
	value = bytes.Trim(value, crlf)

	c.ds.Insert(key, value)
	return nil
}

// Set will set value to db
func (c *commander) Set(command, key, value []byte) (*Schema, error) {
	lock.Lock()
	defer lock.Unlock()

	_, ok := commands[string(command)]
	if !ok {
		return nil, errors.New(ErrorInvalidCommand)
	}

	// remove line feed and carriage return (13/10)/ CF/LF
	key = bytes.Trim(key, crlf)
	value = bytes.Trim(value, crlf)

	newData := c.ds.Insert(key, value)
	return newData, nil
}

// Get will get value from db
func (c *commander) Get(command, key []byte) (*Schema, error) {
	lock.Lock()
	defer lock.Unlock()

	_, ok := commands[string(command)]
	if !ok {
		return nil, errors.New(ErrorInvalidCommand)
	}

	// remove line feed and carriage return (13/10)/ CF/LF
	key = bytes.Trim(key, crlf)
	return c.ds.Search(key)
}

// Delete will get value from db
func (c *commander) Delete(command, key []byte) error {
	lock.Lock()
	defer lock.Unlock()

	_, ok := commands[string(command)]
	if !ok {
		return errors.New(ErrorInvalidCommand)
	}

	// remove line feed and carriage return (13/10)/ CF/LF
	key = bytes.Trim(key, crlf)

	return c.ds.Delete(key)
}

// Publish will publish message to specific topic
//TODO
func (c *commander) Publish(topic string, command, value []byte) ([]byte, error) {
	return nil, nil
}
