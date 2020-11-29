package mutation_test

import (
	"reflect"
	"testing"

	"gitlab.com/michenriksen/jdam/pkg/jdam/mutation"
)

func TestByteSwapID(t *testing.T) {
	m := &mutation.ByteSwap{}
	if m.ID() != "005" {
		t.Errorf("Expected ID to be 005 but was %v", m.ID())
	}
}

func TestByteSwapName(t *testing.T) {
	m := &mutation.ByteSwap{}
	if m.Name() != "Swap Byte" {
		t.Errorf("Expected Name to be Swap Byte but was %v", m.Name())
	}
}

func TestByteSwapCompatibleKinds(t *testing.T) {
	m := &mutation.ByteSwap{}
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

func TestByteSwapMutate(t *testing.T) {
	m := &mutation.ByteSwap{}
	s := "hello"
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}
