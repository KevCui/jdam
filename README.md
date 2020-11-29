# jdam - Structure-aware JSON fuzzing

jdam is a [Radamsa] inspired tool for fuzzing arbitrary JSON objects in a structure-aware fashion, which ensures that fuzzing results will always be valid JSON.

Many existing fuzzing tools will blindly alter the input and will often cause the result to be invalid JSON. This means that fuzz testing will only exercise the target application's JSON parser and the fuzz testing will never reach the underlying application code because the requests are invalid. Jdam is an attempt at solving this problem.

### Note
jdam is still pretty rough and errors should be expected. It is however stable enough to be released for the adventurous people out there, but treat it as alpha software for now!

## Installation
Download a pre-built [release] for your operating system or clone the repository and compile with `go build -o jdam cmd/jdam/*`.

## Usage

```
Usage of jdam:
  -count int
    	Number of fuzzed objects to generate (default 1)
  -ignore string
    	Comma-separated list of fields to exclude from fuzzing
  -list
    	List available mutators
  -max-depth int
    	Maximum object depth to fuzz (default 100)
  -mutators string
    	Comma-separated list of mutator IDs to use (default: all)
  -nil-chance float
    	Probability of value being set to nil (between 0 (no nils) and 1 (all nils)) (default 0.75)
  -output string
    	Output file pattern to use for results (e.g. /tmp/jdam-%d.json)
  -rounds int
    	Number of times to fuzz object (default 1)
  -seed int
    	Seed to use for pseudo-random number generator (default: current UNIX timestamp)
  -verbose
    	Print activity information
  -version
    	Print current jdam version

```

jdam works by piping a valid JSON object (`{...}`) into it. By default, jdam will perform one random permutation on a random field and print out the resulting JSON object:

```bash
$ echo '{"hello":"world"}' | jdam
{"hello":{"id":1}}
```

The subject JSON can be mutated several times with the `-rounds` flag, but keep in mind that previous fuzzing payloads can be mutated as well:

```bash
$ echo '{"hello":"world"}' | jdam -rounds 10
{"hello":{"_constructor":"${42*444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444442}"}}
```

Multiple fuzzed objects can be generated with the `-count` flag:

```bash
$ echo '{"hello":"world"}' | jdam -count 10
{"hello":"sorld"}
{"hello":"../../../../../../../../../../../../../../../../../../../../etc/passwd"}
{"hello":"{{42*42}}"}
{"hello":null}
{"hello":"' \u0026\u0026 this.password.match(/.*/)//+\u0000"}
{"hello":{"$where":"1"}}
{"hello":"null"}
{"hello":"\";touch /tmp/jdam.70797.fail'"}
{"hello":null}
{"hello":"|| 1==1"}
```

### Writing results to files

jdam can write fuzzing results to individual files if you want to prepare test files instead of using the command's STDOUT:

```bash
$ echo '{"hello":"world"}' | jdam -count 10 -output "/tmp/jdam-%d.json"
```

The above command will generate 10 files in `/tmp` from `jdam-1.json` to `jdam-10.json`. The special `%d` verb will be automatically replaced with a sequential number so that results can be written to individual files.

### Ignoring fields

If you want to ensure that jdam never mutates specific fields in the subject JSON object, you can tell jdam to ignore them with the `-ignore` flag:

```bash
$ echo '{"id":13,"articleId":37,"comment":"Hello"}' | ./jdam -count 10 -ignore id,articleId
{"id":13,"articleId":37,"comment":"%x%x%x%x"}
{"id":13,"articleId":37,"comment":"search=')] | //user/*[contains(*,'"}
{"id":13,"articleId":37,"comment":"' and count(/comment())=1 and '1'='1"}
{"id":13,"articleId":37,"comment":"\u0016ello"}
{"id":13,"articleId":37,"comment":"/"}
{"id":13,"articleId":37,"comment":null}
{"id":13,"articleId":37,"comment":"{{{42*42}}}"}
{"id":13,"articleId":37,"comment":"@*"}
{"id":13,"articleId":37,"comment":"Helloooooooooooooooo"}
{"id":13,"articleId":37,"comment":"//*"}
```

### Using jdam with Ffuf

[ffuf] is a fantastic web fuzzer and jdam can easily be used as input for fuzzing:

```bash
ffuf -input-cmd 'cat subject.json | jdam -rounds 3' -d FUZZ -u http://localhost/api/todos/1 -X PUT -v -mc 500 -mr 'error|failed|failure|fault|abort|root:|1764|0x|\d{20}' -od /tmp/ffuf_test
```

The above command will instruct ffuf to get its fuzzing input from jdam and send it as the request body to `PUT /api/todos/1`. request and response details will be written to `tmp/ffuf_test/` if the response is `500 Internal Server Error` or if the response body contains certain interesting strings.

### Using jdam as a package

If you know [Go] and need very specialized fuzzing, you can jdam within your own Go code. See the `examples/` folder for usage examples. 

[Radamsa]: https://gitlab.com/akihe/radamsa
[release]: https://gitlab.com/michenriksen/jdam/-/releases
[ffuf]: https://github.com/ffuf/ffuf/
[Go]: https://golang.org/
