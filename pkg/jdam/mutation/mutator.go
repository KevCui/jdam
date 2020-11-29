package mutation

import (
	"math/rand"
	"reflect"
	"unicode/utf8"
)

// Mutator represents an abstract mutator.
type Mutator interface {
	// ID returns the mutator's 3-digit ID.
	ID() string
	// Name returns the mutator's name.
	Name() string
	// Description returns the mutator's description.
	Description() string
	// CompatibleKinds returns the data types that the mutator can mutate.
	CompatibleKinds() []reflect.Kind
	// Mutate mutates the given subject and returns the result.
	Mutate(subject reflect.Value, r *rand.Rand) interface{}
}

// MutatorList is a list of mutators and is used when constructing a
// new Fuzzer.
type MutatorList []Mutator

// MutatorMap is a map of mutators sorted by their compatible data types.
type MutatorMap map[reflect.Kind][]Mutator

// Mutators is a list of all available mutators
var Mutators = MutatorList{
	new(RandomInt),
	new(RandomFloat),
	new(NumericToNegative),
	new(IntegerOverflow),
	new(ByteSwap),
	new(ByteInsert),
	new(ByteDrop),
	new(ByteRepeat),
	new(ByteSequenceRepeat),
	new(BoolInvert),
	new(SQLInjection),
	new(CommandInjection),
	new(XPATHInjection),
	new(LDAPInjection),
	new(NoSQLInjection),
	new(FormatStringInjection),
	new(LocalFileInclusion),
	new(TemplateInjection),
	new(SillyString),
	new(SillyObject),
}

func randomUTF8Byte(r *rand.Rand) byte {
	buf := make([]byte, 1)
	r.Read(buf)
	if !utf8.ValidString(string(buf)) {
		return randomUTF8Byte(r)
	}
	return buf[0]
}
