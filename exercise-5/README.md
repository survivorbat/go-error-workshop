# ✅ Exercise 5: Testing errors

It's been a long day of automating your

## Goal

Implement the `explainError` method using the new error types.

## Steps

1. Run the tests using `make test`, they are meant to fail (again)
1. Find the `RunProgram` function in the [dishwasher.go](./dishwasher.go)
1. Find the unit tests for the `RunProgram` function in the [dishwasher_test.go](./dishwasher_test.go)
1. Take your implementation from exercise 1 and adjust it in exercise 2 to use the
   new sentinel errors to simplify the implementation.

## Tips

- Check out the [library's source code](https://github.com/survivorbat/go-error-workshop-lib/tree/main/v3/dishwasher.go)
  to understand what errors are returned from it.
- Functions like `errors.AsType` may be useful
- `errors.As` is the older, non-generic version of `errors.AsType`
