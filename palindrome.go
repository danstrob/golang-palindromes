package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

// readTextfile reads a file given as a CL argument and returns its contents as a string
func readTextfile() string {
	filename := string(os.Args[1])
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return string(contents)
}

// letterMap keeps only unicode letters in a string
func letterMap(str string) string {
	return strings.Map(func(rune rune) rune {
		if unicode.IsLetter(rune) {
			return rune
		}
		return -1
	}, str)
}

// cleanString takes in a string and returns a list of strings
// without line breaks, spaces and non-letter chars, so hopefully a list of words
func cleanString(str string) []string {
	words := strings.Split(strings.Replace(str, "\n", " ", -1), " ")

	for i, w := range words {
		words[i] = strings.ToLower(letterMap(w))
	}
	return words
}

// reverseString works by flipping the runes in a string
func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	wordbag := cleanString(readTextfile())
	var palindromes []string

	for _, w := range wordbag {
		if w == reverseString(w) && len(w) > 1 {
			palindromes = append(palindromes, w)
		}
	}

	if len(palindromes) == 1 {
		fmt.Println("I have found 1 palindrome in your text file. It is:", palindromes)
	} else {
		fmt.Printf("I have found %d palindromes in your text file. They are:\n", len(palindromes))
		for i := 0; i < len(palindromes)-1; i++ {
			fmt.Printf("%s, ", palindromes[i])
		}

		fmt.Printf("%s.\n", palindromes[len(palindromes)-1])
	}
}
