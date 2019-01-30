package kece

import (
	"time"
)

// Schema database
type Schema struct {
	Key       []byte
	Value     []byte
	Timestamp time.Time
}
