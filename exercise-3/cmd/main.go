package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	dishwasher "github.com/survivorbat/go-error-workshop/exercise-3"
)

// 👋 Hey there! Nice to find you here. This code isn't part of the workshop, but feel
// free to have a look. The main/run functions here exemplify the run pattern:
//
// https://dev.to/mokiat/go-main-run-pattern-1bin
//
// It allows us to actually test the run method in main_test.go, without having to worry
// about the os.Exit or the stdin/stdout.

func main() {
	err := run(os.Stdin, os.Stdout)
	if err != nil {
		slog.Error("Failed to run program", "error", err)
		os.Exit(1)
	}
}

// run acts as a main() with an error return and global steam inputs. This allows us
// to test the program almost in its entirety in main_test.go. All interactions with
// stdin and stdout are abstracted away, allowing us to input buffers from a test.
func run(stdin io.Reader, stdout io.Writer) error {
	_, _ = fmt.Fprintln(stdout, "What dishwasher program do you want to run?")

	var inputProgram string
	_, err := fmt.Fscan(stdin, &inputProgram)
	if err != nil {
		return fmt.Errorf("failed to read program name: %w", err)
	}

	_, _ = fmt.Fprintln(stdout, "At what intensity? (number only)")

	var inputIntensity int
	_, err = fmt.Fscan(stdin, &inputIntensity)
	if err != nil {
		return fmt.Errorf("failed to read intensity: %w", err)
	}

	_, _ = fmt.Fprintln(stdout, dishwasher.Run(inputProgram, inputIntensity))

	return nil
}
