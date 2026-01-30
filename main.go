package main

import (
	artprinter "ascii-art/Artprinter"
	"fmt"
	"os"
	"strings"
)

func main() {
	//if arguments are not as required stop program
	if len(os.Args) != 2 || os.Args[1] == "" {
		return
	}
	//handle new line characters
	input := strings.Split(strings.ReplaceAll(os.Args[1], `\n`, "\n"), "\n")
	foundNewLinesOnly := artprinter.CheckOnlyNewLines(input)
	if foundNewLinesOnly {
		for i := range input[:len(input)-1] {
			fmt.Println()
			if i == len(input[:len(input)-1])-1 {
				return
			}
		}
	}
	//read the standard.txt banner file
	data, err := os.ReadFile("banners/standard.txt")
	if err != nil {
		fmt.Println("error reading file")
		os.Exit(1)
	}
	banner := strings.Split(string(data), "\n")

	//printing the Ascii art
	fmt.Println(artprinter.PrintAsciiArt(input, banner))

}
