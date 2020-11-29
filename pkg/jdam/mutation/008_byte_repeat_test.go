package mutation_test

import (
	"reflect"
	"testing"

	"gitlab.com/michenriksen/jdam/pkg/jdam/mutation"
)

func TestByteRepeatID(t *testing.T) {
	m := &mutation.ByteRepeat{}
	if m.ID() != "008" {
		t.Errorf("Expected ID to be 008 but was %v", m.ID())
	}
}

func TestByteRepeatName(t *testing.T) {
	m := &mutation.ByteRepeat{}
	if m.Name() != "Repeat Byte" {
		t.Errorf("Expected Name to be Repeat Byte but was %v", m.Name())
	}
}

func TestByteRepeatCompatibleKinds(t *testing.T) {
	m := &mutation.ByteRepeat{}
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

func TestByteRepeatMutate(t *testing.T) {
	m := &mutation.ByteRepeat{}
	s := "hello"
	rs := reflect.ValueOf(s)
	mu := m.Mutate(rs, getRand())
	if len(mu.(string)) < len(s) {
		t.Errorf("Expected length of mutation to be larger than %d but was %d", len(s), len(mu.(string)))
	}
}
