package validatorgo

import (
	"time"
)

// IsDate formats
const (
	StandardDateLayout            = "2006-01-02"
	SlashDateLayout               = "2006/01/02"
	DateTimeLayout                = "2006-01-02 15:04:05"
	ISO8601Layout                 = "2006-01-02T15:04:05"
	ISO8601ZuluLayout             = "2006-01-02T15:04:05Z"
	ISO8601WithMillisecondsLayout = "2006-01-02T15:04:05.000Z"
)

var dateLayouts = [6]string{StandardDateLayout, SlashDateLayout, DateTimeLayout, ISO8601Layout, ISO8601ZuluLayout, ISO8601WithMillisecondsLayout}

var (
	isDateOptsDefaultFormat     string = StandardDateLayout
	isDateOptsDefaultStrictMode bool   = false
)

// IsDateOpts is used to configure IsDate
type IsDateOpts struct {
	Format     string
	StrictMode bool
}

func dateMatchesAnyFormat(str string) bool {
	for _, format := range dateLayouts {
		_, err := time.Parse(format, str)
		if err == nil {
			return true
		}
	}
	return false
}

// A Validator that checks if the string is a valid date. e.g. 2002-07-15.
//
// IsDateOpts is a struct which can contain the keys Format, StrictMode.
//
// Format: is a string and defaults to validatorgo.StandardDateLayout if "any" or no value is provided.
//
// StrictMode: is a boolean and defaults to false. If StrictMode is set to true, the validator will reject strings different from Format.
//
//	ok := validatorgo.IsDate("2006-01-02", &validatorgo.IsDateOpts{})
//	fmt.Println(ok) // true
//	ok := validatorgo.IsDate("01/023/2006", &validatorgo.IsDateOpts{})
//	fmt.Println(ok) // false
func IsDate(str string, opts *IsDateOpts) bool {
	if opts == nil {
		opts = setIsDateOptsToDefault()
	}

	switch opts.Format {
	case StandardDateLayout, SlashDateLayout, DateTimeLayout, ISO8601Layout, ISO8601ZuluLayout, ISO8601WithMillisecondsLayout:
	case "", "any":
		opts.Format = isDateOptsDefaultFormat
	default:
		return false
	}

	if opts.StrictMode {
		_, err := time.Parse(opts.Format, str)
		return err == nil
	} else {
		return dateMatchesAnyFormat(str)
	}
}

func setIsDateOptsToDefault() *IsDateOpts {
	return &IsDateOpts{
		Format:     isDateOptsDefaultFormat,
		StrictMode: isDateOptsDefaultStrictMode,
	}
}
