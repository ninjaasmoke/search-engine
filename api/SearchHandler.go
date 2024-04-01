package api

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"search-server/types"
	"search-server/utils"
	"sort"
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

// func SearchHandler(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	appData, ok := ctx.Value(types.AppDataKey{}).(types.JsonData)

// 	if !ok {
// 		http.Error(w, "Failed to retrieve appData", http.StatusInternalServerError)
// 		return
// 	}

// 	// Tokenize and preprocess the query
// 	query := r.URL.Query().Get("q")
// 	queryTerms := utils.TokenizeText(query)

// 	// Calculate TF-IDF for query terms
// 	queryVector := make(map[string]float64)
// 	for _, term := range queryTerms {
// 		queryVector[term] = math.Log(float64(appData.TotalDocs) / float64(appData.DocumentFrequency[term]+1))
// 	}

// 	// Rank documents based on cosine similarity
// 	rankedDocuments := utils.RankDocuments(queryVector, appData.DocumentVectors)

// 	topDocuments := rankedDocuments[:40]

// 	// Create a map to store unique URLs
// 	uniqueURLs := make(map[string]struct{})

// 	// Slice to store unique ImageData entries
// 	uniqueImageDatas := make([]types.ImageData, 0, len(topDocuments))

// 	// Iterate over topDocuments to filter out duplicates based on URL
// 	for _, docID := range topDocuments {
// 		imageData, found := appData.DocumentInfoMap[docID]
// 		imageData.ID = docID
// 		if found {
// 			// Clean the URL
// 			cleanedURL := utils.CleanImageURL(imageData.URL, 400)

// 			// Check if the cleaned URL is already in the map of unique URLs
// 			_, exists := uniqueURLs[cleanedURL]
// 			if !exists {
// 				// If URL is not found, add the ImageData to the slice and mark the URL as encountered
// 				uniqueImageDatas = append(uniqueImageDatas, imageData)
// 				uniqueURLs[cleanedURL] = struct{}{}
// 			}
// 		}
// 	}

// 	responseJSON, err := json.Marshal(uniqueImageDatas[:33])
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

	// Extract appData from context
	ctx := r.Context()
	appData, ok := ctx.Value(types.AppDataKey{}).(types.JsonData)

	if !ok {

		http.Error(w, "Failed to retrieve appData", http.StatusInternalServerError)
		return
	}

	// Extract query from request
	query := r.URL.Query().Get("q")

	// Tokenize and preprocess the query
	queryTerms := utils.TokenizeText(query)

	// Calculate TF-IDF for query terms
	queryVector := make(map[string]float64)
	for _, term := range queryTerms {
		queryVector[term] = calculateQueryTFIDF(term, appData, queryTerms)
	}

	// Rank documents based on cosine similarity
	rankedDocuments := rankDocuments(queryVector, appData.DocumentVectors)

	// Select top documents, considering additional factors such as document length
	topDocuments := selectTopDocuments(rankedDocuments, 60)

	// Retrieve and format top documents
	uniqueImageDatas := fetchAndFormatDocuments(appData, topDocuments)

	// Marshal JSON response
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

func fetchAndFormatDocuments(appData types.JsonData, documentIDs []string) []types.ImageData {
	var imageDatas []types.ImageData
	uniqueURLs := make(map[string]struct{})

	for _, docID := range documentIDs {
		imageData, found := appData.DocumentInfoMap[docID]
		if found {
			// Clean the URL
			cleanedURL := utils.CleanImageURL(imageData.URL, 400)

			// Check if the cleaned URL is already in the map of unique URLs
			_, exists := uniqueURLs[cleanedURL]
			if !exists {
				// If URL is not found, add the ImageData to the slice and mark the URL as encountered
				imageData.ID = docID
				imageDatas = append(imageDatas, imageData)
				uniqueURLs[cleanedURL] = struct{}{}
			}
		}
	}

	return imageDatas
}

func calculateQueryTFIDF(term string, appData types.JsonData, queryTerms []string) float64 {
	// Calculate TF for the term in the query
	tf := calculateTermFrequency(term, queryTerms)

	// Calculate IDF for the term
	idf := calculateInverseDocumentFrequency(term, appData)

	// Calculate TF-IDF
	return tf * idf
}

func calculateTermFrequency(term string, queryTerms []string) float64 {
	// Count the frequency of the term in the query
	count := 0
	for _, t := range queryTerms {
		if t == term {
			count++
		}
	}

	// Calculate term frequency
	return float64(count) / float64(len(queryTerms))
}

func calculateInverseDocumentFrequency(term string, appData types.JsonData) float64 {
	// Calculate IDF for the term
	df := float64(appData.DocumentFrequency[term])
	totalDocs := float64(appData.TotalDocs)
	return math.Log(totalDocs / (df + 1))
}

func rankDocuments(queryVector map[string]float64, documentVectors map[string]types.DocumentVector) []string {
	// Calculate cosine similarity between query and each document
	similarityScores := make(map[string]float64)
	for docID, docVector := range documentVectors {
		similarityScores[docID] = calculateCosineSimilarity(queryVector, docVector)
	}

	// Rank documents based on similarity scores
	return rankByScore(similarityScores)
}

func calculateCosineSimilarity(vector1, vector2 map[string]float64) float64 {
	// Compute dot product
	dotProduct := 0.0
	for term, tfidf1 := range vector1 {
		tfidf2, exists := vector2[term]
		if exists {
			dotProduct += tfidf1 * tfidf2
		}
	}

	// Compute magnitudes
	magnitude1 := calculateMagnitude(vector1)
	magnitude2 := calculateMagnitude(vector2)

	// Compute cosine similarity
	if magnitude1 != 0 && magnitude2 != 0 {
		return dotProduct / (magnitude1 * magnitude2)
	}
	return 0
}

func calculateMagnitude(vector map[string]float64) float64 {
	sum := 0.0
	for _, value := range vector {
		sum += value * value
	}
	return math.Sqrt(sum)
}

func selectTopDocuments(rankedDocuments []string, count int) []string {
	if len(rankedDocuments) <= count {
		return rankedDocuments
	}
	return rankedDocuments[:count]
}

func rankByScore(scores map[string]float64) []string {
	// Convert map to slice of {docID, score} pairs and sort by score
	type pair struct {
		docID string
		score float64
	}
	pairs := make([]pair, len(scores))
	i := 0
	for docID, score := range scores {
		pairs[i] = pair{docID, score}
		i++
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].score > pairs[j].score
	})

	// Extract sorted document IDs
	result := make([]string, len(pairs))
	for i, pair := range pairs {
		result[i] = pair.docID
	}
	return result
}
