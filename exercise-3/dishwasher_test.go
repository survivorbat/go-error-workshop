package dishwasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		program   string
		intensity int
		expected  string
	}{
		"successful program execution": {
			program:   "basic",
			intensity: 2,
			expected:  "Program complete!",
		},
		"invalid intensity 5": {
			program:   "basic",
			intensity: 5,
			expected:  "Intensity 5 is invalid, it must be between 1 and 4",
		},
		"invalid intensity 20": {
			program:   "basic",
			intensity: 20,
			expected:  "Intensity 20 is invalid, it must be between 1 and 4",
		},
		"invalid program super": {
			program:   "super",
			intensity: 3,
			expected:  "The dishwasher does not know plan super, valid programs are basic or fast",
		},
		"invalid program eco": {
			program:   "eco",
			intensity: 3,
			expected:  "The dishwasher does not know plan eco, valid programs are basic or fast",
		},
	}

	for name, testData := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// Act
			actual := Run(testData.program, testData.intensity)

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}
