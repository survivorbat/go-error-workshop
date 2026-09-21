# 🧼 Exercise 3: Telepathically channeling the dishwasher

The SDK has received another update and introduced error types, allowing you to
read data from the returned errors.

You can find the [_Survivorbat™ WiFi-enabled Dishwasher_ V3 SDK here](https://github.com/survivorbat/go-error-workshop-lib/tree/main/v3/dishwasher.go).

## Goal

Implement the `explainError` method using the new error types.

## Steps

1. Run the tests using `make test`, they are meant to fail (again)
1. Find the `RunProgram` function in the [dishwasher.go](./dishwasher.go)
1. Implement `explainError`

## Tips

- Check out the [library's source code](https://github.com/survivorbat/go-error-workshop-lib/tree/main/v3/dishwasher.go)
  to understand what errors are returned from it.
- Functions like `errors.AsType` may be useful
- `errors.As` is the older, non-generic version of `errors.AsType`
