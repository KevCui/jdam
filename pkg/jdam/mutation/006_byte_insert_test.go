package mutation_test

import (
	"reflect"
	"testing"

	"gitlab.com/michenriksen/jdam/pkg/jdam/mutation"
)

func TestByteInsertID(t *testing.T) {
	m := &mutation.ByteInsert{}
	if m.ID() != "006" {
		t.Errorf("Expected ID to be 006 but was %v", m.ID())
	}
}

func TestByteInsertName(t *testing.T) {
	m := &mutation.ByteInsert{}
	if m.Name() != "Insert Byte" {
		t.Errorf("Expected Name to be Insert Byte but was %v", m.Name())
	}
}

func TestByteInsertCompatibleKinds(t *testing.T) {
	m := &mutation.ByteInsert{}
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

func TestByteInsertMutate(t *testing.T) {
	m := &mutation.ByteInsert{}
	s := "hello"
	rs := reflect.ValueOf(s)
	mu := m.Mutate(rs, getRand())
	if len(mu.(string)) != len(s)+1 {
		t.Errorf("Expected length of mutation to be %d but was %d", len(s)+1, len(mu.(string)))
	}
}
