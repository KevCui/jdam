package mutation_test

import (
	"reflect"
	"testing"

	"gitlab.com/michenriksen/jdam/pkg/jdam/mutation"
)

func TestFormatStringInjectionID(t *testing.T) {
	m := &mutation.FormatStringInjection{}
	if m.ID() != "016" {
		t.Errorf("Expected ID to be 016 but was %v", m.ID())
	}
}

func TestFormatStringInjectionName(t *testing.T) {
	m := &mutation.FormatStringInjection{}
	if m.Name() != "Format String Injection" {
		t.Errorf("Expected Name to be Format String Injection but was %v", m.Name())
	}
}

func TestFormatStringInjectionCompatibleKinds(t *testing.T) {
	m := &mutation.FormatStringInjection{}
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

func TestFormatStringInjectionMutateInt(t *testing.T) {
	m := &mutation.FormatStringInjection{}
	s := 1337
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestFormatStringInjectionMutateFloat(t *testing.T) {
	m := &mutation.FormatStringInjection{}
	s := 13.37
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestFormatStringInjectionMutateString(t *testing.T) {
	m := &mutation.FormatStringInjection{}
	s := "hello"
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestFormatStringInjectionMutateNil(t *testing.T) {
	m := &mutation.FormatStringInjection{}
	s := error(nil)
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}
