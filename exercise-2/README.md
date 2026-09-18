# 🧼 Exercise 2: Negotiating with the dishwasher

The _Survivorbat™ WiFi-enabled Dishwasher_ SDK you're using for your
_Survivorbat™ WiFi-enabled Dishwasher_ got a major upgrade.
It now features sentinel errors and error wrapping.

You can find the [_Survivorbat™ WiFi-enabled Dishwasher_ V2 SDK here](https://github.com/survivorbat/go-error-workshop-lib/v2/dishwasher.go).

## Goal

Take your implementation from exercise 1 and adjust it to use these new errors.

## Steps

1. Run the tests using `make test`, they are meant to fail (again)
1. Find the `RunProgram` function in the [dishwasher.go](./dishwasher.go)
1. Find the unit tests for the `RunProgram` function in the [dishwasher_test.go](./dishwasher_test.go)
1. Take your implementation from exercise 1 and adjust it in exercise 2 to use the
   new sentinel errors to simplify the implementation.

## Tips

- Check out the [library's source code](https://github.com/survivorbat/go-error-workshop-lib/v2/dishwasher.go)
  to understand what errors are returned from it.
- Functions like `errors.Is`, `strings.TrimPrefix` and `strings.TrimSuffix` may be useful
