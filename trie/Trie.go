package trie

import (
	"strings"
)

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			node.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[char]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			return false
		}
		node = node.children[char]
	}
	return node.isEnd
}

func (t *Trie) SpellCheck(query string) (string, bool) {
	words := strings.Fields(query)
	correctedWords := make([]string, len(words))

	// Check each word for spelling errors
	for i, word := range words {
		if !t.Search(word) {
			correctedWord := t.findClosestWord(word)

			if correctedWord != "" {
				correctedWords[i] = correctedWord
			} else {
				correctedWords[i] = word
			}
		} else {
			correctedWords[i] = word
		}
	}

	correctedQuery := strings.Join(correctedWords, " ")

	return correctedQuery, !strings.EqualFold(query, correctedQuery)
}

func (t *Trie) findClosestWord(word string) string {
	minDistance := len(word)
	closestWord := ""

	var search func(node *TrieNode, currentWord string, depth, distance int)
	search = func(node *TrieNode, currentWord string, depth, distance int) {
		if distance > minDistance {
			return
		}
		if node.isEnd && depth == len(word) && distance < minDistance {
			minDistance = distance
			closestWord = currentWord
		}
		if depth == len(word) {
			return
		}
		for char, child := range node.children {
			if char == rune(word[depth]) {
				search(child, currentWord+string(char), depth+1, distance)
			} else {
				search(child, currentWord+string(char), depth+1, distance+1)
			}
		}
	}

	search(t.root, "", 0, 0)
	return closestWord
}
