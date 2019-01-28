package kece

import (
	"testing"
)

func TestCommander(t *testing.T) {
	db := make(map[string]*Schema)
	cmd := NewCommander(db)

	t.Run("should success SET new value to db", func(t *testing.T) {
		key := []byte("1")
		value := []byte("wuriyanto")
		command := []byte("SET")

		newValue, err := cmd.Set(command, key, value)

		if err != nil {
			t.Error(err.Error())
		}

		if string(newValue.Value) != "wuriyanto" {
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

		if string(value.Value) != string(expectedValue) {
			t.Error("value is not equal to expected value")
		}
	})
}
