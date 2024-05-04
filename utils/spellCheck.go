package utils

import (
	"bufio"
	"os"
	"search-server/trie"
	"strings"
)

func LoadWords(filename string, trie *trie.Trie) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		words = append(words, word)
		trie.Insert(strings.ToLower(word))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}
