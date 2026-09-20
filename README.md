# Go Error Workshop

Hey you! Ever worked with errors in Golang? Extremely high chance you have.
But how do they work? When to use `errors.New` or `fmt.Errorf`, how to test if
errors are equal, and how to extract data?

This workshop's goal is to teach you best practices of using errors in Golang.

## Content

This workshop will contain information regarding:

- _✨ Sentinel ✨_ errors in Golang
- Error wrapping
- Error joining
- Custom error types and how to use them
- The `Unwrap` and `Is` methods
- Asserting errors using testify

## Prerequisites

- Basic Golang syntax knowledge
- [Golang 1.27 or higher](https://go.dev/doc/install) installed
- (optional) `make`, to use the makefile

## Goal

The goal of this workshop is to teach the value of using sentinel errors
and error types, as well as how to use them in your code.

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
