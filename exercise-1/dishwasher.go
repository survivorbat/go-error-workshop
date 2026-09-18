package dishwasher

import dishwashsdk "gitnub.com/survivorbat/go-error-workshop-lib"

// Run will run the dishwasher program and return a message about the status
func Run(program string, intensity int) string {
	err := dishwashsdk.RunProgram(program, intensity)
	if err != nil {
		return explainError(err)
	}

	return "Program complete!"
}

func explainError(err error) string {
	panic("implement me :)")
}
