package mutation

import (
	"math/rand"
	"reflect"
)

// BoolInvert inverts a boolean to its opposite value.
// Example: true -> false
type BoolInvert struct{}

// ID returns mutator's 3-digit ID.
func (m *BoolInvert) ID() string {
	return "010"
}

// Name returns the mutator's name.
func (m *BoolInvert) Name() string {
	return "Invert Boolean"
}

// Description returns the mutator's description.
func (m *BoolInvert) Description() string {
	return "Invert boolean type to its opposite value"
}

// CompatibleKinds returns the data types that the mutator can mutate.
func (m *BoolInvert) CompatibleKinds() []reflect.Kind {
	return []reflect.Kind{reflect.Bool}
}

// Mutate returns a boolean's opposite value.
// Example: true -> false
func (m *BoolInvert) Mutate(subject reflect.Value, r *rand.Rand) interface{} {
	return !subject.Bool()
}
