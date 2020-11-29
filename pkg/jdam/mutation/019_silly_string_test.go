package mutation_test

import (
	"reflect"
	"testing"

	"gitlab.com/michenriksen/jdam/pkg/jdam/mutation"
)

func TestSillyStringID(t *testing.T) {
	m := &mutation.SillyString{}
	if m.ID() != "019" {
		t.Errorf("Expected ID to be 019 but was %v", m.ID())
	}
}

func TestSillyStringName(t *testing.T) {
	m := &mutation.SillyString{}
	if m.Name() != "Silly String" {
		t.Errorf("Expected Name to be Silly String but was %v", m.Name())
	}
}

func TestSillyStringCompatibleKinds(t *testing.T) {
	m := &mutation.SillyString{}
	if !mutatorIsCompatibleWith(m, reflect.Int) {
		t.Error("Expected mutator to be compatible with reflect.Int")
	}
	if !mutatorIsCompatibleWith(m, reflect.Float64) {
		t.Error("Expected mutator to be compatible with reflect.Float64")
	}
	if !mutatorIsCompatibleWith(m, reflect.Invalid) {
		t.Error("Expected mutator to be compatible with reflect.Invalid")
	}
	if !mutatorIsCompatibleWith(m, reflect.String) {
		t.Error("Expected mutator to be compatible with reflect.String")
	}
}

func TestSillyStringMutateInt(t *testing.T) {
	m := &mutation.SillyString{}
	s := 1337
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestSillyStringMutateFloat(t *testing.T) {
	m := &mutation.SillyString{}
	s := 13.37
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestSillyStringMutateString(t *testing.T) {
	m := &mutation.SillyString{}
	s := "hello"
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestSillyStringMutateNil(t *testing.T) {
	m := &mutation.SillyString{}
	s := error(nil)
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}
