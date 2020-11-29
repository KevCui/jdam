package mutation

import (
	"math/rand"
	"reflect"
)

var intOverflowPayloads = []float64{
	-0,
	-0.999999999999999,
	-1,
	-1.97684995314e+16,
	-128,
	-129,
	-139333426276771806651771,
	-2147483648,
	-2147483649,
	-32768,
	-32769,
	-9223372036854775808,
	-9223372036854775809,
	-9223372036854775809,
	-999999999999999,
	0x10000,
	0x100000,
	0x3fffffff,
	0x481b49d0f8d5a3e7f821066157c37c,
	0x7ffffffe,
	0x7fffffff,
	0x80000000,
	0xfffffffe,
	0xffffffff,
	1.79769313486e+308,
	1.79769313486e+308,
	16649142472222295162770764775,
	18446744073709551615,
	18446744073709551616,
	2.07564741538e+16,
	2.22507385851e-308,
	2139095040,
	2147483647,
	2147483648,
	3.38800266804e+16,
	32767,
	32768,
	4294967295,
	4294967296,
	65536,
	9223372036854775807,
	9223372036854775807,
	9223372036854775808,
	9223372036854775808,
	999999999999999,
}

// IntegerOverflow replaces values with random integer overflow payloads.
// For more information on integer overflows: https://en.wikipedia.org/wiki/Integer_overflow
// Example: 1337 -> 16649142472222295162770764775
type IntegerOverflow struct{}

// ID returns mutator's 3-digit ID.
func (m *IntegerOverflow) ID() string {
	return "004"
}

// Name returns the mutator's name.
func (m *IntegerOverflow) Name() string {
	return "Integer Overflow"
}

// Description returns the mutator's description.
func (m *IntegerOverflow) Description() string {
	return "Replace numeric type with random integer overflow payload"
}

// CompatibleKinds returns the data types that the mutator can mutate.
func (m *IntegerOverflow) CompatibleKinds() []reflect.Kind {
	return []reflect.Kind{reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Invalid}
}

// Mutate returns a random integer overflow payload.
func (m *IntegerOverflow) Mutate(subject reflect.Value, r *rand.Rand) interface{} {
	return intOverflowPayloads[r.Intn(len(intOverflowPayloads))]
}
