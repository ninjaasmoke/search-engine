package utils

import (
	"strings"

	stemmer "github.com/agonopol/go-stem"
	"github.com/kljensen/snowball/english"
)

func Stem(s string) string {
	return string(stemmer.Stem([]byte(s)))
}

func TokenizeText(text string) []string {
	// Split text into words
	words := strings.Fields(text)

	var tokens []string

	// filter stopwords
	for _, word := range words {
		word = strings.ToLower(word)
		if !english.IsStopWord(word) {
			// Stem the word
			stemmedWord := Stem(word)
			tokens = append(tokens, stemmedWord)
		}
	}

	return tokens
}
