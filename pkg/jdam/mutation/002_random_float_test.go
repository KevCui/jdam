package mutation_test

import (
	"reflect"
	"testing"

	"gitlab.com/michenriksen/jdam/pkg/jdam/mutation"
)

func TestRandomFloatID(t *testing.T) {
	m := &mutation.RandomFloat{}
	if m.ID() != "002" {
		t.Errorf("Expected ID to be 002 but was %v", m.ID())
	}
}

func TestRandomFloatName(t *testing.T) {
	m := &mutation.RandomFloat{}
	if m.Name() != "Random Float" {
		t.Errorf("Expected Name to be Random Float but was %v", m.Name())
	}
}

func TestRandomFloatCompatibleKinds(t *testing.T) {
	m := &mutation.RandomFloat{}
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

func TestRandomFloatMutateInt(t *testing.T) {
	m := &mutation.RandomFloat{}
	s := 1337
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestRandomFloatMutateFloat(t *testing.T) {
	m := &mutation.RandomFloat{}
	s := 13.37
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}
