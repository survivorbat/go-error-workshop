package blender

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 👋 Hey there! Fancy meeting you here. You don't need to look at this code unless you're stuck
// on exercise 5.
//
// In blender_test.go all error tests aren't actually Golang tests, because their names don't start
// with `Test` and their parameter isn't *testing.T. The tests below run  these functions and inspect
// what errors were reported using the `errorRecorder`. We then inspect the errors and test whether
// the fail contains the expected error messages, if any.

func TestBlendReturnsErrorOnNoIngredients(t *testing.T) {
	t.Parallel()
	// Arrange
	capturer := errorRecorder{}

	// Act
	BlendReturnsErrorOnNoIngredients(&capturer)

	// Assert
	const failMessage = "Test should have failed on the error not being ErrNoIngedients"

	if len(capturer) == 0 {
		t.Fatal(failMessage)
	}

	output := strings.Join(capturer, "\n")

	// Check if the correct ErrorIs is called
	if !strings.Contains(output, `expected: "no ingredients were provided"`) {
		t.Fatal(failMessage)
	}

	// Check that the expected failure mentions the wrong error
	if !strings.Contains(output, `in chain: "blender is on fire"`) {
		t.Fatal(failMessage)
	}
}

func TestBlendReturnsErrorOnTooManyIngredients(t *testing.T) {
	t.Parallel()
	// Arrange
	capturer := errorRecorder{}

	// Act
	BlendReturnsErrorOnTooManyIngredients(&capturer)

	// Assert
	const failMessage = "Test should have failed on the error not containing: 6 ingredients given, maximum is 5"

	if len(capturer) == 0 {
		t.Fatal(failMessage)
	}

	output := strings.Join(capturer, "\n")

	// Check that the expected contains is being performed
	if !strings.Contains(output, "does not contain") {
		t.Fatal(failMessage)
	}
	// Check that the expected message is present in the output
	if !strings.Contains(output, "6 ingredients given, maximum is 5") {
		t.Fatal(failMessage)
	}
}

func TestBlendReturnsErrorOnInvalidIngredients(t *testing.T) {
	t.Parallel()
	// Arrange
	capturer := errorRecorder{}

	// Act
	BlendReturnsErrorOnInvalidIngredients(&capturer)

	// Assert
	const failMessage = "Test should have failed on the InvalidIngredientsError not being equal to the invalid ingredients"

	if len(capturer) == 0 {
		t.Fatal(failMessage)
	}

	output := strings.Join(capturer, "\n")

	if !strings.Contains(output, "Not equal") {
		t.Fatal(failMessage)
	}

	// Check that
	if !strings.Contains(output, "Love") {
		t.Fatal(failMessage)
	}

	if !strings.Contains(output, `"Apple", "Banana", "Egg", "Orange", "Pear", "Strawberry", "Yoghurt"`) {
		t.Fatal(failMessage)
	}
}

// errorRecorder collects error calls from the Errorf method so we can inspect
// whether tests failed and what the failure messages are.
type errorRecorder []string

// Errorf implements assert.TestingT so it can be inserted in stretrchr/testify's assert method
func (m *errorRecorder) Errorf(msg string, args ...any) {
	*m = append(*m, fmt.Sprintf(msg, args...))
}

// Parallel does nothing, but it makes the tests look more "real"
func (*errorRecorder) Parallel() {}

// testingT is a combination of assert.testingT and an empty Parallel() call
type testingT interface {
	assert.TestingT
	Parallel()
}
