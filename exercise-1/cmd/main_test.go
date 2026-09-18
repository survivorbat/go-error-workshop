package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// As section is not part of the workshop, and the workshop is about the error handling
// of the program, no tests are present to test the error output. This only serves
// as an example on how the program could be tested with the run pattern.

func TestRun_AsksQuestionsAndRunsProgram(t *testing.T) {
	t.Parallel()
	// Arrange
	inputs := []string{"basic", "3"}

	// Simulate stdin and stdout
	stdin := strings.NewReader(strings.Join(inputs, "\n"))
	stdout := new(bytes.Buffer)

	// Act
	err := run(stdin, stdout)

	// Assert
	require.NoError(t, err)

	expectedLines := []string{
		"What dishwasher program do you want to run?",
		"At what intensity? (number only)",
		"Program complete!",
	}

	lineScanner := bufio.NewScanner(stdout)

	for _, expectedLine := range expectedLines {
		require.Truef(t, lineScanner.Scan(), "No more lines left, but was expecting: %s", expectedLine)
		assert.Equal(t, expectedLine, lineScanner.Text())
	}

	require.Falsef(t, lineScanner.Scan(), "was not expecting more output, got: %s", lineScanner.Text())
}
