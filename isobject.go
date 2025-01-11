package validatorgo

import (
	"encoding/json"
)

var (
	IsObjectOptsDefaultStrict = true
)

// IsObjectOpts is used to configure IsObject
type IsObjectOpts struct {
	Strict bool // looseness or strictness
}

// A validator to check that a value is a json object.
// For example, "{}", "{ foo: 'bar' }" would pass this validator.
//
// IsObjectOpts is a struct which defaults to { Strict: true }.
//
// If the Strict option is set to false, then this validator works, where both arrays ("[]") and the null ("null") value are considered objects.
//
//	ok := validatorgo.IsObject(`{"name": "John", "age": 30}`, &validatorgo.IsObjectOpts{Strict: true})
//	fmt.Println(ok) // true
//	ok := validatorgo.IsObject(`{"name": "John", "age`, &validatorgo.IsObjectOpts{Strict: true})
//	fmt.Println(ok) // false
func IsObject(str string, opts *IsObjectOpts) bool {
	if opts == nil {
		opts = setIsObjectOptsToDefault()
	}

	var obj interface{}

	err := json.Unmarshal([]byte(str), &obj)
	if err != nil {
		return false
	}

	switch obj.(type) {
	case map[string]interface{}:
		return true
	case []interface{}:
		return !opts.Strict
	default:
		if str == "null" && !opts.Strict {
			return true
		}
		return false
	}
}

func setIsObjectOptsToDefault() *IsObjectOpts {
	return &IsObjectOpts{
		Strict: IsObjectOptsDefaultStrict,
	}
}
