package mutation_test

import (
	"reflect"
	"testing"

	"gitlab.com/michenriksen/jdam/pkg/jdam/mutation"
)

func TestTemplateInjectionID(t *testing.T) {
	m := &mutation.TemplateInjection{}
	if m.ID() != "018" {
		t.Errorf("Expected ID to be 018 but was %v", m.ID())
	}
}

func TestTemplateInjectionName(t *testing.T) {
	m := &mutation.TemplateInjection{}
	if m.Name() != "Template Injection" {
		t.Errorf("Expected Name to be Template Injection but was %v", m.Name())
	}
}

func TestTemplateInjectionCompatibleKinds(t *testing.T) {
	m := &mutation.TemplateInjection{}
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

func TestTemplateInjectionMutateInt(t *testing.T) {
	m := &mutation.TemplateInjection{}
	s := 1337
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestTemplateInjectionMutateFloat(t *testing.T) {
	m := &mutation.TemplateInjection{}
	s := 13.37
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestTemplateInjectionMutateString(t *testing.T) {
	m := &mutation.TemplateInjection{}
	s := "hello"
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}

func TestTemplateInjectionMutateNil(t *testing.T) {
	m := &mutation.TemplateInjection{}
	s := error(nil)
	rs := reflect.ValueOf(s)
	if s == m.Mutate(rs, getRand()) {
		t.Error("Expected mutation to be different from original")
	}
}
