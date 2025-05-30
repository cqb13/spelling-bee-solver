package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

//go:embed words.txt
var words string

var centerLetter rune
var outerLetters []rune

func validWord(word string) bool {
	if !strings.Contains(word, string(centerLetter)) {
		return false
	}

	for _, char := range word {
		if char == centerLetter {
			continue
		}

		if !slices.Contains(outerLetters, char) {
			return false
		}
	}

	return true
}

func main() {
	centerLetterPtr := flag.String("center", "-", "required letter")
	outerLettersPtr := flag.String("outer", "-", "options for included letters")

	flag.Parse()

	if *centerLetterPtr == "-" {
		fmt.Println("Center letter must be provided")
		fmt.Println()
		flag.Usage()
		os.Exit(1)
	}

	if len(*centerLetterPtr) != 1 {
		fmt.Println("Center letter must only be 1 letter")
		fmt.Println()
		flag.Usage()
		os.Exit(1)
	}

	if *outerLettersPtr == "-" {
		fmt.Println("Outer letters must be provided")
		fmt.Println()
		flag.Usage()
		os.Exit(1)
	}

	if len(*outerLettersPtr) != 6 {
		fmt.Println("There must be 6 outer letters")
		fmt.Println()
		flag.Usage()
		os.Exit(1)
	}

	centerLetter = rune((*centerLetterPtr)[0])

	for _, letter := range *outerLettersPtr {
		outerLetters = append(outerLetters, letter)
	}

	var word int = 1
	for line := range strings.SplitSeq(words, "\n") {
		parts := strings.Split(line, ",")

		if !validWord(parts[0]) {
			continue
		}

		fmt.Printf("%-3v| %-20v%v\n", word, parts[0], parts[1])
		word++
	}
}
