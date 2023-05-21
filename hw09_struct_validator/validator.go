package hw09structvalidator

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type DataError struct {
	Err error
}

func (d DataError) Error() string {
	panic(d.Err)
}

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

const (
	NotGreaterThanOrEqualMin = "not greater than or equal min"
	NotLessThanOrEqualMax    = "not less than or equal max"
	NotInEnumeration         = "not in enumeration"
	InvalidLength            = "invalid length"
	DoesNotMatchRegExp       = "does not match regular expression"
	UnsupportedDataType      = "unsupported data type"
	UnsupportedFieldType     = "unsupported field type"
	InvalidTag               = "invalid tag"
)

// Validation errors.
var (
	ErrValidationNotGreaterThanOrEqualMin = errors.New(NotGreaterThanOrEqualMin)
	ErrValidationNotLessThanOrEqualMax    = errors.New(NotLessThanOrEqualMax)
	ErrValidationNotInEnumeration         = errors.New(NotInEnumeration)
	ErrValidationInvalidLength            = errors.New(InvalidLength)
	ErrValidationDoesNotMatchRegExp       = errors.New(DoesNotMatchRegExp)
)

// Data errors.
var (
	ErrDataUnsupportedType      = DataError{Err: errors.New(UnsupportedDataType)}
	ErrDataUnsupportedFieldType = DataError{Err: errors.New(UnsupportedFieldType)}
	ErrDataInvalidTag           = DataError{Err: errors.New(InvalidTag)}
)

func (ve ValidationErrors) Error() string {
	var buffer bytes.Buffer
	for _, v := range ve {
		buffer.WriteString(fmt.Sprintf("field \"%s\" has error: %s\n", v.Field, v.Err))
	}
	return buffer.String()
}

func Validate(v interface{}) error { //nolint:gocognit
	var errors ValidationErrors

	if reflect.TypeOf(v).Kind() != reflect.Struct {
		return ErrDataUnsupportedType
	}
	rv := reflect.ValueOf(v)
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Type().Field(i)
		tag := field.Tag.Get("validate")
		if tag == "" {
			continue
		}
		tags := strings.Split(tag, "|")
		value := rv.Field(i)
		switch value.Kind() { //nolint:exhaustive
		case reflect.Int:
			for _, tag = range tags {
				if err := validateIntField(field.Name, value.Int(), tag); err != nil {
					errors = append(errors, *err)
				}
			}
		case reflect.String:
			if err := validateStringField(field.Name, value.String(), tag); err != nil {
				errors = append(errors, *err)
			}
		case reflect.Slice:
			if slice, ok := value.Interface().([]string); ok {
				for _, s := range slice {
					if err := validateStringField(field.Name, s, tag); err != nil {
						errors = append(errors, *err)
					}
				}
			} else if slice, ok := value.Interface().([]int64); ok {
				for _, s := range slice {
					if err := validateIntField(field.Name, s, tag); err != nil {
						errors = append(errors, *err)
					}
				}
			} else {
				return ErrDataUnsupportedFieldType
			}
		default:
			return ErrDataUnsupportedFieldType
		}
	}

	return errors
}

func validateIntField(f string, v int64, tag string) *ValidationError {
	t := strings.Split(tag, ":")
	if len(t) != 2 {
		return &ValidationError{Field: f, Err: ErrDataInvalidTag}
	}
	key := t[0]
	val := t[1]
	switch key {
	case "in":
		isInRange := false
		for _, stringItem := range strings.Split(val, ",") {
			intItem, err := strconv.ParseInt(stringItem, 10, 64)
			if err != nil {
				return &ValidationError{Field: f, Err: ErrDataInvalidTag}
			}
			if v == intItem {
				isInRange = true
				// with no 'break' to validate the entire tag
			}
		}
		if !isInRange {
			return &ValidationError{Field: f, Err: ErrValidationNotInEnumeration}
		}
		return nil
	case "max", "min":
		l, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return &ValidationError{Field: f, Err: ErrDataInvalidTag}
		}
		if key == "max" && v > l {
			return &ValidationError{Field: f, Err: ErrValidationNotLessThanOrEqualMax}
		}
		if key == "min" && v < l {
			return &ValidationError{Field: f, Err: ErrValidationNotGreaterThanOrEqualMin}
		}
		return nil
	}
	return nil
}

func validateStringField(f string, v string, tag string) *ValidationError {
	t := strings.Split(tag, ":")
	if len(t) != 2 {
		// invalid tag, not processed
		return nil
	}
	key := t[0]
	val := t[1]
	switch key {
	case "len":
		l, err := strconv.Atoi(val)
		if err != nil {
			return &ValidationError{Field: f, Err: ErrDataInvalidTag}
		}
		if len(v) < l {
			return &ValidationError{Field: f, Err: ErrValidationInvalidLength}
		}
		return nil
	case "regexp":
		r := regexp.MustCompile(val)
		if !r.MatchString(v) {
			return &ValidationError{Field: f, Err: ErrValidationDoesNotMatchRegExp}
		}
		return nil
	case "in":
		isInRange := false
		for _, vs := range strings.Split(val, ",") {
			if v == vs {
				isInRange = true
				break
			}
		}
		if !isInRange {
			return &ValidationError{Field: f, Err: ErrValidationNotInEnumeration}
		}
		return nil
	}
	return nil
}
