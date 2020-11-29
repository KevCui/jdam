package mutation

import (
	"bytes"
	"math/rand"
	"reflect"
)

// ByteSequenceRepeat repeats a random sequence of bytes a random amount of times.
// Example: Hello -> Helelelelelelelelelelelelelelelelelelelello
type ByteSequenceRepeat struct{}

// ID returns mutator's 3-digit ID.
func (m *ByteSequenceRepeat) ID() string {
	return "009"
}

// Name returns the mutator's name.
func (m *ByteSequenceRepeat) Name() string {
	return "Repeat Byte Sequence"
}

// Description returns the mutator's description.
func (m *ByteSequenceRepeat) Description() string {
	return "Repeat a random length byte sequence at random position a random amount of times"
}

// CompatibleKinds returns the data types that the mutator can mutate.
func (m *ByteSequenceRepeat) CompatibleKinds() []reflect.Kind {
	return []reflect.Kind{reflect.String}
}

// Mutate returns a string with a random sequence of bytes repeated a random amount of times
// Example: Hello -> Helelelelelelelelelelelelelelelelelelelello
func (m *ByteSequenceRepeat) Mutate(subject reflect.Value, r *rand.Rand) interface{} {
	subjectBytes := []byte(subject.String())
	rPos := r.Intn(len(subjectBytes))
	rLen := r.Intn(len(subjectBytes)) - rPos
	if rLen < 1 {
		rLen = 1
	}
	seq := subjectBytes[rPos : rPos+rLen]
	rRep := r.Intn(32)
	rSeq := bytes.Repeat(seq, rRep)
	return string(append(subjectBytes[:rPos], append(rSeq, subjectBytes[rPos:]...)...))
}
