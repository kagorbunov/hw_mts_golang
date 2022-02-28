package contract

import "time"

type Signature interface {
	Date() time.Time
	Size() string
	// Name of file
	Name() string
	Signature() []byte
	SignatureByte() []byte
	Equals (s Signature) bool
}
