package fridge

type Config struct {
	// Name allows you to give your fridge a name
	Name string `json:"name"`

	// Temperature refers to the fridge's temperature
	Temperature int `json:"temperature"`
}
