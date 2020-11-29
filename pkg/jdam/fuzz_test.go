package jdam_test

import (
	"reflect"
	"testing"

	"gitlab.com/michenriksen/jdam/pkg/jdam"
	"gitlab.com/michenriksen/jdam/pkg/jdam/mutation"
)

func getFuzzer() *jdam.Fuzzer {
	return jdam.New(mutation.Mutators).NilChance(1)
}

func TestFuzzIntValue(t *testing.T) {
	subject := map[string]interface{}{"f": 1337}
	fuzzed := getFuzzer().Fuzz(subject)
	if subject["f"] == fuzzed["f"] {
		t.Error("Expected subject['f'] to not equal fuzzed['f']")
	}
}

func TestFuzzFloatValue(t *testing.T) {
	subject := map[string]interface{}{"f": 13.37}
	fuzzed := getFuzzer().Fuzz(subject)
	if subject["f"] == fuzzed["f"] {
		t.Errorf("Expected subject['f'] to not equal fuzzed['f']")
	}
}

func TestFuzzBoolValue(t *testing.T) {
	subject := map[string]interface{}{"f": true}
	fuzzed := getFuzzer().Fuzz(subject)
	if subject["f"] == fuzzed["f"] {
		t.Errorf("Expected subject['f'] to not equal fuzzed['f']")
	}
}

func TestFuzzStringValue(t *testing.T) {
	subject := map[string]interface{}{"f": "hello"}
	fuzzed := getFuzzer().Fuzz(subject)
	if subject["f"] == fuzzed["f"] {
		t.Error("Expected subject['f'] to not equal fuzzed['f']")
	}
}

func TestFuzzStringSliceValue(t *testing.T) {
	subject := map[string]interface{}{"f": []interface{}{"hello", "world"}}
	fuzzed := getFuzzer().Fuzz(subject)
	if reflect.DeepEqual(subject["f"], fuzzed["f"]) {
		t.Error("Expected subject['f'] to not equal fuzzed['f']")
	}
}

func TestFuzzIntSliceValue(t *testing.T) {
	subject := map[string]interface{}{"f": []interface{}{13, 37}}
	fuzzed := getFuzzer().Fuzz(subject)
	if reflect.DeepEqual(subject["f"], fuzzed["f"]) {
		t.Error("Expected subject['f'] to not equal fuzzed['f']")
	}
}

func TestFuzzMapValue(t *testing.T) {
	subject := map[string]interface{}{"f": map[string]interface{}{"sf": "hello"}}
	fuzzed := getFuzzer().Fuzz(subject)
	if reflect.DeepEqual(subject["f"], fuzzed["f"]) {
		t.Error("Expected subject['f'] to not equal fuzzed['f']")
	}
}

func TestNilChanceNever(t *testing.T) {
	subject := map[string]interface{}{"f": "hello"}
	fuzzer := getFuzzer().NilChance(1)
	for i := 0; i < 10; i++ {
		if fuzzer.Fuzz(subject)["f"] == nil {
			t.Error("Expected subject['f'] to not be nil")
		}
	}
}

func TestNilChanceAlways(t *testing.T) {
	subject := map[string]interface{}{"f": "hello"}
	fuzzer := getFuzzer().NilChance(0)
	for i := 0; i < 10; i++ {
		if fuzzer.Fuzz(subject)["f"] != nil {
			t.Error("Expected subject['f'] to be nil")
		}
	}
}

func TestIgnoreFields(t *testing.T) {
	subject := map[string]interface{}{"ignored": true, "other": "hello"}
	fuzzer := getFuzzer().IgnoreFields([]string{"ignored"})
	for i := 0; i < 10; i++ {
		if !fuzzer.Fuzz(subject)["ignored"].(bool) {
			t.Error("Expected subject['ignored'] to not change")
		}
	}
}

func TestNilChanceAboveOnePanics(t *testing.T) {
	defer func() { recover() }()
	getFuzzer().NilChance(1.1)
	t.Error("Expected NilChance to panic when p > 1")
}

func TestNilChanceBelowZeroPnanics(t *testing.T) {
	defer func() { recover() }()
	getFuzzer().NilChance(-1)
	t.Error("Expected NilChance to panic when p < 0")
}
