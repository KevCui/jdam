package mutation

import (
	"math/rand"
	"reflect"
)

// RandomFloat replaces a numeric type with a random float.
type RandomFloat struct{}

// ID returns mutator's 3-digit ID.
func (m *RandomFloat) ID() string {
	return "002"
}

// Name returns the mutator's name.
func (m *RandomFloat) Name() string {
	return "Random Float"
}

// Description returns the mutator's description.
func (m *RandomFloat) Description() string {
	return "Replace numeric type with random float"
}

// CompatibleKinds returns the data types that the mutator can mutate.
func (m *RandomFloat) CompatibleKinds() []reflect.Kind {
	return []reflect.Kind{reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Invalid}
}

// Mutate returns a random float.
func (m *RandomFloat) Mutate(subject reflect.Value, r *rand.Rand) interface{} {
	return r.Float64()
}
