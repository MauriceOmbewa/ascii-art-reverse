package ascii

import (
	"strings"
)

// AsciiArt processes words, printing their ASCII art
// character by character and adding new lines as needed.

// AsciiArtToText transforms ASCII art back to text
func AsciiArt(asciiArt []string, contents2 []string) string {
	if len(asciiArt) == 0 {
		return ""
	}

	// Map ASCII art to corresponding characters
	artToChar := make(map[string]rune)
	for i := 0; i < 95; i++ {
		charArt := ""
		for j := 0; j < 8; j++ {
			charArt += contents2[i*9+1+j]
		}
		artToChar[charArt] = rune(i + 32)
	}

	var result strings.Builder
	for i := 0; i < len(asciiArt); i += 8 {
		charArt := ""
		for j := 0; j < 8; j++ {
			if i+j < len(asciiArt) {
				charArt += asciiArt[i+j]
			}
		}
		char := artToChar[charArt];
		result.WriteRune(char)
		// if char, exists := artToChar[charArt]; exists {
		// 	result.WriteRune(char)
		// } else {
		// 	fmt.Println("Error: Unrecognized ASCII art pattern")
		// 	return ""
		// }
	}

	return result.String()
}