package dishwasher

import (
	"fmt"
	"strings"

	dishwashsdk "github.com/survivorbat/go-error-workshop-lib"
)

// Run will run the dishwasher program and return a message about the status
func Run(program string, intensity int) string {
	err := dishwashsdk.RunProgram(program, intensity)
	if err != nil {
		return explainError(err)
	}

	return "Program complete!"
}

// explainError translates the error to the output
//
// ❗ Test requirements:
//
// | error                                                           | Expected output
// | intensity must be a positive number                             | Intensity must be a positive number
// | intensity 5 is too great, the maximum is 4                      | Intensity 5 is invalid, it must be between 1 and 4
// | intensity 20 is too great, the maximum is 4                     | Intensity 20 is invalid, it must be between 1 and 4
// | program super does not exist, allowed programs are: fast, basic | The dishwasher does not know plan super, valid programs are basic or fast
// | program eco does not exist, allowed programs are: fast, basic   | The dishwasher does not know plan eco, valid programs are basic or fast
func explainError(err error) string {
	msg := err.Error()

	if msg == "intensity must be a positive number" {
		return "Intensity must be a positive number"
	}

	segments := strings.Split(msg, " ")

	if strings.Contains(msg, "is too great") {
		return fmt.Sprintf("Intensity %s is invalid, it must be between 1 and 4", segments[1])
	}

	if strings.Contains(msg, "does not exist") {
		return fmt.Sprintf("The dishwasher does not know plan %s, valid programs are basic or fast", segments[1])
	}

	panic("unknown error")
}
