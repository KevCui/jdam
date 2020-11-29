package mutation

import (
	"math/rand"
	"reflect"
)

// ByteSwap swaps a byte at a random position with a random byte.
// Example: Hello -> He`lo
type ByteSwap struct{}

// ID returns mutator's 3-digit ID.
func (m *ByteSwap) ID() string {
	return "005"
}

// Name returns the mutator's name.
func (m *ByteSwap) Name() string {
	return "Swap Byte"
}

// Description returns the mutator's description.
func (m *ByteSwap) Description() string {
	return "Swap a byte at random position with random byte"
}

// CompatibleKinds returns the data types that the mutator can mutate.
func (m *ByteSwap) CompatibleKinds() []reflect.Kind {
	return []reflect.Kind{reflect.String}
}

// Mutate returns a string with a random byte swapped.
// Example: Hello -> He`lo
func (m *ByteSwap) Mutate(subject reflect.Value, r *rand.Rand) interface{} {
	bytes := []byte(subject.String())
	if len(bytes) == 0 {
		return subject.String()
	}
	bytes[r.Intn(len(bytes))] = randomUTF8Byte(r)
	return string(bytes)
}
