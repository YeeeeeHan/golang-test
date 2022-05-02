package utils

import "strings"

// ParseInput Takes inputString and splits into command and arguments
func ParseInput(inputString string) (string, []string) {
	if inputString == "" {
		return "", nil
	}
	// Splits by space, handles double spacing
	stringSlice := strings.Fields(inputString)
	for i, str := range stringSlice {
		// Standardise input by converting to lowercase
		stringSlice[i] = strings.ToLower(str)
	}

	return stringSlice[0], stringSlice[1:]
}
