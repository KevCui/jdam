package mutation

import (
	"math/rand"
	"reflect"
)

// ByteRepeat repeats a byte at a random position a random amount of times.
// Example: Hello -> HHHHHHHHHHHHHHHHHHHHHHHello
type ByteRepeat struct{}

// ID returns mutator's 3-digit ID.
func (m *ByteRepeat) ID() string {
	return "008"
}

// Name returns the mutator's name.
func (m *ByteRepeat) Name() string {
	return "Repeat Byte"
}

// Description returns the mutator's description.
func (m *ByteRepeat) Description() string {
	return "Repeat a byte at random position a random amount of times"
}

// CompatibleKinds returns the data types that the mutator can mutate.
func (m *ByteRepeat) CompatibleKinds() []reflect.Kind {
	return []reflect.Kind{reflect.String}
}

// Mutate returns a string with a random byte repeated a random amount of times
// Example: Hello -> HHHHHHHHHHHHHHHHHHHHHHHello
func (m *ByteRepeat) Mutate(subject reflect.Value, r *rand.Rand) interface{} {
	bytes := []byte(subject.String())
	if len(bytes) == 0 {
		return subject.String()
	}
	rPos := r.Intn(len(bytes))
	b := bytes[rPos]
	rRep := r.Intn(126) + 2
	rb := make([]byte, rRep)
	for i := range rb {
		rb[i] = b
	}
	return string(append(bytes[:rPos], append(rb, bytes[rPos:]...)...))
}
