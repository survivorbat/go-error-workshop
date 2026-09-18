package dishwasher

import (
	"log/slog"

	dishwashsdk "github.com/survivorbat/go-error-workshop-lib"
)

// Run will run the dishwasher program and return a message about the status
func Run(program string, intensity int) string {
	err := dishwashsdk.RunProgram(program, intensity)
	if err != nil {
		slog.Error("Call to dishwashsdk failed", "error", err)
		return explainError(err)
	}

	return "Program complete!"
}

// ❗ The code below is yours to implement.

func explainError(err error) string {
	return "implement me :)"
}
