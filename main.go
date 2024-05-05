package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"search-server/api"
	"search-server/models"
	"search-server/types"
	"search-server/utils"
)

// CorsMiddleware adds CORS headers to every request
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")                            // Allow requests from any origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")          // Allow the HTTP methods you need
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Allow the headers you need
		w.Header().Set("Access-Control-Allow-Credentials", "true")                    // Allow credentials such as cookies (if your client sends them)

		// Check if it's a preflight request
		if r.Method == "OPTIONS" {
			// Preflight requests need to be handled with a 200 OK response
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func main() {
	docInfoFile := "document_info_map.json"
	invertedIndexFile := "final_inverted_index.json"
	wordListFile := "words.txt"

	// Load words from the text file
	trie := models.NewTrie()
	_, err := utils.LoadWords(wordListFile, trie)
	if err != nil {
		fmt.Println("Error loading words:", err)
		return
	}

	docInfoMap, err := utils.ReadDocumentInfoJson(docInfoFile)
	if err != nil {
		log.Fatal(err)
	}

	invertedIndexMap, docFreq, totalDocs, err := utils.ReadInvertedIndexJson(invertedIndexFile)
	if err != nil {
		log.Fatal(err)
	}

	var appData types.JsonData

	appData.DocumentInfoMap = docInfoMap
	appData.InvertedIndexMap = invertedIndexMap
	appData.TotalDocs = totalDocs
	appData.DocumentFrequency = docFreq
	appData.Trie = trie

	// documentVectors := utils.GenerateDocumentVectors(appData)
	// appData.DocumentVectors = documentVectors

	appData.AveraageDocLength = utils.GetAverageDocumentLength(docInfoMap)

	mux := http.NewServeMux()
	ctx := context.WithValue(context.Background(), types.AppDataKey{}, appData)

	// Define API routes with "/api" prefix
	mux.HandleFunc("/api/ping/", http.HandlerFunc(api.PingHandler))
	mux.HandleFunc("/api/imageData/", func(w http.ResponseWriter, r *http.Request) {
		api.ImageDataHandler(w, r.WithContext(ctx))
	})
	mux.HandleFunc("/api/search", func(w http.ResponseWriter, r *http.Request) {
		api.SearchHandler(w, r.WithContext(ctx), false)
	})
	mux.HandleFunc("/api/search/noCheck", func(w http.ResponseWriter, r *http.Request) {
		api.SearchHandler(w, r.WithContext(ctx), true)
	})

	// Define a route for the root URL
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		api.FrontendHandler(w, r.WithContext(ctx))
	})

	wrappedMux := CorsMiddleware(mux)

	// Start HTTP server
	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", wrappedMux))
}
