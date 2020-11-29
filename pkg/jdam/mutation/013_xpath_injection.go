package mutation

import (
	"math/rand"
	"reflect"
)

var xpathInjectionPayloads = []string{
	"/",
	"//",
	"//*",
	"*/*",
	"@*",
	"count(/child::node())",
	"x' or name()='username' or 'x'='y",
	"' and count(/*)=1 and '1'='1",
	"' and count(/@*)=1 and '1'='1",
	"' and count(/comment())=1 and '1'='1",
	"search=')] | //user/*[contains(*,'",
	"search=Har') and contains(../password,'c",
	"search=Har') and starts-with(../password,'c",
}

// XPATHInjection replaces a value with a random XPATH injection payload.
// For more information on XPATH injection: https://owasp.org/www-community/attacks/XPATH_Injection
// Example: Hello -> x' or name()='username' or 'x'='y
type XPATHInjection struct{}

// ID returns mutator's 3-digit ID.
func (m *XPATHInjection) ID() string {
	return "013"
}

// Name returns the mutator's name.
func (m *XPATHInjection) Name() string {
	return "XPATH Injection"
}

// Description returns the mutator's description.
func (m *XPATHInjection) Description() string {
	return "Replace value with random XPATH injection payload"
}

// CompatibleKinds returns the data types that the mutator can mutate.
func (m *XPATHInjection) CompatibleKinds() []reflect.Kind {
	return []reflect.Kind{reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Invalid}
}

// Mutate returns a random XPATH injection payload.
// Example: x' or name()='username' or 'x'='y
func (m *XPATHInjection) Mutate(subject reflect.Value, r *rand.Rand) interface{} {
	return xpathInjectionPayloads[r.Intn(len(xpathInjectionPayloads))]
}
