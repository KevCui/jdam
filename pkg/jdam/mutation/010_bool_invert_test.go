package mutation_test

import (
	"reflect"
	"testing"

	"gitlab.com/michenriksen/jdam/pkg/jdam/mutation"
)

func TestBoolInvertID(t *testing.T) {
	m := &mutation.BoolInvert{}
	if m.ID() != "010" {
		t.Errorf("Expected ID to be 010 but was %v", m.ID())
	}
}

func TestBoolInvertName(t *testing.T) {
	m := &mutation.BoolInvert{}
	if m.Name() != "Invert Boolean" {
		t.Errorf("Expected Name to be Invert Boolean but was %v", m.Name())
	}
}

func TestBoolInvertCompatibleKinds(t *testing.T) {
	m := &mutation.BoolInvert{}
	if !mutatorIsCompatibleWith(m, reflect.Bool) {
		t.Error("Expected mutator to not be compatible with reflect.Bool")
	}
	if mutatorIsCompatibleWith(m, reflect.Int) {
		t.Error("Expected mutator to not be compatible with reflect.Int")
	}
	if mutatorIsCompatibleWith(m, reflect.Float64) {
		t.Error("Expected mutator to not be compatible with reflect.Float64")
	}
	if mutatorIsCompatibleWith(m, reflect.Invalid) {
		t.Error("Expected mutator to not be compatible with reflect.Invalid")
	}
	if mutatorIsCompatibleWith(m, reflect.String) {
		t.Error("Expected mutator to not be compatible with reflect.String")
	}
}

func TestBoolInvertMutateTrue(t *testing.T) {
	m := &mutation.BoolInvert{}
	s := true
	rs := reflect.ValueOf(s)
	mu := m.Mutate(rs, getRand())
	if mu.(bool) != false {
		t.Errorf("Expected mutation to be false but was %v", mu)
	}
}

func TestBoolInvertMutateFalse(t *testing.T) {
	m := &mutation.BoolInvert{}
	s := false
	rs := reflect.ValueOf(s)
	mu := m.Mutate(rs, getRand())
	if mu.(bool) != true {
		t.Errorf("Expected mutation to be true but was %v", mu)
	}
}
