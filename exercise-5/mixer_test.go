package mixer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMix_ReturnsNoErrorOnSuccess(t *testing.T) {
	// Arrange
	ingredients := []Ingredient{Apple, Banana}

	// Act
	err := Mix(ingredients...)

	// Assert
	assert.NoError(t, err)
}

// The tests below are executed through mixer_test_test.go

func testMix_ReturnsErrorOnTooManyIngredients(t assert.TestingT) {
	// Arrange
	ingredients := []Ingredient{Apple, Pear, Banana, Orange, Yoghurt, Egg}

	// Act
	err := Mix(ingredients...)

	// Assert
	assert.ErrorIs(t, err, ErrTooManyIngredients)
}

func testMix_ReturnsErrorOnInvalidIngredients(t assert.TestingT) {
	// Arrange
	ingredients := []Ingredient{"Love"}

	// Act
	err := Mix(ingredients...)

	// Assert
	assert.Error(t, err)
}
