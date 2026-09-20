package mixer

import (
	"errors"
	"fmt"
)

var ErrTooManyIngredients = errors.New("too many ingredients")

// InvalidIngredientError is returned if invalid ingredients were put into the mixer
type InvalidIngredientError struct {
	// Invalids is the list of all invalid ingredients that were provided
	Invalids []Ingredient
}

// Error implements the error interface and lists the invalid ingredients
func (i *InvalidIngredientError) Error() string {
	return fmt.Sprintf("invalid ingredients: %v", i.Invalids)
}
