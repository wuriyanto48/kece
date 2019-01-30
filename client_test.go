package kece

import (
	"testing"
)

func TestClient(t *testing.T) {

	client := &Client{ID: "001"}

	cm := &ClientMessage{Client: client}

	t.Run("should success with command GET, DEL, or AUTH with valid message", func(t *testing.T) {
		cm.Message = []byte("GET 1")

		err := cm.ValidateMessage()

		if err != nil {
			t.Errorf("error validate client message with GET command %s", err.Error())
		}
	})

	t.Run("should error with command GET, DEL, or AUTH with invalid message", func(t *testing.T) {
		cm.Message = []byte("DEL 1 bla")

		err := cm.ValidateMessage()

		if err == nil {
			t.Errorf("error validate client message with DEL command %s", err.Error())
		}
	})

	t.Run("should success with command SET with valid message", func(t *testing.T) {
		cm.Message = []byte("SET 1 wury")

		err := cm.ValidateMessage()

		if err != nil {
			t.Errorf("error validate client message with SET command %s", err.Error())
		}
	})

	t.Run("should error with command SET with invalid message", func(t *testing.T) {
		cm.Message = []byte("SET 1 wury yanto")

		err := cm.ValidateMessage()

		if err == nil {
			t.Errorf("error validate client message with SET command %s", err.Error())
		}
	})
}
