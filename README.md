# jdam - Structure-aware JSON fuzzing

![jdam](.gitlab/jdam.gif)

jdam is a [Radamsa] inspired tool for fuzzing arbitrary JSON objects in a structure-aware fashion, which ensures that fuzzing results will always be valid JSON.

Many existing fuzzing tools will blindly alter the input and will often cause the result to be invalid JSON. This means that fuzz testing will only exercise the target application's JSON parser and the fuzz testing will never reach the underlying application code because the requests are invalid. Jdam is an attempt at solving this problem.

jdam comes with several mutation modules that all aim to uncover potential problems and vulnerabilities in the systems that process the fuzzed data. Some modules perform random changes such as dropping, replacing, swapping, and repeating random bytes, inverting boolean values, and number negation, while other modules replace values with payloads that attempt to trigger specific vulnerability types such as:

- [SQL Injection]
- [Command Injection]
- [LDAP Injection]
- [NoSQL Injection]
- [Format String Injection]
- [Local File Inclusion]
- [Integer Overflow]
- and more

To see a full list of available mutators, use `jdam -list`.

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

If you know [Go] and need very specialized fuzzing, you can use jdam within your own Go code. See the `examples/` folder for usage examples.

## What to look for

Fuzzing is only useful if you can detect that a particular mutation caused a problem or triggered a vulnerability. These signs will vary depending on the target system and environment, but here are some good general things to look out for:

1. **Crashes:** if the target system is no longer accepting connections there is a good chance that it crashed because of some sort of memory corruption bug. Replay the last request and see if the crash can be reproduced.
1. **Unusually slow responses:** if a request causes the server to respond unusually slow, there is a good chance that a payload caused problems for an underlying system or instructed the database to `sleep`.
1. **Unusually large responses:** if a response is unusually large compared to the average, it might contain interesting data that wasn't intended to be returned.
1. **500 Internal Server Error** and other error responses: these usually happen when a payload triggered an unexpected or unhandled error in the application.
1. **Error and failure keywords:** Responses that contain error-related keywords such as `error`, `failure`, `failed`, etc.
1. **/etc/passwd file contents:** all local file inclusion payloads attempt to include the `/etc/passwd` file in order to have predictable content. So any response containing `root:` is very interesting!
1. **The magic number 1764:** all template injection payloads will attempt to make vulnerable templating engines evaluate `42*42` which equals 1.764. If this number ever occurs in a response, there is a high chance that the system is vulnerable to template injection.
1. **0x and long numbers:** if a response contains `0x` followed by a long hex string or a long numeric string, it could be an indication that a payload triggered a format string vulnerability.
1. **The presence of /tmp/jdam.*.fail:** All command injection payloads attempt to create a file in the system's `/tmp` folder with the name `jdam.<number>.fail`. The number will be a random five  digit number to make it easier to identify the responsible payload.

## Trophy Room

jdam is still a very new tool so it hasn't uncovered any good bugs yet. I would love to hear from you if jdam has helped you score a CVE or a nice bug bounty and I will gladly link to your write-up or report here. :)

## Credit

I would like to give credit and thanks to a couple of other projects:

- [Radamsa] for the inspiration for jdam.
- [gofuzz] for giving me ideas on how to code jdam.
- [SecLists] and [PayloadsAllTheThings] for fuzzing payloads.
- [ffuf] for making it super easy to use jdam against web apps.

**Happy Fuzzing!**

[Radamsa]: https://gitlab.com/akihe/radamsa
[SQL Injection]: https://owasp.org/www-community/attacks/SQL_Injection
[Command Injection]: https://owasp.org/www-community/attacks/Command_Injection
[LDAP Injection]: https://owasp.org/www-community/attacks/LDAP_Injection
[NoSQL Injection]: https://owasp.org/www-project-web-security-testing-guide/latest/4-Web_Application_Security_Testing/07-Input_Validation_Testing/05.6-Testing_for_NoSQL_Injection
[Format String Injection]: https://owasp.org/www-community/attacks/Format_string_attack
[Local File Inclusion]: https://owasp.org/www-project-web-security-testing-guide/latest/4-Web_Application_Security_Testing/07-Input_Validation_Testing/11.1-Testing_for_Local_File_Inclusion
[Integer Overflow]: https://en.wikipedia.org/wiki/Integer_overflow
[release]: https://gitlab.com/michenriksen/jdam/-/releases
[ffuf]: https://github.com/ffuf/ffuf/
[Go]: https://golang.org/
[gofuzz]: https://github.com/google/gofuzz
[SecLists]: https://github.com/danielmiessler/SecLists
[PayloadsAllTheThings]: https://github.com/swisskyrepo/PayloadsAllTheThings
