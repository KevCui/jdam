package mutation_test

import (
	"reflect"
	"testing"

	"gitlab.com/michenriksen/jdam/pkg/jdam/mutation"
)

func TestLocalFileInclusionID(t *testing.T) {
	m := &mutation.LocalFileInclusion{}
	if m.ID() != "017" {
		t.Errorf("Expected ID to be 017 but was %v", m.ID())
	}
}

func TestLocalFileInclusionName(t *testing.T) {
	m := &mutation.LocalFileInclusion{}
	if m.Name() != "Local File Inclusion" {
		t.Errorf("Expected Name to be Local File Inclusion but was %v", m.Name())
	}
}

func TestLocalFileInclusionCompatibleKinds(t *testing.T) {
	m := &mutation.LocalFileInclusion{}
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

func TestLocalFileInclusionMutateInt(t *testing.T) {
	m := &mutation.LocalFileInclusion{}
	s := 1337
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestLocalFileInclusionMutateFloat(t *testing.T) {
	m := &mutation.LocalFileInclusion{}
	s := 13.37
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestLocalFileInclusionMutateString(t *testing.T) {
	m := &mutation.LocalFileInclusion{}
	s := "hello"
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestLocalFileInclusionMutateNil(t *testing.T) {
	m := &mutation.LocalFileInclusion{}
	s := error(nil)
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}
