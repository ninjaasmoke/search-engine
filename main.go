package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/kljensen/snowball/english"
)

type ImageData struct {
	URL                string   `json:"url"`
	Title              string   `json:"title"`
	RelatedImageTags   []string `json:"related_image_tags"`
	AnnotatedImageTags []string `json:"annotated_image_tags"`
}

type InvertedIndex struct {
	TF map[string]int `json:"tf"`
	DF int            `json:"df"`
}

var (
	documentInfoMap   map[string]ImageData
	mapLock           sync.RWMutex
	invertedIndexMap  map[string]InvertedIndex
	invertedIndexLock sync.RWMutex
	totalDocs         int
	documentFrequency map[string]int
)

func tokenizeText(text string) []string {
	// Split text into words
	words := strings.Fields(text)

	var tokens []string

	// Lemmatize and filter stopwords
	for _, word := range words {
		word = strings.ToLower(word)
		if !english.IsStopWord(word) {
			// Lemmatize the word
			word = english.Stem(word, true)
			tokens = append(tokens, word)
		}
	}

	return tokens
}

// Function to read and parse the JSON file
func readDocumentInfoJson(filename string) error {
	// Read the JSON file
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// Define a map to hold the dynamic IDs and ImageData
	var docInfoMap map[string]ImageData

	// Unmarshal JSON data into map
	err = json.Unmarshal(data, &docInfoMap)
	if err != nil {
		return err
	}

	// Update global variable with parsed data
	mapLock.Lock()
	defer mapLock.Unlock()
	documentInfoMap = docInfoMap

	return nil
}

func readInvertedIndexJson(filename string) error {
	// Read the JSON file
	data, err := os.ReadFile(filename)

	if err != nil {
		return err
	}

	var invertedMap map[string]InvertedIndex

	err = json.Unmarshal(data, &invertedMap)
	if err != nil {
		return err
	}

	docFreq := make(map[string]int)
	for token, info := range invertedMap {
		docFreq[token] = info.DF
	}
	tD := 0
	for _, df := range docFreq {
		tD += df
	}

	documentFrequency = docFreq
	totalDocs = tD

	invertedIndexLock.Lock()
	defer invertedIndexLock.Unlock()
	invertedIndexMap = invertedMap

	return nil
}

func imageDataHandler(w http.ResponseWriter, r *http.Request) {
	// Extract ID from the request URL
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]

	mapLock.RLock()
	defer mapLock.RUnlock()

	// Retrieve ImageData for the given ID
	imageData, found := documentInfoMap[id]
	if !found {
		http.NotFound(w, r)
		return
	}

	// Marshal ImageData to JSON
	responseJSON, err := json.Marshal(imageData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error marshaling JSON:", err)
		return
	}

	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write response
	w.Write(responseJSON)
}

// a seach handler function that acccepts a query and returns the relevant images
// it is a GET /search?query=... endpoint
/*

query_terms = tokenize_text(query)
idf = {term: math.log(total_docs / (document_frequency.get(term, 0) + 1)) for term in query_terms}

document_scores = {}
for term in query_terms:
    if term in final_inverted_index:
        for doc_id, tf in final_inverted_index[term]['tf'].items():
            if doc_id not in document_scores:
                document_scores[doc_id] = 0
            document_scores[doc_id] += tf * idf[term]

# Rank documents based on combined TF-IDF scores
ranked_documents = sorted(document_scores.items(), key=lambda x: x[1], reverse=True)

return top 20 ranked_documents (image_data)
*/

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	queryTerms := tokenizeText(query)

	idf := make(map[string]float64)
	for _, term := range queryTerms {
		idf[term] = math.Log(float64(totalDocs) / float64(documentFrequency[term]+1))
	}

	documentScores := make(map[string]float64)
	for _, term := range queryTerms {
		if info, found := invertedIndexMap[term]; found {
			for docID, tf := range info.TF {
				if _, found := documentScores[docID]; !found {
					documentScores[docID] = 0
				}
				documentScores[docID] += float64(tf) * idf[term]
			}
		}
	}

	rankedDocuments := make([]string, 0, len(documentScores))
	for docID := range documentScores {
		rankedDocuments = append(rankedDocuments, docID)
	}

	sort.Slice(rankedDocuments, func(i, j int) bool {
		return documentScores[rankedDocuments[i]] > documentScores[rankedDocuments[j]]
	})

	// Return top 20 documents
	if len(rankedDocuments) > 20 {
		rankedDocuments = rankedDocuments[:20]
	}

	// Retrieve ImageData for the top 20 documents
	mapLock.RLock()
	defer mapLock.RUnlock()

	var imageDatas []ImageData
	for _, docID := range rankedDocuments {
		imageData, found := documentInfoMap[docID]
		if found {
			imageDatas = append(imageDatas, imageData)
		}
	}

	// Marshal ImageData to JSON
	responseJSON, err := json.Marshal(imageDatas)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error marshaling JSON:", err)
		return
	}

	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write response
	w.Write(responseJSON)

}

func frontendHandler(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("./front/dist")).ServeHTTP(w, r)

}

func main() {
	docInfoFile := "document_info_map.json"
	invertedIndexFile := "final_inverted_index.json"

	if err := readDocumentInfoJson(docInfoFile); err != nil {
		log.Fatal(err)
	}

	if err := readInvertedIndexJson(invertedIndexFile); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	// Define API routes with "/api" prefix
	mux.HandleFunc("/api/imageData/", imageDataHandler)
	mux.HandleFunc("/api/search", searchHandler)

	// Define a route for the root URL
	mux.HandleFunc("/", frontendHandler)

	// Start HTTP server
	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
