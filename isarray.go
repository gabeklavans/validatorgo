package validatorgo

import "encoding/json"

var (
	isArrayOptsDefaultMin *uint = nil
	isArrayOptsDefaultMax *uint = nil
)

// IsArrayOpts is used to configure IsArray
type IsArrayOpts struct {
	Min *uint // minimum array length
	Max *uint // maximum array length
}

// A validator to check that a value is an array.
//
// IsArrayOpts is a struct which defaults to { Min: nil, Max: nil }.
//
// You can also check that the array's length is greater than or equal to IsArrayOpts.Min and/or that it's less than or equal to IsArrayOpts.Max.
//
//	ok := validatorgo.IsArray(`["item1", "item2"]`, nil)
//	fmt.Println(ok) // true
//	ok := validatorgo.IsArray(`{"name": "John", "age": 30}`, nil)
//	fmt.Println(ok) // false
func IsArray(str string, opts *IsArrayOpts) bool {
	if opts == nil {
		opts = setIsArrayOptsToDefault()
	}

	var arr []interface{}

	err := json.Unmarshal([]byte(str), &arr)
	if err != nil {
		return false
	}

	arrLength := len(arr)
	withinLimits := true

	if opts.Min != nil {
		isMin := *(opts.Min) <= uint(arrLength)
		withinLimits = withinLimits && isMin
	}

	if opts.Max != nil {
		isMax := *(opts.Max) >= uint(arrLength)
		withinLimits = withinLimits && isMax
	}

	return withinLimits

}

func setIsArrayOptsToDefault() *IsArrayOpts {
	return &IsArrayOpts{
		Min: isArrayOptsDefaultMin,
		Max: isArrayOptsDefaultMax,
	}
}