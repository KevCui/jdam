package mutation_test

import (
	"math/rand"
	"reflect"
	"time"

	"gitlab.com/michenriksen/jdam/pkg/jdam/mutation"
)

func mutatorIsCompatibleWith(m mutation.Mutator, k reflect.Kind) bool {
	for _, v := range m.CompatibleKinds() {
		if v == k {
			return true
		}
	}
	return false
}

func getRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
