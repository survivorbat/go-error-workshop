package fridge

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResponseError_IsEqualToAPIError(t *testing.T) {
	t.Parallel()
	// Arrange
	err := &ResponseError{
		Body:      []byte("🦁"),
		ActualErr: assert.AnError,
	}

	// Act
	actual := errors.Is(err, ErrAPIError)

	// Assert
	assert.True(t, actual, "not equal to API error")
}

func TestResponseError_InnerErrorIsAccessible(t *testing.T) {
	t.Parallel()
	// Arrange
	err := &ResponseError{
		Body:      []byte("🦁"),
		ActualErr: &dummyError{message: "foo, bar, baz"},
	}

	// Act
	actual, ok := errors.AsType[*dummyError](err)

	// Assert
	require.True(t, ok, "does not AsType to dummy error")
	assert.Equal(t, "foo, bar, baz", actual.message)
}

// dummyError is used as an inner error in the tests above
type dummyError struct {
	message string
}

func (d *dummyError) Error() string {
	return d.message
}
