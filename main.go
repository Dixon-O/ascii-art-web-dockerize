package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1️⃣ Ensure exactly one argument
	if len(os.Args) != 2 {
		return
	}

	// 2️⃣ Get input argument
	input := os.Args[1]

	// Convert literal "\n" sequences to actual newlines
	input = strings.ReplaceAll(input, `\n`, "\n")

	// 3️⃣ Read shadow banner file
	data, err := os.ReadFile("banners/standard.txt")
	if err != nil {
		return
	}

	// 4️⃣ Split banner into lines
	banner := strings.Split(string(data), "\n")

	// Fix shadow banner offset if first line is empty
	if banner[0] == "" {
		banner = banner[1:]
	}

	// 5️⃣ Split input into lines by actual newline
	lines := strings.Split(input, "\n")

	// 6️⃣ Print each line in ASCII art
	for _, line := range lines {

		// If the line is empty, print 8 empty rows (height of a character)
		if line == "" {
			for i := 0; i < 8; i++ {
				fmt.Println()
			}
			continue
		}

		// Print each row of ASCII art for the line
		for row := 0; row < 8; row++ {
			for _, char := range line {

				// Only printable ASCII characters
				if char < 32 || char > 126 {
					continue
				}

				// Compute the starting index for this character in the banner
				index := (int(char) - 32) * 9

				// Print the current row of the character
				fmt.Print(banner[index+row])
			}
			// Move to next row
			fmt.Println()
		}
	}
}
