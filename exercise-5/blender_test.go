package blender

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ⚠️ Do not change the names or input parameters of the tests below. They are executed in
// blender_test_test.go for the purpose of this exercise.

func TestBlend_ReturnsNoErrorOnSuccess(t *testing.T) {
	t.Parallel()
	// Arrange
	ingredients := []Ingredient{Apple, Banana}

	// Act
	err := Blend(ingredients...)

	// Assert
	assert.NoError(t, err)
}

// ❗ Test requirements:
//
// - err must be an error
// - err must errors.Is to ErrNoIngredients
func BlendReturnsErrorOnNoIngredients(t testingT) {
	t.Parallel()
	// Arrange
	ingredients := []Ingredient{}

	// Act
	err := Blend(ingredients...)

	// Assert
	assert.ErrorIs(t, err, ErrNoIngredients)
}

// ❗ Test requirements:
//
// - err must be an error
// - err must errors.Is to ErrTooManyIngredients
// - err must contain "6 ingredients given, maximum is 5"
func BlendReturnsErrorOnTooManyIngredients(t testingT) {
	t.Parallel()
	// Arrange
	ingredients := []Ingredient{Apple, Pear, Banana, Orange, Yoghurt, Egg}

	// Act
	err := Blend(ingredients...)

	// Assert
	assert.ErrorIs(t, err, ErrTooManyIngredients)
	assert.ErrorContains(t, err, "6 ingredients given, maximum is 5")
}

// ❗ Test requirements:
//
// - err must be an error
// - err must errors.As to InvalidIngredientsError
// - err must equal to InvalidIngredientsError{Invalid: []Ingredient{"Love"}}
func BlendReturnsErrorOnInvalidIngredients(t testingT) {
	t.Parallel()
	// Arrange
	ingredients := []Ingredient{"Love"}

	// Act
	err := Blend(ingredients...)

	// Assert
	var actualErr *InvalidIngredientsError
	assert.ErrorAs(t, err, &actualErr)

	expected := &InvalidIngredientsError{
		Invalid: ingredients,
	}
	assert.Equal(t, expected, actualErr)
}
