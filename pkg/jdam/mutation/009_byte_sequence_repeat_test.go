package mutation_test

import (
	"reflect"
	"testing"

	"gitlab.com/michenriksen/jdam/pkg/jdam/mutation"
)

func TestByteSequenceRepeatID(t *testing.T) {
	m := &mutation.ByteSequenceRepeat{}
	if m.ID() != "009" {
		t.Errorf("Expected ID to be 009 but was %v", m.ID())
	}
}

func TestByteSequenceRepeatName(t *testing.T) {
	m := &mutation.ByteSequenceRepeat{}
	if m.Name() != "Repeat Byte Sequence" {
		t.Errorf("Expected Name to be Repeat Byte Sequence but was %v", m.Name())
	}
}

func TestByteSequenceRepeatCompatibleKinds(t *testing.T) {
	m := &mutation.ByteSequenceRepeat{}
	if mutatorIsCompatibleWith(m, reflect.Int) {
		t.Error("Expected mutator to not be compatible with reflect.Int")
	}
	if mutatorIsCompatibleWith(m, reflect.Float64) {
		t.Error("Expected mutator to not be compatible with reflect.Float64")
	}
	if mutatorIsCompatibleWith(m, reflect.Invalid) {
		t.Error("Expected mutator to not be compatible with reflect.Invalid")
	}
	if !mutatorIsCompatibleWith(m, reflect.String) {
		t.Error("Expected mutator to be compatible with reflect.String")
	}
}

func TestByteSequenceRepeatMutate(t *testing.T) {
	m := &mutation.ByteSequenceRepeat{}
	s := "hello"
	rs := reflect.ValueOf(s)
	mu := m.Mutate(rs, getRand())
	if len(mu.(string)) < len(s) {
		t.Errorf("Expected length of mutation to be larger than %d but was %d", len(s), len(mu.(string)))
	}
}
