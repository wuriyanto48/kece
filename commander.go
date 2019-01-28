package kece

import (
	"bytes"
	"errors"
	"sync"
	"time"
)

var (
	commands = map[string]string{
		"SET":     "SET",
		"GET":     "GET",
		"DEL":     "DEL",
		"PUBLISH": "PUBLISH",
		"SUCCESS": "OK\r\n",
	}
)

// Schema database
type Schema struct {
	Key       []byte
	Value     []byte
	Timestamp time.Time
}

// Commander interface
type Commander interface {
	Set(command, key, value []byte) (*Schema, error)
	Get(command, key []byte) (*Schema, error)
	Delete(command, key []byte) (*Schema, error)
	Publish(topic string, command, value []byte) ([]byte, error)
}

// NewCommander function, Commander's constructor
func NewCommander(db map[string]*Schema) Commander {
	return &commander{db: db}
}

type commander struct {
	db map[string]*Schema
	sync.RWMutex
}

// Set will set value to db
func (c *commander) Set(command, key, value []byte) (*Schema, error) {
	_, ok := commands[string(command)]
	if !ok {
		return nil, errors.New(ErrorInvalidCommand)
	}

	// remove line feed (10)/ LF
	value = bytes.Trim(value, "\n")
	value = bytes.Trim(value, "\r")

	c.Lock()
	newData := &Schema{Key: key, Value: value, Timestamp: time.Now()}
	c.db[string(key)] = newData
	c.Unlock()
	return newData, nil
}

// Get will get value from db
func (c *commander) Get(command, key []byte) (*Schema, error) {
	_, ok := commands[string(command)]
	if !ok {
		return nil, errors.New(ErrorInvalidCommand)
	}

	// remove line feed (10)/ LF
	key = bytes.Trim(key, "\n")

	c.RLock()
	value, ok := c.db[string(key)]
	if !ok {
		return nil, errors.New(ErrorEmptyValue)
	}
	c.RUnlock()
	return value, nil
}

// Delete will get value from db
func (c *commander) Delete(command, key []byte) (*Schema, error) {
	_, ok := commands[string(command)]
	if !ok {
		return nil, errors.New(ErrorInvalidCommand)
	}

	// remove line feed (10)/ LF
	key = bytes.Trim(key, "\n")

	c.RLock()
	value, ok := c.db[string(key)]
	if !ok {
		return nil, errors.New(ErrorEmptyValue)
	}
	delete(c.db, string(key))
	c.RUnlock()
	return value, nil
}

// Publish will publish message to specific topic
//TODO
func (c *commander) Publish(topic string, command, value []byte) ([]byte, error) {
	return nil, nil
}
