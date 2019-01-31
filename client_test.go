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
			t.Errorf("error validate client message with DEL command")
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
			t.Errorf("error validate client message with SET command")
		}
	})

	t.Run("should success with command SET with many spaces", func(t *testing.T) {
		cm.Message = []byte("SET         1 agung")

		err := cm.ValidateMessage()
		if err != nil {
			t.Errorf("error validate client message with SET command %s", err.Error())
		}
	})

	t.Run("should error with command SET with invalid args", func(t *testing.T) {
		cm.Message = []byte(`SET k "agung dwi p'`)

		err := cm.ValidateMessage()
		if err == nil {
			t.Errorf("error validate client message with SET command")
		}
	})

	t.Run("should success with command SET with JSON string", func(t *testing.T) {
		cm.Message = []byte(`SET adp {"id": 1, "name": "agung dwi prasetyo"}`)

		err := cm.ValidateMessage()
		if err != nil {
			t.Errorf("error validate client message with SET command for JSON data with spaces %s", err.Error())
		}

		if string(cm.Value) != `{"id": 1, "name": "agung dwi prasetyo"}` {
			t.Errorf("value is not equal")
		}
	})

	t.Run("should success with command SET with a string", func(t *testing.T) {
		cm.Message = []byte(`SET v "this is value stored to      kece"`)

		err := cm.ValidateMessage()
		if err != nil {
			t.Errorf("error validate client message with SET command for string data %s", err.Error())
		}

		if string(cm.Value) != `this is value stored to      kece` {
			t.Errorf("value is not equal")
		}
	})

	t.Run("should error with command SET with invalid value", func(t *testing.T) {
		cm.Message = []byte(`SET v "test'`)

		err := cm.ValidateMessage()
		if err == nil {
			t.Errorf("error validate client message with SET command for invalid value")
		}
	})
}

func TestIsValidValue(t *testing.T) {
	tests := []struct {
		name       string
		args       string
		wantValid  bool
		wantResult string
	}{
		{
			name:       "Testcase #1: Positive (store json string)",
			args:       `{"field": "this is value from field a"}`,
			wantValid:  true,
			wantResult: `{"field": "this is value from field a"}`,
		},
		{
			name:       "Testcase #2: Positive (store a string with spaces)",
			args:       `"this is value"`,
			wantValid:  true,
			wantResult: `this is value`,
		},
		{
			name:       "Testcase #3: Positive (store a string with spaces)",
			args:       `'this is value'`,
			wantValid:  true,
			wantResult: `this is value`,
		},
		{
			name:      "Testcase #4: Negative (invalid sign)",
			args:      `"this is value'`,
			wantValid: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid, res := isValidValue(tt.args)
			if isValid != tt.wantValid {
				t.Errorf("isValidValue() = %v, want %v", isValid, tt.wantValid)
			}
			if isValid && res != tt.wantResult {
				t.Errorf("isValidValue() = %v, want %v", res, tt.wantResult)
			}
		})
	}
}
