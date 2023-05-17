package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int             `validate:"min:18|max:50"`
		Email  string          `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole        `validate:"in:admin,stuff"`
		Phones []string        `validate:"len:11"`
		meta   json.RawMessage //nolint:unused
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

var (
	email = "em@a.il"
	role  = "admin"
	age   = 18
	user  = User{
		ID:     "d174b2a2-be11-4695-871b-ecebe524058d",
		Name:   "name",
		Age:    age,
		Email:  email,
		Role:   UserRole(role),
		Phones: []string{"88005555555", "88003333333"},
	}
	app      = App{Version: `0.0.0`}
	response = Response{Code: 200}
	token    = Token{}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{in: user, expectedErr: ValidationErrors(nil)},
		{in: app, expectedErr: ValidationErrors(nil)},
		{in: response, expectedErr: ValidationErrors(nil)},
		{in: token, expectedErr: ValidationErrors(nil)},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("valid data, case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			for _, test := range tests {
				err := Validate(test.in)
				require.Equal(t, test.expectedErr, err)
			}
			_ = tt
		})
	}

	t.Run("integers", func(t *testing.T) {
		u := user
		var ve ValidationErrors

		u.Age = 0
		ve = append(ve, ValidationError{Field: "Age", Err: ErrNotGreaterThanOrEqualMin})
		err := Validate(u)
		require.Equal(t, ve, err)

		u.Age = 100
		ve = ve[1:]
		ve = append(ve, ValidationError{Field: "Age", Err: ErrNotLessThanOrEqualMax})
		err = Validate(u)
		require.Equal(t, ve, err)

		u.Age = 100
		ve = ve[1:]
		ve = append(ve, ValidationError{Field: "Age", Err: ErrNotLessThanOrEqualMax})
		err = Validate(u)
		require.Equal(t, ve, err)

		u.Age = age
		u.Phones = []string{"8005555555", "8003333333"}
		ve = ve[1:]
		ve = append(ve, ValidationError{Field: "Phones", Err: ErrInvalidLength})
		ve = append(ve, ValidationError{Field: "Phones", Err: ErrInvalidLength})
		err = Validate(u)
		require.Equal(t, ve, err)

		r := response
		err = Validate(r)
		require.Nil(t, err)

		r.Code = 201
		ve = ve[2:]
		ve = append(ve, ValidationError{Field: "Code", Err: ErrNotInEnumeration})
		err = Validate(r)
		require.Equal(t, ve, err)
	})

	t.Run("strings", func(t *testing.T) {
		u := user
		var ve ValidationErrors

		u.Email = "@user"
		ve = append(ve, ValidationError{Field: "Email", Err: ErrDoesNotMatchRegExp})
		err := Validate(u)
		require.Equal(t, ve, err)

		u.Email = email
		u.Role = "invalid_role"
		ve = ve[1:]
		ve = append(ve, ValidationError{Field: "Role", Err: ErrNotInEnumeration})
		err = Validate(u)
		require.Equal(t, ve, err)

		u.Email = email
		u.Role = UserRole(role)
		u.ID = "0"
		ve = ve[1:]
		ve = append(ve, ValidationError{Field: "ID", Err: ErrInvalidLength})
		err = Validate(u)
		require.Equal(t, ve, err)
	})

	t.Run("invalid input", func(t *testing.T) {
		var ve ValidationErrors
		ve = append(ve, ValidationError{Field: "", Err: ErrUnsupportedDataType})
		err := Validate("")
		require.Equal(t, ve, err)
		err = Validate(0)
		require.Equal(t, ve, err)
	})
}
