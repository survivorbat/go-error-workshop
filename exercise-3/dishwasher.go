package dishwasher

import (
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
// | IntensityError | intensity 5 is invalid, it must be between 1 and 4
// | IntensityError | intensity 20 is invalid, it must be between 1 and 4
// | ProgramError   | The dishwasher does not know plan super, valid programs are basic or fast
// | ProgramError   | The dishwasher does not know plan eco, valid programs are basic or fast
func explainError(err error) string {
	return "an error happened 🤷"
}
