package artprinter

// CheckOnlyNewLines() checks if input has only new line characters
func CheckOnlyNewLines(input []string) bool {
	found := true
	for _, ch := range input {
		if ch != "" {
			found = false
		}
	}
	return found
}
