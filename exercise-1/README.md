# 🧼 Exercise 1: Talking to the dishwasher

You bought a new _Survivorbat™ WiFi-enabled Dishwasher_, congrats!
You want to run it from your computer, and decided to write a custom Go program
to instrument it.
Luckily for you, the dishwasher has its own SDK available, allowing you contact
its API by calling methods on a Golang library.

You can find the [_Survivorbat™ WiFi-enabled Dishwasher_ SDK here](https://github.com/survivorbat/go-error-workshop-lib/tree/maindishwasher.go).

You've been working hard on your first Proof-of-Concept, and the only step
left for you is to properly display errors in your application.
This is where the `explainError` function comes into play.

## Goal

Get the tests for the `Run` function working.

## Steps

1. Run the program using `make run`, follow the prompts and see what happens
1. Run the tests using `make test`, they are meant to fail
1. Find the `RunProgram` function in the [dishwasher.go](./dishwasher.go)
1. Finish the implementation of `explainError` to make the unit tests succeed.

## Tips

- Check out the [library's source code](https://github.com/survivorbat/go-error-workshop-lib/tree/main/dishwasher.go)
  to understand what errors are returned from it.
- Functions like `strings.Contains` and `strings.Split` may be useful
