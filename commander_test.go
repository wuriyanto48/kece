package kece

import (
	"bytes"
	"testing"
)

func TestCommander(t *testing.T) {
	dataStructures := []DataStructure{NewHashMap(), NewBST()}

	for _, ds := range dataStructures {
		cmd := NewCommander(ds)

		t.Run("should success SET new value to db", func(t *testing.T) {
			key := []byte("1")
			value := []byte("wuriyanto")
			command := []byte("SET")

			expectedValue := []byte("wuriyanto")

			newValue, err := cmd.Set(command, key, value)

			if err != nil {
				t.Error("command is not valid")
			}

			if !bytes.Equal(newValue.Value, expectedValue) {
				t.Error("new value is not equal to value")
			}
		})

		t.Run("should success GET value from db", func(t *testing.T) {
			key := []byte("1")
			expectedValue := []byte("wuriyanto")
			command := []byte("GET")

			value, err := cmd.Get(command, key)

			if err != nil {
				t.Error(err.Error())
			}

			if !bytes.Equal(value.Value, expectedValue) {
				t.Error("value is not equal to expected value")
			}
		})

		t.Run("should error SET new value to db with invalid command", func(t *testing.T) {
			key := []byte("1")
			value := []byte("wuriyanto")
			command := []byte("ET")

			newValue, err := cmd.Set(command, key, value)

			if err == nil {
				t.Error("command should invalid")
			}

			if newValue != nil {
				t.Error("new value should be nil")
			}
		})

		t.Run("should error GET value from db with invalid command", func(t *testing.T) {
			key := []byte("1")
			command := []byte("ET")

			value, err := cmd.Get(command, key)

			if err == nil {
				t.Error("command is not valid")
			}

			if value != nil {
				t.Error("value should be nil")
			}
		})
	}
}
