package mutation

import (
	"math/rand"
	"reflect"
)

// ByteDrop drops a byte at a random position.
// Example: Hello -> Helo
type ByteDrop struct{}

// ID returns mutator's 3-digit ID.
func (m *ByteDrop) ID() string {
	return "007"
}

// Name returns the mutator's name.
func (m *ByteDrop) Name() string {
	return "Drop Byte"
}

// Description returns the mutator's description.
func (m *ByteDrop) Description() string {
	return "Drop a byte at random position"
}

// CompatibleKinds returns the data types that the mutator can mutate.
func (m *ByteDrop) CompatibleKinds() []reflect.Kind {
	return []reflect.Kind{reflect.String}
}

// Mutate returns a string with a byte at a random position dropped.
// Example: Hello -> Helo
func (m *ByteDrop) Mutate(subject reflect.Value, r *rand.Rand) interface{} {
	bytes := []byte(subject.String())
	if len(bytes) == 0 {
		return subject.String()
	}
	rPos := r.Intn(len(bytes))
	return string(append(bytes[:rPos], bytes[rPos+1:]...))
}
