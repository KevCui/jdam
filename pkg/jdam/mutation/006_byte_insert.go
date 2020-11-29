package mutation

import (
	"math/rand"
	"reflect"
)

// ByteInsert inserts a random byte at a random position.
// Example: Hello -> H^ello
type ByteInsert struct{}

// ID returns mutator's 3-digit ID.
func (m *ByteInsert) ID() string {
	return "006"
}

// Name returns the mutator's name.
func (m *ByteInsert) Name() string {
	return "Insert Byte"
}

// Description returns the mutator's description.
func (m *ByteInsert) Description() string {
	return "Insert a random byte at random position"
}

// CompatibleKinds returns the data types that the mutator can mutate.
func (m *ByteInsert) CompatibleKinds() []reflect.Kind {
	return []reflect.Kind{reflect.String}
}

// Mutate returns a string with a random byte inserted at random position.
// Example: Hello -> H^ello
func (m *ByteInsert) Mutate(subject reflect.Value, r *rand.Rand) interface{} {
	bytes := []byte(subject.String())
	if len(bytes) == 0 {
		return subject.String()
	}
	rPos := r.Intn(len(bytes))
	return string(append(bytes[:rPos], append([]byte{randomUTF8Byte(r)}, bytes[rPos:]...)...))
}
