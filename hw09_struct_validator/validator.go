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
	InvalidTag               = "invalid tag"
)

var (
	ErrNotGreaterThanOrEqualMin = errors.New(NotGreaterThanOrEqualMin)
	ErrNotLessThanOrEqualMax    = errors.New(NotLessThanOrEqualMax)
	ErrNotInEnumeration         = errors.New(NotInEnumeration)
	ErrInvalidLength            = errors.New(InvalidLength)
	ErrDoesNotMatchRegExp       = errors.New(DoesNotMatchRegExp)
	ErrUnsupportedDataType      = errors.New(UnsupportedDataType)
	ErrInvalidTag               = errors.New(InvalidTag)
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
		errors = append(errors, ValidationError{Field: "", Err: ErrUnsupportedDataType})
		return errors
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
			}
			if slice, ok := value.Interface().([]int64); ok {
				for _, s := range slice {
					if err := validateIntField(field.Name, s, tag); err != nil {
						errors = append(errors, *err)
					}
				}
			}
		default:
			errors = append(errors, ValidationError{Field: field.Name, Err: ErrUnsupportedDataType})
		}
	}

	return errors
}

func validateIntField(f string, v int64, tag string) *ValidationError {
	t := strings.Split(tag, ":")
	if len(t) != 2 {
		return &ValidationError{Field: f, Err: ErrInvalidTag}
	}
	key := t[0]
	val := t[1]
	switch key {
	case "in":
		isInRange := false
		for _, stringItem := range strings.Split(val, ",") {
			intItem, err := strconv.ParseInt(stringItem, 10, 64)
			if err != nil {
				return &ValidationError{Field: f, Err: ErrInvalidTag}
			}
			if v == intItem {
				isInRange = true
				// with no 'break' to validate the entire tag
			}
		}
		if !isInRange {
			return &ValidationError{Field: f, Err: ErrNotInEnumeration}
		}
		return nil
	case "max", "min":
		l, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return &ValidationError{Field: f, Err: ErrInvalidTag}
		}
		if key == "max" && v > l {
			return &ValidationError{Field: f, Err: ErrNotLessThanOrEqualMax}
		}
		if key == "min" && v < l {
			return &ValidationError{Field: f, Err: ErrNotGreaterThanOrEqualMin}
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
			return &ValidationError{Field: f, Err: ErrInvalidTag}
		}
		if len(v) < l {
			return &ValidationError{Field: f, Err: ErrInvalidLength}
		}
		return nil
	case "regexp":
		r := regexp.MustCompile(val)
		if !r.MatchString(v) {
			return &ValidationError{Field: f, Err: ErrDoesNotMatchRegExp}
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
			return &ValidationError{Field: f, Err: ErrNotInEnumeration}
		}
		return nil
	}
	return nil
}
