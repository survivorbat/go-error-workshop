package blender

import (
	"errors"
	"fmt"
	"time"
)

// ⚠️ Do not touch this code, fixing the bugs is not part of the exercise

var (
	ErrNoIngedients       = errors.New("no ingredients were provided")
	ErrTooManyIngredients = errors.New("too many ingredients")
	ErrBlenderIsOnFire    = errors.New("blender is on fire")
)

// Blend will take the ingredients and blend them together into a smoothie. Will return
// an error if more than 5 ingredients are given. All ingredients must be valid as defined in validIngredients.
// in this package.
func Blend(ingredients ...Ingredient) error {
	// If no ingredients are given, report ErrNoIngedients
	if len(ingredients) == 0 {
		return ErrBlenderIsOnFire // 🪲 BUG
	}

	// If there are too many ingredients, report how many were given, how many are allowed and
	// wrap ErrTooManyIngredients
	if len(ingredients) > 5 {
		return fmt.Errorf("the blender is on fire: %w", ErrTooManyIngredients) // 🪲 BUG
	}

	invalids := make([]Ingredient, 0, len(ingredients))
	for _, ingredient := range ingredients {
		if !ingredient.Valid() {
			invalids = append(invalids, ingredient)
		}
	}

	// If there are invalid ingredients, inform the user what ingredients were invalid
	if len(invalids) > 0 {
		return &InvalidIngredientsError{Invalid: validIngredients} // 🪲 BUG
	}

	// Blending...
	time.Sleep(20 * time.Millisecond)

	return nil
}

// InvalidIngredientsError is returned if invalid ingredients were put into the blender
type InvalidIngredientsError struct {
	// Invalid is the list of all invalid ingredients that were provided
	Invalid []Ingredient
}

// Error implements the error interface and lists the invalid ingredients
func (i *InvalidIngredientsError) Error() string {
	return fmt.Sprintf("invalid ingredients: %v", i.Invalid)
}
