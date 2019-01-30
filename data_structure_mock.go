package kece

import (
	"errors"
	"time"
)

type dataStructureMock struct {
	db map[string]*Schema
}

// newHashMap init new hashmap data store
func newStructureMock() DataStructure {
	ds := new(dataStructureMock)
	ds.db = make(map[string]*Schema)
	return ds
}

// Insert new data to storage with new key and value
func (h *dataStructureMock) Insert(key, value []byte) *Schema {
	newData := &Schema{Key: key, Value: value, Timestamp: time.Now()}
	h.db[string(key)] = newData
	return newData
}

// Search data based on key
func (h *dataStructureMock) Search(key []byte) (*Schema, error) {
	value, ok := h.db[string(key)]
	if !ok {
		return nil, errors.New(ErrorEmptyValue)
	}
	return value, nil
}

// Delete data based on key
func (h *dataStructureMock) Delete(key []byte) error {
	_, ok := h.db[string(key)]
	if !ok {
		return errors.New(ErrorEmptyValue)
	}
	delete(h.db, string(key))
	return nil
}
