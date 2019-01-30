package storage

import (
	"errors"
	"time"

	"github.com/Bhinneka/kece"
)

type hashMap struct {
	db map[string]*kece.Schema
}

// NewHashMap init new hashmap data store
func NewHashMap() kece.DataStructure {
	ds := new(hashMap)
	ds.db = make(map[string]*kece.Schema)
	return ds
}

// Insert new data to storage with new key and value
func (h *hashMap) Insert(key, value []byte) *kece.Schema {
	newData := &kece.Schema{Key: key, Value: value, Timestamp: time.Now()}
	h.db[string(key)] = newData
	return newData
}

// Search data based on key
func (h *hashMap) Search(key []byte) (*kece.Schema, error) {
	value, ok := h.db[string(key)]
	if !ok {
		return nil, errors.New(kece.ErrorEmptyValue)
	}
	return value, nil
}

// Delete data based on key
func (h *hashMap) Delete(key []byte) error {
	_, ok := h.db[string(key)]
	if !ok {
		return errors.New(kece.ErrorEmptyValue)
	}
	delete(h.db, string(key))
	return nil
}
