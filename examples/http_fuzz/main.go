package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

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

	// initialize http client
	client := &http.Client{}

	// Create a new fuzzer with all available mutators.
	// We instruct the fuzzer to ignore the id field so that its
	// value is never changed. This is useful in cases where you know
	// that altering this value would result in uninteresting errors.
	fuzzer := jdam.New(mutation.Mutators).IgnoreFields([]string{"id"})

	// Fuzz the target API endpoint 1000 times
	for i := 0; i < 1000; i++ {
		// Fuzz a random field with a random mutator.
		fuzzed := fuzzer.Fuzz(subject)

		// Encode the fuzzed object into JSON.
		fuzzedJSON, err := json.Marshal(fuzzed)
		if err != nil {
			panic(err)
		}

		// prepare a PUT (update) request to /api/todos/1 with our fuzzed payload.
		req, err := http.NewRequest(http.MethodPut, "http://localhost/api/todos/1", bytes.NewBuffer(fuzzedJSON)) // DevSkim: ignore DS137138
		if err != nil {
			panic(err)
		}
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		req.Header.Set("User-Agent", "MyCustomFuzzer")

		// Send request to the server.
		resp, err := client.Do(req)
		if err != nil {
			// Our payload may have crashed the server!
			// Write the payload to a file for further research and stop fuzzing.
			ioutil.WriteFile(fmt.Sprintf("crash-%d.json", i), fuzzedJSON, 0644)
			break
		}

		if resp.StatusCode == 500 {
			// Our payload has caused some sort of internal server error!
			// Write the payload to a file for further research.
			ioutil.WriteFile(fmt.Sprintf("error-%d.json", i), fuzzedJSON, 0644)
		}

		// Sleep for a bit to be nice to the server (lol).
		time.Sleep(200 * time.Millisecond)
	}
}
