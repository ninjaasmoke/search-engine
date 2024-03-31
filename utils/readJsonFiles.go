package utils

import (
	"encoding/json"
	"os"
	"search-server/types"
)

// Function to read and parse the JSON file
func ReadDocumentInfoJson(filename string) (map[string]types.ImageData, error) {
	// Read the JSON file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Define a map to hold the dynamic IDs and ImageData
	var docInfoMap map[string]types.ImageData

	// Unmarshal JSON data into map
	err = json.Unmarshal(data, &docInfoMap)
	if err != nil {
		return nil, err
	}

	return docInfoMap, nil
}

func ReadInvertedIndexJson(filename string) (map[string]types.InvertedIndex, map[string]int, int, error) {
	// Read the JSON file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, nil, 0, err
	}

	// Define a map to hold the inverted index data
	var invertedMap map[string]types.InvertedIndex

	// Unmarshal JSON data into map
	err = json.Unmarshal(data, &invertedMap)
	if err != nil {
		return nil, nil, 0, err
	}

	// Calculate document frequency and total number of documents
	docFreq := make(map[string]int)
	for token, info := range invertedMap {
		docFreq[token] = info.DF
	}

	tD := 0
	for _, df := range docFreq {
		tD += df
	}

	return invertedMap, docFreq, tD, nil
}
