package program

import (
	"fmt"
	"strings"
)

func PrettyPrintHashtable(input string) (string, error) {
	// Remove curly braces and extra spaces
	input = strings.TrimSpace(strings.Trim(input, "{}"))

	// Split key-value pairs
	pairs := strings.Split(input, ",")
	hashtable := make(map[int]int)

	// Parse key-value pairs
	for _, pair := range pairs {
		var key, value int
		_, err := fmt.Sscanf(pair, "%d:%d", &key, &value)
		if err != nil {
			continue
		}
		hashtable[key] = value
	}

	// Pretty print the hashtable
	var output strings.Builder
	output.WriteString("{\n")
	for key, value := range hashtable {
		output.WriteString(fmt.Sprintf("  %d: %d,\n", key, value))
	}
	output.WriteString("}")

	return output.String(), nil
}
