package hw09structvalidator

import (
	"reflect"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (ve ValidationErrors) Error() string {
	err := ""
	if len(ve) == 0 {
		return err
	}
	for _, v := range ve {
		err += v.Err.Error()
	}
	return err
}

func Validate(v interface{}) error {
	var errors ValidationErrors

	rv := reflect.ValueOf(v)
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Type().Field(i)
		t := field.Tag.Get("validate")
		if t == "" {
			continue
		}
		value := rv.Field(i)
		switch value.Kind() {
		case reflect.Int:
			if err := validateIntField(field.Name, value.Int(), t); err != nil {
				errors = append(errors, *err)
			}
		case reflect.String:
			if err := validateStringField(field.Name, value.String(), t); err != nil {
				errors = append(errors, *err)
			}
		case reflect.Slice:
			if slice, ok := value.Interface().([]string); ok {
				for _, s := range slice {
					if err := validateStringField(field.Name, s, t); err != nil {
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

func validateIntField(n string, v int64, t string) *ValidationError {
	/*	"in:200,404,500"
		"min:18|max:50"
	*/
	return nil
}

func validateStringField(n string, v string, t string) *ValidationError {
	/*	"len:36"
		"regexp:^\w+@\w+\.\w+$"
		"in:admin,stuff"
	*/
	return nil
}
