package utils

import (
	"math"
	"search-server/types"
	"sort"
)

func CosineSimilarity(a, b map[string]float64) float64 {
	dotProduct := 0.0
	magnitudeA := 0.0
	magnitudeB := 0.0

	// Calculate dot product and magnitudes
	for term, weightA := range a {
		dotProduct += weightA * b[term]
		magnitudeA += math.Pow(weightA, 2)
	}
	for _, weightB := range b {
		magnitudeB += math.Pow(weightB, 2)
	}

	// Handle division by zero
	if magnitudeA == 0 || magnitudeB == 0 {
		return 0.0
	}

	// Calculate cosine similarity
	return dotProduct / (math.Sqrt(magnitudeA) * math.Sqrt(magnitudeB))
}

func RankDocuments(queryVector map[string]float64, documentVectors map[string]types.DocumentVector) []string {
	scores := make(map[string]float64)
	for docID, docVector := range documentVectors {
		scores[docID] = CosineSimilarity(queryVector, docVector)
	}

	// Sort document IDs by descending scores
	rankedDocuments := make([]string, 0, len(scores))
	for docID := range scores {
		rankedDocuments = append(rankedDocuments, docID)
	}
	sort.Slice(rankedDocuments, func(i, j int) bool {
		return scores[rankedDocuments[i]] > scores[rankedDocuments[j]]
	})

	return rankedDocuments
}

func GenerateDocumentVectors(appData types.JsonData) map[string]types.DocumentVector {
	documentVectors := make(map[string]types.DocumentVector)

	idfs := make(map[string]float64)
	for term, posting := range appData.InvertedIndexMap {
		idfs[term] = math.Log(float64(appData.TotalDocs) / float64(posting.DF+1))
	}

	for term, posting := range appData.InvertedIndexMap {
		idf := idfs[term]

		for docID, tf := range posting.TF {
			tfIDF := float64(tf) * idf

			if _, found := documentVectors[docID]; !found {
				documentVectors[docID] = make(types.DocumentVector)
			}

			documentVectors[docID][term] = tfIDF
		}
	}

	return documentVectors
}
