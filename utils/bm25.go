package utils

import (
	"math"
	"search-server/types"
	"sort"
)

func GetAverageDocumentLength(docInfoMap map[string]types.ImageData) float64 {
	totalLength := 0.0
	for _, docInfo := range docInfoMap {
		totalLength += float64(len(docInfo.Title))
		for _, tag := range docInfo.RelatedImageTags {
			totalLength += float64(len(tag))
		}
		for _, tag := range docInfo.AnnotatedImageTags {
			totalLength += float64(len(tag))
		}
	}
	return totalLength / float64(len(docInfoMap))
}

func BM25Similarity(queryVector map[string]float64, k1 float64, b float64, appData types.JsonData) []string {
	scores := make(map[string]float64)

	averageDocLength := appData.AveraageDocLength

	for term, queryWeight := range queryVector {
		if invertedIndex, found := appData.InvertedIndexMap[term]; found {
			df := float64(invertedIndex.DF)
			idf := math.Log(float64(appData.TotalDocs) / df)

			for docID, tf := range invertedIndex.TF {
				docLength := float64(len(appData.DocumentInfoMap[docID].Title))
				docLength += float64(len(appData.DocumentInfoMap[docID].RelatedImageTags))
				docLength += float64(len(appData.DocumentInfoMap[docID].AnnotatedImageTags))

				// Calculate BM25 score
				score := idf * ((float64(tf) * (k1 + 1)) / (float64(tf) + k1*(1-b+b*(docLength/averageDocLength))) * queryWeight)
				if _, found := scores[docID]; !found {
					scores[docID] = 0
				}
				scores[docID] += score
			}
		}
	}

	// Filter out documents with zero similarity scores
	rankedDocuments := make([]string, 0)

	for docID, score := range scores {
		if score > 0 {
			rankedDocuments = append(rankedDocuments, docID)
		}
	}

	// If there are no documents, return empty result
	if len(rankedDocuments) == 0 {
		return []string{}
	}

	// Sort document IDs by descending scores
	sort.Slice(rankedDocuments, func(i, j int) bool {
		return scores[rankedDocuments[i]] > scores[rankedDocuments[j]]
	})

	return rankedDocuments
}
