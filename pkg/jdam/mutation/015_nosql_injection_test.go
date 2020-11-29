package mutation_test

import (
	"reflect"
	"testing"

	"gitlab.com/michenriksen/jdam/pkg/jdam/mutation"
)

func TestNoSQLInjectionID(t *testing.T) {
	m := &mutation.NoSQLInjection{}
	if m.ID() != "015" {
		t.Errorf("Expected ID to be 015 but was %v", m.ID())
	}
}

func TestNoSQLInjectionName(t *testing.T) {
	m := &mutation.NoSQLInjection{}
	if m.Name() != "NoSQL Injection" {
		t.Errorf("Expected Name to be NoSQL Injection but was %v", m.Name())
	}
}

func TestNoSQLInjectionCompatibleKinds(t *testing.T) {
	m := &mutation.NoSQLInjection{}
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

func TestNoSQLInjectionMutateInt(t *testing.T) {
	m := &mutation.NoSQLInjection{}
	s := 1337
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestNoSQLInjectionMutateFloat(t *testing.T) {
	m := &mutation.NoSQLInjection{}
	s := 13.37
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestNoSQLInjectionMutateString(t *testing.T) {
	m := &mutation.NoSQLInjection{}
	s := "hello"
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestNoSQLInjectionMutateNil(t *testing.T) {
	m := &mutation.NoSQLInjection{}
	s := error(nil)
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}
