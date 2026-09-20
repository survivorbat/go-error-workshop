# ❄️ Exercise 4: Keeping it cool

You finally got your dishwasher operational from your tool, and it feels great.
But you have more appliances to automate, and next up is your
_Survivorbat™ Ethernet-enabled Fridge_.
Unfortunately your fridge has no SDK available, so it's up to you to communicate
with its API directly.
You even want to publish this library on your personal GitHub profile, so it has
to be top-notch in terms of error handling.

Your first priority is to implement the `Configure` functionality.
This will allow you to set the name and temperature of the fridge.

## Goal

Implement the `Configure` method on the `FridgeClient` so that all tests pass.

## 1. Fixing the errors in errors

### Steps

1. Find the `Configure` method in the [fridge.go](./fridge.go). It has comments
   explaining what the test requirements are in the error paths. We will
   work on these in a moment.
1. Find the library's errors in [errors.go](./errors.go)
1. Find the tests for these errors in [errors_test.go](./errors_test.go)
1. Run `make test.errors`, these should fail
1. Add the necessary code to the errors to get the tests to succeed

### Tips

- What makes an error an error? Check out the stdlib's `error` type if you're stuck
- Check out the [stdlib's errors package documentation](https://pkg.go.dev/errors)
  to learn what other methods can be implemented to change error behaviour

## 2. Implementing the errors from errors

### Steps

1. Take another look at the `Configure` method in [fridge.go](./fridge.go).
   Some of the test requirements may look more familiar now.
1. Run `make test.all`, these should fail
1. Implement the error returns according to the specified requirements

### Tips

- Use your implementations in [errors.go](./errors.go)
- You don't need to create new errors to finish this exercise
- Remember error wrapping and error joining
