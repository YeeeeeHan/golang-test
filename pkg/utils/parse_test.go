package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInput_Commands(t *testing.T) {
	testTable := []struct {
		name          string
		input         string
		outputCommand string
		outputArgs    []string
	}{
		{
			name:          "parse empty input",
			input:         "",
			outputCommand: "",
			outputArgs:    []string(nil),
		},
		{
			name:          "parse only command",
			input:         "login",
			outputCommand: "login",
			outputArgs:    []string{},
		},
		{
			name:          "parse erratic spaces",
			input:         "  login   abc   def",
			outputCommand: "login",
			outputArgs:    []string{"abc", "def"},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			outputCommand, outputArgs := ParseInput(tt.input)
			assert.Equal(t, tt.outputCommand, outputCommand)
			assert.Equal(t, tt.outputArgs, outputArgs)
		})
	}
}
