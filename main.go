package main

import (
	"fmt"
	"os"
	"strings"

	"ascii/ascii"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: Incorrect number of arguments ", len(os.Args))
		os.Exit(1)
	}
	// Get an input text from the commandline.
	Asciifile := os.Args[1]

	// Check if there are any args in the commandline and if the input is also empty.
	if Asciifile == "" {
		return
	}
	contascii, err := os.ReadFile(Asciifile)
	if err != nil{
		fmt.Println("Error reading from file:", err)
		os.Exit(1)
	}

	contascii2 := strings.Split(string(contascii), "\n")

	// Read content from one of the banner files.
	contents, err := os.ReadFile("standard.txt")
	// Apply error handling if file is incorrect.
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	fileInfo, err := os.Stat("standard.txt")
	if err != nil {
		fmt.Println("Error reading file information", err)
		return
	}
	fileSize := fileInfo.Size()

	if fileSize != 6623 {
		fmt.Println("Error with the file size", fileSize)
		return
	}

	// Split the content to a string slice and separate with newline.
	contents2 := strings.Split(string(contents), "\n")

	fmt.Println(ascii.AsciiArt(contascii2, contents2))
}
