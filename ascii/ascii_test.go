package ascii

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestAsciiArt(t *testing.T) {
	type parameters struct {
		words     []string
		contents2 []string
	}
	type testCases struct {
		name string
		args parameters
	}
	// We read the bannerfile directly to save on space/make the code readable
	// otherwise we would have provided the contents in banner file as a []string.
	contents, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Print("Error reading from file", err)
	}
	textArt := strings.Split(string(contents), "\n")

	tests := []testCases{
		{name: "test empty", args: parameters{words: []string{}, contents2: textArt}},
		{name: "test newline(\n)", args: parameters{words: []string{""}, contents2: textArt}},
		{name: "test Non ascii", args: parameters{words: []string{"你好"}, contents2: textArt}},
		{name: "test HeL10#\n'", args: parameters{words: []string{"'HeL10#'", ""}, contents2: textArt}},
		{name: "test {3Hello !@There}", args: parameters{words: []string{"{3Hello !@There}"}, contents2: textArt}},
		{name: "test Hello\nThere", args: parameters{words: []string{"Hello", "There"}, contents2: textArt}},
		{name: "test Hello\n\nThere", args: parameters{words: []string{"Hello", "", "There"}, contents2: textArt}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AsciiArt(tt.args.words, tt.args.contents2)
		})
	}
}
