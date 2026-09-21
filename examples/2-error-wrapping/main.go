package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("------------------------------------------")
	fmt.Println()

	fmt.Println("Are 2 errors.New equal with ==?")
	fmt.Println(`errors.New("foo") == errors.New("foo") =`, errors.New("foo") == errors.New("foo"))

	fmt.Println()
	fmt.Println("------------------------------------------")
	fmt.Println()

	fmt.Println("Are their .Error() strings equal with ==?")
	fmt.Println(`errors.New("foo").Error() == errors.New("foo").Error() =`, errors.New("foo").Error() == errors.New("foo").Error())

	fmt.Println()
	fmt.Println("------------------------------------------")
	fmt.Println()

	fmt.Println("What if we use errors.Is?")
	fmt.Println(`errors.Is(errors.New("foo"), errors.New("foo")) =`, errors.Is(errors.New("foo"), errors.New("foo")))

	fmt.Println()
	fmt.Println("------------------------------------------")
	fmt.Println()

	fmt.Println("What if I declare the error beforehand, are they equal with ==?")
	fmt.Println(`errFoo := errors.New("foo")`)

	errFoo := errors.New("foo")
	fmt.Println("errFoo == errFoo = ", errFoo == errFoo)

	fmt.Println()
	fmt.Println("------------------------------------------")
	fmt.Println()

	fmt.Println("What about errors.Is?")
	fmt.Println("errors.Is(errFoo, errFoo) =", errors.Is(errFoo, errFoo))

	fmt.Println()
	fmt.Println("------------------------------------------")
	fmt.Println()

	fmt.Println("A sentinel error is an error that you declare beforehand at the top of the package.")
	fmt.Println("This allows you to both use == and errors.Is to verify what you received is a specific error.")
	fmt.Println("What's the benefit of errors.Is over ==? You'll see that in the next examples.")

	fmt.Println()
	fmt.Println("------------------------------------------")
}
