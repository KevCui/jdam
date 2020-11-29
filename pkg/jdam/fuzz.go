package jdam

import (
	"math/rand"
	"reflect"
	"time"

	"gitlab.com/michenriksen/jdam/pkg/jdam/mutation"
)

// Fuzzer can mutate random values in maps.
type Fuzzer struct {
	curDepth     int
	maxDepth     int
	ignoreFields []string
	nilChance    float64
	mutators     mutation.MutatorMap
	r            *rand.Rand
}

// New returns a new fuzzer. The fuzzer can be configured with
// MaxDepth, NilChance, and IgnoreFields.
func New(mutators mutation.MutatorList) *Fuzzer {
	return NewWithSeed(time.Now().UnixNano(), mutators)
}

// NewWithSeed returns a new fuzzer with the given seed used for
// its internal pseudo-random number generator.
func NewWithSeed(seed int64, mutators mutation.MutatorList) *Fuzzer {
	f := &Fuzzer{
		curDepth:  0,
		maxDepth:  100,
		nilChance: .75,
		r:         rand.New(rand.NewSource(seed)),
	}
	return f.makeMutatorMap(mutators)
}

// MaxDepth sets the maximum number of recursive fuzz calls that will be made
// before stopping.
func (f *Fuzzer) MaxDepth(d int) *Fuzzer {
	f.maxDepth = d
	return f
}

// NilChance sets the probability of setting a field value to nil.
// 'p'. 'p' should be between 0 (no nils) and 1 (all nils), inclusive.
func (f *Fuzzer) NilChance(p float64) *Fuzzer {
	if p < 0 || p > 1 {
		panic("p should be between 0 and 1, inclusive.")
	}
	f.nilChance = p
	return f
}

// IgnoreFields instructs the fuzzer to ignore map keys with the given names.
func (f *Fuzzer) IgnoreFields(s []string) *Fuzzer {
	f.ignoreFields = s
	return f
}

// Fuzz fuzzes a random value in a map and returns the resulting map.
func (f *Fuzzer) Fuzz(obj map[string]interface{}) map[string]interface{} {
	return f.doFuzz(obj)
}

func (f *Fuzzer) doFuzz(subject map[string]interface{}) map[string]interface{} {
	f.curDepth++
	if f.curDepth > f.maxDepth {
		return subject
	}
	fuzzed := map[string]interface{}{}
	targetField := f.findTargetField(subject)
	for k, v := range subject {
		rv := reflect.ValueOf(v)
		if k != targetField {
			fuzzed[k] = v
			continue
		}
		if f.r.Float64() > f.nilChance {
			fuzzed[k] = nil
			continue
		}
		switch rv.Kind() {
		case reflect.Map:
			fuzzed[k] = f.doFuzz(v.(map[string]interface{}))
		case reflect.Slice:
			fuzzed[k] = f.fuzzSlice(v.([]interface{}))
		default:
			mutator := f.randMutator(rv.Kind())
			if mutator == nil {
				fuzzed[k] = v
				continue
			}
			fuzzed[k] = mutator.Mutate(rv, f.r)
		}
	}
	return fuzzed
}

func (f *Fuzzer) findTargetField(m map[string]interface{}) string {
	for {
		field := f.randMapKey(m)
		if !f.ignoredField(field) {
			return field
		}
	}
}

func (f *Fuzzer) ignoredField(field string) bool {
	for _, v := range f.ignoreFields {
		if field == v {
			return true
		}
	}
	return false
}

func (f *Fuzzer) randMapKey(m map[string]interface{}) string {
	keys := reflect.ValueOf(m).MapKeys()
	return keys[f.r.Intn(len(keys))].String()
}

func (f *Fuzzer) randMutator(k reflect.Kind) mutation.Mutator {
	if val, ok := f.mutators[k]; ok {
		if len(val) == 0 {
			return nil
		}
		return val[f.r.Intn(len(val))]
	}
	return nil
}

func (f *Fuzzer) fuzzSlice(a []interface{}) []interface{} {
	if len(a) == 0 {
		return a
	}
	tmp := make([]interface{}, len(a))
	copy(tmp, a)
	f.r.Shuffle(len(a), func(i, j int) { tmp[i], tmp[j] = tmp[j], tmp[i] })
	for i, v := range tmp {
		rv := reflect.ValueOf(v)
		mutator := f.randMutator(rv.Kind())
		if mutator == nil {
			continue
		}
		tmp[i] = mutator.Mutate(rv, f.r)
		break
	}
	return tmp
}

func (f *Fuzzer) makeMutatorMap(mutators mutation.MutatorList) *Fuzzer {
	f.mutators = make(mutation.MutatorMap)
	for _, m := range mutators {
		for _, k := range m.CompatibleKinds() {
			if val, ok := f.mutators[k]; ok {
				f.mutators[k] = append(val, m)
			} else {
				f.mutators[k] = []mutation.Mutator{m}
			}
		}
	}
	return f
}
