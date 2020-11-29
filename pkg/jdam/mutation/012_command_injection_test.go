package mutation_test

import (
	"reflect"
	"testing"

	"gitlab.com/michenriksen/jdam/pkg/jdam/mutation"
)

func TestCommandInjectionID(t *testing.T) {
	m := &mutation.CommandInjection{}
	if m.ID() != "012" {
		t.Errorf("Expected ID to be 012 but was %v", m.ID())
	}
}

func TestCommandInjectionName(t *testing.T) {
	m := &mutation.CommandInjection{}
	if m.Name() != "Command Injection" {
		t.Errorf("Expected Name to be Command Injection but was %v", m.Name())
	}
}

func TestCommandInjectionCompatibleKinds(t *testing.T) {
	m := &mutation.CommandInjection{}
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

func TestCommandInjectionMutateInt(t *testing.T) {
	m := &mutation.CommandInjection{}
	s := 1337
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestCommandInjectionMutateFloat(t *testing.T) {
	m := &mutation.CommandInjection{}
	s := 13.37
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestCommandInjectionMutateString(t *testing.T) {
	m := &mutation.CommandInjection{}
	s := "hello"
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestCommandInjectionMutateNil(t *testing.T) {
	m := &mutation.CommandInjection{}
	s := error(nil)
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}
