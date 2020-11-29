package mutation_test

import (
	"reflect"
	"testing"

	"gitlab.com/michenriksen/jdam/pkg/jdam/mutation"
)

func TestLDAPInjectionID(t *testing.T) {
	m := &mutation.LDAPInjection{}
	if m.ID() != "014" {
		t.Errorf("Expected ID to be 014 but was %v", m.ID())
	}
}

func TestLDAPInjectionName(t *testing.T) {
	m := &mutation.LDAPInjection{}
	if m.Name() != "LDAP Injection" {
		t.Errorf("Expected Name to be LDAP Injection but was %v", m.Name())
	}
}

func TestLDAPInjectionCompatibleKinds(t *testing.T) {
	m := &mutation.LDAPInjection{}
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

func TestLDAPInjectionMutateInt(t *testing.T) {
	m := &mutation.LDAPInjection{}
	s := 1337
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestLDAPInjectionMutateFloat(t *testing.T) {
	m := &mutation.LDAPInjection{}
	s := 13.37
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestLDAPInjectionMutateString(t *testing.T) {
	m := &mutation.LDAPInjection{}
	s := "hello"
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestLDAPInjectionMutateNil(t *testing.T) {
	m := &mutation.LDAPInjection{}
	s := error(nil)
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}
