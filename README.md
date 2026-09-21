# Go Error Workshop

Hey you! Ever worked with errors in Go? Extremely high chance you have.
But how do they work? When to use `errors.New` or `fmt.Errorf`, how to test if
errors are equal, and how to extract data?

This workshop's goal is to teach you best practices of using errors in Go.

## Content

This workshop will cover:

- _✨ Sentinel ✨_ errors in Go
- Error wrapping
- Error joining
- Custom error types and how to use them
- The `Unwrap` and `Is` methods
- Asserting errors using testify

## Prerequisites

- Basic Go syntax knowledge
- [Go 1.27 or higher](https://go.dev/doc/install) installed
- (optional) `make`, to use the makefile

## Goal

The goal of this workshop is to teach the value of using sentinel errors
and error types, as well as how to use them in your code.

All tests in the repository currently fail.
At the end of the workshop, they should succeed.

## Exercises

The first batch of exercises will be dealing with the [dishwasher SDK](https://github.com/survivorbat/go-error-workshop-lib).
A library that has not followed this workshop, but gets better each exercise.

- [Exercise 1](./exercise-1/README.md)
- [Exercise 2](./exercise-2/README.md)
- [Exercise 3](./exercise-3/README.md)

The fourth exercise involves returning errors in a library.
It does not involve dishwashers.

- [Exercise 4](./exercise-4/README.md)

The final exercise involves asserting errors using `stretchr/testify`.

- [Exercise 5](./exercise-5/README.md)
