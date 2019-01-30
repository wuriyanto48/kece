package kece

// DataStructure abstract interface
type DataStructure interface {
	Insert(key, value []byte) *Schema
	Search(key []byte) (*Schema, error)
	Delete(key []byte) error
}
