package kece

// Commander interface
type Commander interface {
	Set(value []byte) ([]byte, error)
	Get(value []byte) ([]byte, error)
	Publish(value []byte) ([]byte, error)
}

type commander struct {
}
