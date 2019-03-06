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
		name        string
		args        string
		wantValue   string
		wantExpired int
		wantError   bool
	}{
		{
			name:      "Testcase #1: Positive (store json string)",
			args:      `{"field": "this is value from field a"}`,
			wantError: false,
			wantValue: `{"field": "this is value from field a"}`,
		},
		{
			name:      "Testcase #2: Positive (store a string with spaces)",
			args:      `"this is value"`,
			wantError: false,
			wantValue: `this is value`,
		},
		{
			name:      "Testcase #3: Positive (store a string with spaces)",
			args:      `'this is value'`,
			wantError: false,
			wantValue: `this is value`,
		},
		{
			name:        "Testcase #4: Positive (store a string with expired value)",
			args:        `'this is value with lifetime 20 seconds' 20`,
			wantError:   false,
			wantValue:   `this is value with lifetime 20 seconds`,
			wantExpired: 20,
		},
		{
			name:      "Testcase #5: Negative (invalid sign)",
			args:      `"this is value'`,
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, expired, err := processingValue(tt.args)
			if (err != nil) != tt.wantError {
				t.Errorf("error: processingValue() = %v, want %v", err, tt.wantError)
			}

			if err == nil && value != tt.wantValue {
				t.Errorf("value: processingValue() = %v, want %v", value, tt.wantValue)
			}

			if err == nil && expired != tt.wantExpired {
				t.Errorf("expired: processingValue() = %v, want %v", expired, tt.wantExpired)
			}
		})
	}
}
