package mixer

import (
	"fmt"
	"time"
)

// Mix will take the ingredients and mix them together into a smoothie. Will return
// an error if more than 5 ingredients are given. All ingredients must be valid as defined
// in this package.
func Mix(ingredients ...Ingredient) error {
	if len(ingredients) > 5 {
		return fmt.Errorf("you put %d ingredients into the mixer, maximum is 5: %w", len(ingredients), ErrTooManyIngredients)
	}

	invalids := make([]Ingredient, 0, len(ingredients))
	for _, ingredient := range ingredients {
		if !ingredient.Valid() {
			invalids = append(invalids, ingredient)
		}
	}

	if len(invalids) > 0 {
		return &InvalidIngredientError{Invalids: invalids}
	}

	time.Sleep(20 * time.Millisecond)

	return nil
}
