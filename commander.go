package kece

import (
	"bytes"
	"errors"
	"sync"
	"time"
)

var (
	commands = map[string]string{
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

	// remove line feed and carnige return (13/10)/ CF/LF
	value = bytes.Trim(value, crlf)

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

	// remove line feed and carnige return (13/10)/ CF/LF
	key = bytes.Trim(key, crlf)

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

	// remove line feed and carnige return (13/10)/ CF/LF
	key = bytes.Trim(key, crlf)

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
