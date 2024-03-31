package api

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"search-server/types"
	"search-server/utils"
)

// func SearchHandler(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	appData, ok := ctx.Value(types.AppDataKey{}).(types.JsonData)

// 	if !ok {
// 		http.Error(w, "Failed to retrieve appData", http.StatusInternalServerError)
// 		return
// 	}

// 	query := r.URL.Query().Get("q")

// 	queryTerms := utils.TokenizeText(query)

// 	idf := make(map[string]float64)
// 	for _, term := range queryTerms {
// 		idf[term] = math.Log(float64(appData.TotalDocs) / float64(appData.DocumentFrequency[term]+1))
// 	}

// 	documentScores := make(map[string]float64)
// 	for _, term := range queryTerms {
// 		if info, found := appData.InvertedIndexMap[term]; found {
// 			for docID, tf := range info.TF {
// 				if _, found := documentScores[docID]; !found {
// 					documentScores[docID] = 0
// 				}
// 				documentScores[docID] += float64(tf) * idf[term]
// 			}
// 		}
// 	}

// 	rankedDocuments := make([]string, 0, len(documentScores))
// 	for docID := range documentScores {
// 		rankedDocuments = append(rankedDocuments, docID)
// 	}

// 	sort.Strings(rankedDocuments)

// 	documentTitles := make(map[string]bool)
// 	var uniqueRankedDocuments []string

// 	for _, docID := range rankedDocuments {
// 		imageData, found := appData.DocumentInfoMap[docID]
// 		if found {
// 			if _, exists := documentTitles[imageData.Title]; !exists {
// 				documentTitles[imageData.Title] = true
// 				uniqueRankedDocuments = append(uniqueRankedDocuments, docID)
// 				if len(uniqueRankedDocuments) >= 30 {
// 					break
// 				}
// 			}
// 		}
// 	}

// 	// Retrieve ImageData for the top 30 unique documents
// 	var uniqueImageDatas []types.ImageData
// 	for _, docID := range uniqueRankedDocuments {
// 		imageData, found := appData.DocumentInfoMap[docID]
// 		if found {
// 			uniqueImageDatas = append(uniqueImageDatas, imageData)
// 		}
// 	}

// 	// Marshal ImageData to JSON
// 	responseJSON, err := json.Marshal(uniqueImageDatas)
// 	if err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		log.Println("Error marshaling JSON:", err)
// 		return
// 	}

// 	// Set Content-Type header
// 	w.Header().Set("Content-Type", "application/json")

// 	// Write response
// 	w.Write(responseJSON)

// }

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	appData, ok := ctx.Value(types.AppDataKey{}).(types.JsonData)

	if !ok {
		http.Error(w, "Failed to retrieve appData", http.StatusInternalServerError)
		return
	}

	// Tokenize and preprocess the query
	query := r.URL.Query().Get("q")
	queryTerms := utils.TokenizeText(query)

	// Calculate TF-IDF for query terms
	queryVector := make(map[string]float64)
	for _, term := range queryTerms {
		queryVector[term] = math.Log(float64(appData.TotalDocs) / float64(appData.DocumentFrequency[term]+1))
	}

	// Rank documents based on cosine similarity
	rankedDocuments := utils.RankDocuments(queryVector, appData.DocumentVectors)

	topDocuments := rankedDocuments[:60]

	// Create a map to store unique URLs
	uniqueURLs := make(map[string]struct{})

	// Slice to store unique ImageData entries
	uniqueImageDatas := make([]types.ImageData, 0, len(topDocuments))

	// Iterate over topDocuments to filter out duplicates based on URL
	for _, docID := range topDocuments {
		imageData, found := appData.DocumentInfoMap[docID]
		if found {
			// Clean the URL
			cleanedURL := utils.CleanImageURL(imageData.URL, 400)

			// Check if the cleaned URL is already in the map of unique URLs
			_, exists := uniqueURLs[cleanedURL]
			if !exists {
				// If URL is not found, add the ImageData to the slice and mark the URL as encountered
				uniqueImageDatas = append(uniqueImageDatas, imageData)
				uniqueURLs[cleanedURL] = struct{}{}
			}
		}
	}

	responseJSON, err := json.Marshal(uniqueImageDatas)
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
