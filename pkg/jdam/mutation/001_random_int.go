package mutation

import (
	"math/rand"
	"reflect"
)

// RandomInt replaces a numeric type with a random integer.
type RandomInt struct{}

// ID returns mutator's 3-digit ID.
func (m *RandomInt) ID() string {
	return "001"
}

// Name returns the mutator's name.
func (m *RandomInt) Name() string {
	return "Random Integer"
}

// Description returns the mutator's description.
func (m *RandomInt) Description() string {
	return "Replace numeric type with random integer"
}

// CompatibleKinds returns the data types that the mutator can mutate.
func (m *RandomInt) CompatibleKinds() []reflect.Kind {
	return []reflect.Kind{reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Invalid}
}

// Mutate returns a random integer.
func (m *RandomInt) Mutate(subject reflect.Value, r *rand.Rand) interface{} {
	return uint64(r.Uint32())<<32 | uint64(r.Uint32())
}
