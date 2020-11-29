package mutation_test

import (
	"reflect"
	"testing"

	"gitlab.com/michenriksen/jdam/pkg/jdam/mutation"
)

func TestXPATHInjectionID(t *testing.T) {
	m := &mutation.XPATHInjection{}
	if m.ID() != "013" {
		t.Errorf("Expected ID to be 013 but was %v", m.ID())
	}
}

func TestXPATHInjectionName(t *testing.T) {
	m := &mutation.XPATHInjection{}
	if m.Name() != "XPATH Injection" {
		t.Errorf("Expected Name to be XPATH Injection but was %v", m.Name())
	}
}

func TestXPATHInjectionCompatibleKinds(t *testing.T) {
	m := &mutation.XPATHInjection{}
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

func TestXPATHInjectionMutateInt(t *testing.T) {
	m := &mutation.XPATHInjection{}
	s := 1337
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestXPATHInjectionMutateFloat(t *testing.T) {
	m := &mutation.XPATHInjection{}
	s := 13.37
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestXPATHInjectionMutateString(t *testing.T) {
	m := &mutation.XPATHInjection{}
	s := "hello"
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestXPATHInjectionMutateNil(t *testing.T) {
	m := &mutation.XPATHInjection{}
	s := error(nil)
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}
