package mutation_test

import (
	"reflect"
	"testing"

	"gitlab.com/michenriksen/jdam/pkg/jdam/mutation"
)

func TestRandomIntID(t *testing.T) {
	m := &mutation.RandomInt{}
	if m.ID() != "001" {
		t.Errorf("Expected ID to be 001 but was %v", m.ID())
	}
}

func TestRandomIntName(t *testing.T) {
	m := &mutation.RandomInt{}
	if m.Name() != "Random Integer" {
		t.Errorf("Expected Name to be Random Integer but was %v", m.Name())
	}
}

func TestRandomIntCompatibleKinds(t *testing.T) {
	m := &mutation.RandomInt{}
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

func TestRandomIntMutateInt(t *testing.T) {
	m := &mutation.RandomInt{}
	s := 1337
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestRandomIntMutateFloat(t *testing.T) {
	m := &mutation.RandomInt{}
	s := 13.37
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}
