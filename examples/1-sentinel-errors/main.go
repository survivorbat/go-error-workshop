package main

import (
	"errors"
	"fmt"
)

var ErrFoo = errors.New("foo")

func main() {
	fmt.Println("------------------------------------------")

	fmt.Println("Are 2 ErrFoo's equal with ==?")
	fmt.Println(`ErrFoo == ErrFoo`, ErrFoo == ErrFoo)

	fmt.Println()
	fmt.Println("------------------------------------------")
	fmt.Println()

	fmt.Println("Are 2 ErrFoo's equal with ==?")
	fmt.Println(`ErrFoo == ErrFoo`, ErrFoo == ErrFoo)

	fmt.Println()
	fmt.Println("------------------------------------------")
}
