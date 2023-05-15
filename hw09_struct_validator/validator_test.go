package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
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

var user = User{
	ID:     "d174b2a2-be11-4695-871b-ecebe524058d",
	Name:   "name",
	Age:    18,
	Email:  "em@ai.l",
	Role:   "admin",
	Phones: []string{"88005555555", "88003333333"},
}
var app = App{Version: `0.0.0`}
var response = Response{Code: 200}

func TestValidate(t *testing.T) {

	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{in: user, expectedErr: nil},
		{in: app, expectedErr: nil},
		{in: response, expectedErr: nil},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("valid data, case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			for _, test := range tests {
				err := Validate(test.in)
				require.Nil(t, err)
			}
			_ = tt
		})
	}

	t.Run("integers", func(t *testing.T) {

		err := Validate(user)
		require.Nil(t, err)

		user.Age = 0
		err = Validate(user)
		require.NotNil(t, err)
		require.Equal(t, fmt.Sprintf("field \"Age\" has error: %s\n", NotGreaterThanOrEqualMin), err.Error())

		user.Age = 100
		err = Validate(user)
		require.NotNil(t, err)
		require.Equal(t, fmt.Sprintf("field \"Age\" has error: %s\n", NotLessThanOrEqualMax), err.Error())

		err = Validate(response)
		require.Nil(t, err)

		response.Code = 201
		err = Validate(response)
		require.NotNil(t, err)
		require.Equal(t, fmt.Sprintf("field \"Code\" has error: %s\n", NotInEnumeration), err.Error())
	})
}
