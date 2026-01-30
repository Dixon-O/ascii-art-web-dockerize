package artprinter

import (
	"os"
	"strings"
	"testing"
)

// TestCharIndexCalculaton tests character index calculation
func TestCharIndexCalculaton(t *testing.T) {
	tests := []struct {
		char rune
		want int
	}{
		{' ', 0},
		{'!', 9},
		{'0', 144},
		{'A', 297},
		{'a', 585},
		{'~', 846},
	}

	for _, tt := range tests {
		got := (int(tt.char) - 32) * 9
		if got != tt.want {
			t.Errorf("Index for %c: got %d, want %d", tt.char, got, tt.want)
		}
	}
}

func TestPrintAsciiArt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"words separated by `\n` character", "hello\nthere", ` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
 _     _                           
| |   | |                          
| |_  | |__     ___   _ __    ___  
| __| |  _ \   / _ \ | '__|  / _ \ 
\ |_  | | | | |  __/ | |    |  __/ 
 \__| |_| |_|  \___| |_|     \___| 
                                   
                                   `},
		{"single word", "hello", ` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               `},
		{"special characters", "@#!%^$*^", "             _  _     _   _   __  /\\    _       _      /\\  \n   ____    _| || |_  | | (_) / / |/\\|  | |   /\\| |/\\  |/\\| \n  / __ \\  |_  __  _| | |    / /       / __)  \\ ` ' /       \n / / _` |  _| || |_  | |   / /        \\__ \\ |_     _|      \n| | (_| | |_  __  _| |_|  / / _       (   /  / , . \\       \n \\ \\__,_|   |_||_|   (_) /_/ (_)       |_|   \\/|_|\\/       \n  \\____/                                                   \n                                                           "},
	}
	banner, err := os.ReadFile("../banners/standard.txt")
	if err != nil {
		t.Fatalf("Unexpected error: %v reading file", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := strings.Split(strings.ReplaceAll(tt.input, `\n`, "\n"), "\n")
			result := PrintAsciiArt(str, strings.Split(string(banner), "\n"))
			if result != tt.expected {
				t.Errorf("PrintAsciiArt() failed for input %q. Got:\n%q\nWant:\n%q",
					tt.input, result, tt.expected)
			}
		})
	}
}
