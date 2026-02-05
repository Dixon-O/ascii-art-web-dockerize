package artprinter

import "strings"

// PrintAsciiArt returns a string with the ascii art
func PrintAsciiArt(input, banner []string) string {
	// final := ""
	var final strings.Builder
	//printing the asci art
	for _, word := range input {
		if word == "" {
			final.WriteString("\n")
			continue
		}
		for row := 0; row < 8; row++ {
			for _, char := range word {
				startIndex := (int(char)-32)*9 + 1
				final.WriteString(string(banner[startIndex+row]))
			}
			if row < 8 {
				final.WriteString("\n")
			}
		}
	}
	// return final[:len(final)-1]
	return final.String()
}
