package artprinter

// PrintAsciiArt returns a string with the ascii art
func PrintAsciiArt(input, banner []string) string {
	final := ""
	//printing the asci art
	for _, word := range input {
		if word == "" {
			final += "\n"
			continue
		}
		for row := 0; row < 8; row++ {
			for _, char := range word {
				startIndex := (int(char)-32)*9 + 1
				final += string(banner[startIndex+row])
			}
			if row < 8 {
				final += "\n"
			}
		}
	}
	return final[:len(final)-1]
}
