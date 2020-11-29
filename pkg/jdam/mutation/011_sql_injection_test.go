package mutation_test

import (
	"reflect"
	"testing"

	"gitlab.com/michenriksen/jdam/pkg/jdam/mutation"
)

func TestSQLInjectionID(t *testing.T) {
	m := &mutation.SQLInjection{}
	if m.ID() != "011" {
		t.Errorf("Expected ID to be 011 but was %v", m.ID())
	}
}

func TestSQLInjectionName(t *testing.T) {
	m := &mutation.SQLInjection{}
	if m.Name() != "SQL Injection" {
		t.Errorf("Expected Name to be SQL Injection but was %v", m.Name())
	}
}

func TestSQLInjectionCompatibleKinds(t *testing.T) {
	m := &mutation.SQLInjection{}
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

func TestSQLInjectionMutateInt(t *testing.T) {
	m := &mutation.SQLInjection{}
	s := 1337
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestSQLInjectionMutateFloat(t *testing.T) {
	m := &mutation.SQLInjection{}
	s := 13.37
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestSQLInjectionMutateString(t *testing.T) {
	m := &mutation.SQLInjection{}
	s := "hello"
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestSQLInjectionMutateNil(t *testing.T) {
	m := &mutation.SQLInjection{}
	s := error(nil)
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}
