package validator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/validator"
)

func TestValidator(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		expectedMessages []string
		shouldHasError   bool
		errors           map[string][]string
	}{
		{
			name:           "has no errors",
			shouldHasError: false,
		},
		{
			name:             "should return error message",
			shouldHasError:   true,
			expectedMessages: []string{"err-context: err-message"},
			// {context: [message1, message2, ...]}
			errors: map[string][]string{
				"err-context": {"err-message"},
			},
		},
		{
			name:             "should show all error messages",
			shouldHasError:   true,
			expectedMessages: []string{"context1: error message 1", "context2: error message 2"},
			errors: map[string][]string{
				"context2": {"error message 2"},
				"context1": {"error message 1"},
			},
		},
		{
			name:             "should agroup all context errors",
			shouldHasError:   true,
			expectedMessages: []string{"field: error 1, error 2"},
			errors: map[string][]string{
				"field": {"error 1", "error 2"},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := validator.New()

			for context, messages := range tc.errors {
				for _, msg := range messages {
					sut.AddError(context, msg)
				}
			}

			if tc.shouldHasError {
				assert.True(t, sut.HasErrors())
				errorMessage := sut.Validate().Error()

				for _, expectedMessage := range tc.expectedMessages {
					assert.Contains(t, errorMessage, expectedMessage)
				}
			} else {
				assert.False(t, sut.HasErrors())
				assert.Nil(t, sut.Validate())
			}
		})
	}
}
