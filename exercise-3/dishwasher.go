package dishwasher

import (
	"errors"
	"fmt"

	dishwashsdk "github.com/survivorbat/go-error-workshop-lib/v3"
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
// | error          | Expected output
// | IntensityError | Intensity 5 is invalid, it must be between 1 and 4
// | IntensityError | Intensity 20 is invalid, it must be between 1 and 4
// | ProgramError   | The dishwasher does not know plan super, valid programs are basic or fast
// | ProgramError   | The dishwasher does not know plan eco, valid programs are basic or fast
func explainError(err error) string {
	intensityErr, ok := errors.AsType[*dishwashsdk.IntensityError](err)
	if ok {
		return fmt.Sprintf("Intensity %d is invalid, it must be between 1 and %d", intensityErr.Value, intensityErr.Maximum)
	}

	programErr, ok := errors.AsType[*dishwashsdk.ProgramError](err)
	if ok {
		return fmt.Sprintf("The dishwasher does not know plan %s, valid programs are basic or fast", programErr.Value)
	}

	panic("unknown error")
}
