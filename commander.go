package kece

import (
	"errors"
	"sync"
)

var (
	commands = map[string]string{
		"SET":     "SET",
		"GET":     "GET",
		"PUBLISH": "PUBLISH",
	}
)

// Commander interface
type Commander interface {
	Set(command, key, value []byte) ([]byte, error)
	Get(command, key []byte) ([]byte, error)
	Publish(topic string, command, value []byte) ([]byte, error)
}

// NewCommander function, Commander's constructor
func NewCommander(db map[string][]byte) Commander {
	return &commander{db: db}
}

type commander struct {
	db map[string][]byte
	sync.RWMutex
}

// Set will set value to db
func (c *commander) Set(command, key, value []byte) ([]byte, error) {
	_, ok := commands[string(command)]
	if !ok {
		return nil, errors.New(ErrorInvalidCommand)
	}

	c.Lock()
	c.db[string(key)] = value
	c.Unlock()
	return value, nil
}

// Get will get value from db
func (c *commander) Get(command, key []byte) ([]byte, error) {
	_, ok := commands[string(command)]
	if !ok {
		return nil, errors.New(ErrorInvalidCommand)
	}

	c.RLock()
	value, ok := c.db[string(key)]
	if !ok {
		return nil, errors.New(ErrorEmptyValue)
	}
	c.RUnlock()
	return value, nil
}

// Publish will publish message to specific topic
//TODO
func (c *commander) Publish(topic string, command, value []byte) ([]byte, error) {
	return nil, nil
}
