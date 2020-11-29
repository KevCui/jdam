package mutation

import (
	"math/rand"
	"reflect"
)

var sqlInjectionPayloads = []string{
	"1; SELECT 1",
	"1'; SELECT 1-- 1",
	`"' OR 1=1 -- 1`,
	"' OR '1'='1",
	`" or ""-"`,
	`" or "" "`,
	`" or ""&"`,
	`" or ""^"`,
	`" or ""*"`,
	`or true--`,
	`" or true--`,
	`' or true--`,
	`") or true--`,
	`') or true--`,
	`' or 'x'='x`,
	`') or ('x')=('x`,
	`')) or (('x'))=(('x`,
	`" or "x"="x`,
	`") or ("x")=("x`,
	`")) or (("x"))=(("x`,
	`or 1=1`,
	`or 1=1-- `,
	`or 1=1#`,
	`or 1=1/*`,
	`SLEEP(1) /*‘ or SLEEP(1) or ‘“ or SLEEP(1) or “*/`,
	`SLEEP(1) /*' or SLEEP(1) or '" or SLEEP(1) or "*/`,
}

// SQLInjection replaces a value with a random SQL injection payload
// For more information on SQL injection: https://en.wikipedia.org/wiki/SQL_injection
// Example: Hello -> "' OR 1=1 -- 1
type SQLInjection struct{}

// ID returns mutator's 3-digit ID.
func (m *SQLInjection) ID() string {
	return "011"
}

// Name returns the mutator's name.
func (m *SQLInjection) Name() string {
	return "SQL Injection"
}

// Description returns the mutator's description.
func (m *SQLInjection) Description() string {
	return "Replace value with random SQL injection payload"
}

// CompatibleKinds returns the data types that the mutator can mutate.
func (m *SQLInjection) CompatibleKinds() []reflect.Kind {
	return []reflect.Kind{reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Invalid}
}

// Mutate returns a random SQL injection payload.
// Example: "' OR 1=1 -- 1
func (m *SQLInjection) Mutate(subject reflect.Value, r *rand.Rand) interface{} {
	return sqlInjectionPayloads[r.Intn(len(sqlInjectionPayloads))]
}
