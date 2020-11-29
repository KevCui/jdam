package mutation_test

import (
	"reflect"
	"testing"

	"gitlab.com/michenriksen/jdam/pkg/jdam/mutation"
)

func TestIntegerOverflowID(t *testing.T) {
	m := &mutation.IntegerOverflow{}
	if m.ID() != "004" {
		t.Errorf("Expected ID to be 004 but was %v", m.ID())
	}
}

func TestIntegerOverflowName(t *testing.T) {
	m := &mutation.IntegerOverflow{}
	if m.Name() != "Integer Overflow" {
		t.Errorf("Expected Name to be Integer Overflow but was %v", m.Name())
	}
}

func TestIntegerOverflowCompatibleKinds(t *testing.T) {
	m := &mutation.IntegerOverflow{}
	if !mutatorIsCompatibleWith(m, reflect.Int) {
		t.Error("Expected mutator to be compatible with reflect.Int")
	}
	if !mutatorIsCompatibleWith(m, reflect.Float64) {
		t.Error("Expected mutator to be compatible with reflect.Float64")
	}
	if !mutatorIsCompatibleWith(m, reflect.Invalid) {
		t.Error("Expected mutator to be compatible with reflect.Invalid")
	}
	if mutatorIsCompatibleWith(m, reflect.String) {
		t.Error("Expected mutator to not be compatible with reflect.String")
	}
}

func TestIntegerOverflowMutateInt(t *testing.T) {
	m := &mutation.IntegerOverflow{}
	s := 1337
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestIntegerOverflowMutateFloat(t *testing.T) {
	m := &mutation.IntegerOverflow{}
	s := 13.37
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}
