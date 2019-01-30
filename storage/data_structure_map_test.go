package storage

import (
	"bytes"
	"testing"
)

func TestHashMapStorage(t *testing.T) {
	cmd := NewHashMap()

	t.Run("should success Insert new value to db", func(t *testing.T) {
		key := []byte("1")
		value := []byte("wuriyanto")

		expectedValue := []byte("wuriyanto")

		newValue := cmd.Insert(key, value)

		if !bytes.Equal(newValue.Value, expectedValue) {
			t.Error("new value is not equal to value")
		}
	})

	t.Run("should success Search value from db", func(t *testing.T) {
		key := []byte("1")
		expectedValue := []byte("wuriyanto")

		value, err := cmd.Search(key)

		if err != nil {
			t.Error(err.Error())
		}

		if !bytes.Equal(value.Value, expectedValue) {
			t.Error("value is not equal to expected value")
		}
	})

	t.Run("should success DELETE value from db", func(t *testing.T) {
		key := []byte("1")

		err := cmd.Delete(key)

		if err != nil {
			t.Error("command is not valid")
		}

		_, err = cmd.Search(key)

		if err == nil {
			t.Error(err.Error())
		}
	})
}
