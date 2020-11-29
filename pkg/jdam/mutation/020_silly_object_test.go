package mutation_test

import (
	"reflect"
	"testing"

	"gitlab.com/michenriksen/jdam/pkg/jdam/mutation"
)

func TestSillyObjectID(t *testing.T) {
	m := &mutation.SillyObject{}
	if m.ID() != "020" {
		t.Errorf("Expected ID to be 020 but was %v", m.ID())
	}
}

func TestSillyObjectName(t *testing.T) {
	m := &mutation.SillyObject{}
	if m.Name() != "Silly Object" {
		t.Errorf("Expected Name to be Silly Object but was %v", m.Name())
	}
}

func TestSillyObjectCompatibleKinds(t *testing.T) {
	m := &mutation.SillyObject{}
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

func TestSillyObjectMutateInt(t *testing.T) {
	m := &mutation.SillyObject{}
	s := 1337
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestSillyObjectMutateFloat(t *testing.T) {
	m := &mutation.SillyObject{}
	s := 13.37
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestSillyObjectMutateString(t *testing.T) {
	m := &mutation.SillyObject{}
	s := "hello"
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestSillyObjectMutateNil(t *testing.T) {
	t.Skip("Skipping until I find out why the test is flaky")
	m := &mutation.SillyObject{}
	s := error(nil)
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}
