package blender

import "slices"

// Ingredient can be put inside of the blender
type Ingredient string

const (
	Apple      Ingredient = "Apple"
	Banana     Ingredient = "Banana"
	Egg        Ingredient = "Egg"
	Orange     Ingredient = "Orange"
	Pear       Ingredient = "Pear"
	Strawberry Ingredient = "Strawberry"
	Yoghurt    Ingredient = "Yoghurt"
)

// validIngredients is a list of all valid ingredients
var validIngredients = []Ingredient{
	Apple,
	Banana,
	Egg,
	Orange,
	Pear,
	Strawberry,
	Yoghurt,
}

// Valid returns whether the ingredient is valid
func (i Ingredient) Valid() bool {
	return slices.Contains(validIngredients, i)
}

// String is a convenience method to cast the type back to a string
func (i Ingredient) String() string {
	return string(i)
}
