# ❄️ Exercise 4: Keeping it cool

You finally got your dishwasher operational from your tool, and it feels great.
But you have more appliances to automate, and next up is your
_Survivorbat™ Ethernet-enabled Fridge_.
Unfortunately your fridge has no SDK available, so it's up to you to communicate
with its API directly.
You even want to publish this library on your personal GitHub profile, so it has
to be top-notch.

Having learned from the mistakes of the dishwasher SDK, you want to
make proper error handling a priority.

## Goal

Implement the `explainError` method using the new error types.

## Steps

1. Run the tests using `make test`, they are meant to fail (again)
1. Find the `RunProgram` function in the [dishwasher.go](./dishwasher.go)
1. Find the unit tests for the `RunProgram` function in the [dishwasher_test.go](./dishwasher_test.go)
1. Take your implementation from exercise 1 and adjust it in exercise 2 to use the
   new sentinel errors to simplify the implementation.

## Tips

- Check out the [library's source code](https://github.com/survivorbat/go-error-workshop-lib/v3/dishwasher.go)
  to understand what errors are returned from it.
- Functions like `errors.AsType` may be useful
- `errors.As` is the older, non-generic version of `errors.AsType`
