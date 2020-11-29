package mutation

import (
	"math/rand"
	"reflect"
)

// NumericToNegative replaces a numeric type with its negative equivalent.
// Example: 1337 -> -1337
type NumericToNegative struct{}

// ID returns mutator's 3-digit ID.
func (m *NumericToNegative) ID() string {
	return "003"
}

// Name returns the mutator's name.
func (m *NumericToNegative) Name() string {
	return "Numeric To Negative"
}

// Description returns the mutator's description.
func (m *NumericToNegative) Description() string {
	return "Replace numeric type with its negative equivalent"
}

// CompatibleKinds returns the data types that the mutator can mutate.
func (m *NumericToNegative) CompatibleKinds() []reflect.Kind {
	return []reflect.Kind{reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64}
}

// Mutate returns the numeric value's negative equivalent.
// Example: 1337 -> -1337
func (m *NumericToNegative) Mutate(subject reflect.Value, r *rand.Rand) interface{} {
	switch subject.Kind() {
	case reflect.Int:
		return -subject.Int()
	case reflect.Uint:
		return -subject.Uint()
	default:
		return -subject.Float()
	}
}
