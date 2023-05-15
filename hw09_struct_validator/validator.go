package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
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
)

var (
	ErrNotGreaterThanOrEqualMin = errors.New(NotGreaterThanOrEqualMin)
	ErrNotLessThanOrEqualMax    = errors.New(NotLessThanOrEqualMax)
	ErrNotInEnumeration         = errors.New(NotInEnumeration)
)

func (ve ValidationErrors) Error() string {
	err := ""
	if len(ve) == 0 {
		return err
	}
	for _, v := range ve {
		err += fmt.Sprintf("field \"%s\" has error: %s\n", v.Field, v.Err)
	}
	return err
}

func Validate(v interface{}) error {
	var errors ValidationErrors

	rv := reflect.ValueOf(v)
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Type().Field(i)
		tag := field.Tag.Get("validate")
		if tag == "" {
			continue
		}
		tags := strings.Split(tag, "|")
		value := rv.Field(i)
		switch value.Kind() {
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
		case
			reflect.Invalid,
			reflect.Bool,
			reflect.Int8,
			reflect.Int16,
			reflect.Int32,
			reflect.Int64,
			reflect.Uint,
			reflect.Uint8,
			reflect.Uint16,
			reflect.Uint32,
			reflect.Uint64,
			reflect.Uintptr,
			reflect.Float32,
			reflect.Float64,
			reflect.Complex64,
			reflect.Complex128,
			reflect.Array,
			reflect.Chan,
			reflect.Func,
			reflect.Interface,
			reflect.Map,
			reflect.Pointer,
			reflect.Struct,
			reflect.UnsafePointer:
			break
		}
	}

	return errors
}

func validateIntField(f string, v int64, tag string) *ValidationError {
	t := strings.Split(tag, ":")
	if len(t) != 2 {
		// invalid tag, not processed
		return nil
	}
	switch t[0] {
	case "in":
		values := make([]int64, 0)
		for _, vs := range strings.Split(t[1], ",") {
			v, err := strconv.ParseInt(vs, 10, 64)
			// invalid tag, not processed
			if err != nil {
				return nil
			}
			values = append(values, v)
		}
		// invalid tag, not processed
		if len(values) < 1 {
			return nil
		}
		inv := true
		for _, s := range values {
			if v == s {
				inv = false
				break
			}
		}
		if inv {
			return &ValidationError{Field: f, Err: ErrNotInEnumeration}
		}
		return nil
	case "max", "min":
		l, err := strconv.ParseInt(t[1], 10, 64)
		// invalid tag, not processed
		if err != nil {
			break
		}
		if t[0] == "max" && v > l {
			return &ValidationError{Field: f, Err: ErrNotLessThanOrEqualMax}
		}
		if t[0] == "min" && v < l {
			return &ValidationError{Field: f, Err: ErrNotGreaterThanOrEqualMin}
		}
		break
	}
	return nil
}

func validateStringField(n string, v string, t string) *ValidationError {
	/*	"len:36"
		"regexp:^\w+@\w+\.\w+$"
		"in:admin,stuff"
	*/
	return nil
}
