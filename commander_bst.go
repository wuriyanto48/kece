package kece

import (
	"bytes"
	"errors"
	"time"
)

type commanderBst struct {
	tree *BST
}

// NewCommanderBST constructor using BST data structure
func NewCommanderBST() Commander {
	bst := new(commanderBst)
	bst.tree = NewBST()
	return bst
}

func (c *commanderBst) Auth(command, key, value []byte) error {
	return nil
}

// Set will set value to db
func (c *commanderBst) Set(command, key, value []byte) (*Schema, error) {
	lock.Lock()
	defer lock.Unlock()

	_, ok := commands[string(command)]
	if !ok {
		return nil, errors.New(ErrorInvalidCommand)
	}

	// remove line feed and carnige return (13/10)/ CF/LF
	key = bytes.Trim(key, crlf)
	value = bytes.Trim(value, crlf)

	newData := &Schema{Key: key, Value: value, Timestamp: time.Now()}
	c.tree.Insert(key, value)

	return newData, nil
}

// Get will get value from db
func (c *commanderBst) Get(command, key []byte) (*Schema, error) {
	lock.Lock()
	defer lock.Unlock()

	_, ok := commands[string(command)]
	if !ok {
		return nil, errors.New(ErrorInvalidCommand)
	}

	// remove line feed and carnige return (13/10)/ CF/LF
	key = bytes.Trim(key, crlf)

	value := c.tree.Search(key)
	if value == nil {
		return nil, errors.New(ErrorEmptyValue)
	}

	return value, nil
}

func (c *commanderBst) Delete(command, key []byte) (*Schema, error) {
	lock.Lock()
	defer lock.Unlock()

	_, ok := commands[string(command)]
	if !ok {
		return nil, errors.New(ErrorInvalidCommand)
	}

	// remove line feed and carnige return (13/10)/ CF/LF
	key = bytes.Trim(key, crlf)

	c.tree.Delete(key)
	return nil, nil
}

func (c *commanderBst) Publish(topic string, command, value []byte) ([]byte, error) {
	return nil, nil
}
