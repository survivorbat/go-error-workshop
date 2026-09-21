# ✅ Exercise 5: Testing errors

It's been a long day of automating your kitchen appliances, and you've become very
thirsty.
You walk over to your blender, get some ingredients ready, and then you realise...
You could automate that too!

You've worked hard on your implementation and are almost done, but then you realise
there are a lot of bugs in the error handling in your code.
But how is that possible?
You've written unit tests meant to catch those bugs.

## Goal

Improve the assertions in the unit tests to make sure the correct error is asserted.

**_You're not meant to fix the bugs 🪲_**

## Steps

1. Find the `Blend` function in the [blender.go](./blender.go), look for
   the **3** bugs
1. Find the unit tests for the `Blend` function in the [blender_test.go](./blender_test.go)
   They are not catching the bugs in our code. Spot the use of `assert.Error`.
1. Run the tests using `make test`, meta-tests are asserting whether
   the bugs are caught in the unit tests.
1. Update the assertions performed in [blender_test.go](./blender_test.go) to ensure
   the correct errors are returned from the functions.

## Tips

- `testify/assert` has several stronger error checks available, such as:
  - `assert.ErrorIs`
  - `assert.ErrorAs`
  - `assert.ErrorContains`
- The tests in [blender_test.go](./blender_test.go) may not show up as tests in
  your IDE. This is intentional. Please run the entire package or use `make test`
  instead.
- The `testify/require` package should not be used in this exercise, as that
  would break the test runners.
