package mutation_test

import (
	"reflect"
	"testing"

	"gitlab.com/michenriksen/jdam/pkg/jdam/mutation"
)

func TestNumericToNegativeID(t *testing.T) {
	m := &mutation.NumericToNegative{}
	if m.ID() != "003" {
		t.Errorf("Expected ID to be 003 but was %v", m.ID())
	}
}

func TestNumericToNegativeName(t *testing.T) {
	m := &mutation.NumericToNegative{}
	if m.Name() != "Numeric To Negative" {
		t.Errorf("Expected Name to be Numeric To Negative but was %v", m.Name())
	}
}

func TestNumericToNegativeCompatibleKinds(t *testing.T) {
	m := &mutation.NumericToNegative{}
	if !mutatorIsCompatibleWith(m, reflect.Int) {
		t.Error("Expected mutator to be compatible with reflect.Int")
	}
	if !mutatorIsCompatibleWith(m, reflect.Float64) {
		t.Error("Expected mutator to be compatible with reflect.Float64")
	}
	if mutatorIsCompatibleWith(m, reflect.Invalid) {
		t.Error("Expected mutator to not be compatible with reflect.Invalid")
	}
	if mutatorIsCompatibleWith(m, reflect.String) {
		t.Error("Expected mutator to not be compatible with reflect.String")
	}
}

func TestNumericToNegativeMutateInt(t *testing.T) {
	m := &mutation.NumericToNegative{}
	s := 1337
	rs := reflect.ValueOf(s)
	mu := m.Mutate(rs, getRand())
	if mu.(int64) != -1337 {
		t.Errorf("Expected mutation to equal -1337 but was %v", mu)
	}
}

func TestNumericToNegativeMutateFloat(t *testing.T) {
	m := &mutation.NumericToNegative{}
	s := 13.37
	rs := reflect.ValueOf(s)
	mu := m.Mutate(rs, getRand())
	if mu.(float64) != -13.37 {
		t.Errorf("Expected mutation to equal -1337 but was %v", mu)
	}
}
