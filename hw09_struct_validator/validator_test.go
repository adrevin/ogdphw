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

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: User{
				ID:     "d174b2a2-be11-4695-871b-ecebe524058d",
				Name:   "name",
				Age:    18,
				Email:  "em@ai.l",
				Role:   "admin",
				Phones: []string{"88005555555", "88003333333"},
			},
			expectedErr: nil,
		},
		{
			in: App{
				Version: `0.0.0`,
			},
			expectedErr: nil,
		},
		{
			in: Response{
				Code: 200,
			},
			expectedErr: nil,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			for _, test := range tests {
				err := Validate(test.in)
				require.Nil(t, err)
			}
			_ = tt
		})
	}
}
