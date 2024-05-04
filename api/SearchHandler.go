package api

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"search-server/types"
	"search-server/utils"
)

func SearchHandler(w http.ResponseWriter, r *http.Request, skipSpellCheck bool) {

	// Extract appData from context
	ctx := r.Context()
	appData, ok := ctx.Value(types.AppDataKey{}).(types.JsonData)

	if !ok {

		http.Error(w, "Failed to retrieve appData", http.StatusInternalServerError)
		return
	}

	// Extract query and spell check flag from URL parameters
	values := r.URL.Query()

	query := values.Get("q")

	correctedQuery, isCorrected := query, false

	if !skipSpellCheck {
		correctedQuery, isCorrected = appData.Trie.SpellCheck(query)
	}

	// Tokenize and preprocess the query
	queryTerms := utils.TokenizeText(correctedQuery)

	// Calculate TF-IDF for query terms
	queryVector := make(map[string]float64)
	for _, term := range queryTerms {
		queryVector[term] = calculateQueryTFIDF(term, appData, queryTerms)
	}

	// Rank documents based on cosine similarity
	// rankedDocuments := rankDocuments(queryVector, appData.DocumentVectors)
	rankedDocuments := utils.BM25Similarity(queryVector, 1.2, 0.75, appData)

	// Select top documents, considering additional factors such as document length
	topDocuments := selectTopDocuments(rankedDocuments, 30)

	// Retrieve and format top documents
	uniqueImageDatas := fetchAndFormatDocuments(appData, topDocuments)

	// Marshal JSON response
	// if `isCorrected` is true, return an extra field in the response for corrected query
	var response interface{}
	if isCorrected {
		response = struct {
			Query          string
			CorrectedQuery string
			Documents      []types.ImageData
		}{
			Query:          query,
			CorrectedQuery: correctedQuery,
			Documents:      uniqueImageDatas,
		}
	} else {
		response = struct {
			Query     string
			Documents []types.ImageData
		}{
			Query:     query,
			Documents: uniqueImageDatas,
		}
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println("Error marshalling JSON response:", err)
		http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
		return
	}

	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
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

func selectTopDocuments(rankedDocuments []string, count int) []string {
	if len(rankedDocuments) <= count {
		return rankedDocuments
	}
	return rankedDocuments[:count]
}

// func rankDocuments(queryVector map[string]float64, documentVectors map[string]types.DocumentVector) []string {
// 	// Calculate cosine similarity between query and each document
// 	similarityScores := make(map[string]float64)
// 	for docID, docVector := range documentVectors {
// 		similarityScores[docID] = calculateCosineSimilarity(queryVector, docVector)
// 	}

// 	// Rank documents based on similarity scores
// 	return rankByScore(similarityScores)
// }

// func calculateCosineSimilarity(vector1, vector2 map[string]float64) float64 {
// 	// Compute dot product
// 	dotProduct := 0.0
// 	for term, tfidf1 := range vector1 {
// 		tfidf2, exists := vector2[term]
// 		if exists {
// 			dotProduct += tfidf1 * tfidf2
// 		}
// 	}

// 	// Compute magnitudes
// 	magnitude1 := calculateMagnitude(vector1)
// 	magnitude2 := calculateMagnitude(vector2)

// 	// Compute cosine similarity
// 	if magnitude1 != 0 && magnitude2 != 0 {
// 		return dotProduct / (magnitude1 * magnitude2)
// 	}
// 	return 0
// }

// func calculateMagnitude(vector map[string]float64) float64 {
// 	sum := 0.0
// 	for _, value := range vector {
// 		sum += value * value
// 	}
// 	return math.Sqrt(sum)
// }

// func rankByScore(scores map[string]float64) []string {
// 	// Convert map to slice of {docID, score} pairs and sort by score
// 	type pair struct {
// 		docID string
// 		score float64
// 	}
// 	pairs := make([]pair, len(scores))
// 	i := 0
// 	for docID, score := range scores {
// 		pairs[i] = pair{docID, score}
// 		i++
// 	}
// 	sort.Slice(pairs, func(i, j int) bool {
// 		return pairs[i].score > pairs[j].score
// 	})

// 	// Extract sorted document IDs
// 	result := make([]string, len(pairs))
// 	for i, pair := range pairs {
// 		result[i] = pair.docID
// 	}
// 	return result
// }
