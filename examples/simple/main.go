package main

import (
	"encoding/json"
	"fmt"

	"gitlab.com/michenriksen/jdam/pkg/jdam"
	"gitlab.com/michenriksen/jdam/pkg/jdam/mutation"
)

func main() {
	data := `{"userId":1,"id":1,"title":"Buy cookies","completed":false}`
	subject := map[string]interface{}{}

	// Decode data as JSON and put resulting object into subject variable
	err := json.Unmarshal([]byte(data), &subject)
	if err != nil {
		panic(err)
	}

	// Create a new fuzzer with all available mutators.
	// We instruct the fuzzer to ignore the id field so that its
	// value is never changed. This is useful in cases where you know
	// that altering this value would result in uninteresting errors.
	fuzzer := jdam.New(mutation.Mutators).IgnoreFields([]string{"id"})

	// Fuzz a random field with a random mutator.
	fuzzed := fuzzer.Fuzz(subject)

	// Encode the fuzzed object into JSON
	fuzzedJSON, err := json.Marshal(fuzzed)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(fuzzedJSON))
}
