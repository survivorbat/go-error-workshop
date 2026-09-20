package mixer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestMix_FailsWithExpectedOutput(t *testing.T) {
	t.Run("TestMix_ReturnsErrorOnTooManyIngredients", func(t *testing.T) {
		// Arrange
		capturer := testFailures{}

		// Act
		testMix_ReturnsErrorOnTooManyIngredients(&capturer)

		// Assert
		expected := `
Error Trace:	/home/ruben/dev/go-workshop/go-error-workshop/exercise-5/mixer_test.go:30
/home/ruben/dev/go-workshop/go-error-workshop/exercise-5/mixer_test_test.go:16
	Error
		:      	Target error should be in err chain:
	            	expected: "`
		assert.Contains(t, capturer, expected)
	})
}

type testFailures []string

func (m *testFailures) Errorf(msg string, args ...any) {
	*m = append(*m, fmt.Sprintf(msg, args...))
}
