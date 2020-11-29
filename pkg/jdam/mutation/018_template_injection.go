package mutation

import (
	"math/rand"
	"reflect"
)

var templateInjectionPayloads = []string{
	"42*42",
	"{42*42}",
	"{{42*42}}",
	"{{{42*42}}}",
	"#{42*42}",
	"${42*42}",
	"<%=42*42 %>",
	"{{=42*42}}",
	"${donotexists|42*42}",
	"[[${42*42}]]",
}

// TemplateInjection replaces a value with a random template injection payload
// For more information on template injection: https://owasp.org/www-project-web-security-testing-guide/stable/4-Web_Application_Security_Testing/07-Input_Validation_Testing/18-Testing_for_Server_Side_Template_Injection
// Example: Hello -> {{{42*42}}}
type TemplateInjection struct{}

// ID returns mutator's 3-digit ID.
func (m *TemplateInjection) ID() string {
	return "018"
}

// Name returns the mutator's name.
func (m *TemplateInjection) Name() string {
	return "Template Injection"
}

// Description returns the mutator's description.
func (m *TemplateInjection) Description() string {
	return "Replace value with random template injection payload"
}

// CompatibleKinds returns the data types that the mutator can mutate.
func (m *TemplateInjection) CompatibleKinds() []reflect.Kind {
	return []reflect.Kind{reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Invalid}
}

// Mutate returns a random template injection payload.
// All payloads attempt to get the templating engine to evaluate 42 * 42
// which means that successful exploitation can be identified by looking for
// the number 1764 in server responses.
// Example: {{{42*42}}}
func (m *TemplateInjection) Mutate(subject reflect.Value, r *rand.Rand) interface{} {
	return templateInjectionPayloads[r.Intn(len(templateInjectionPayloads))]
}
