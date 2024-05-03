package types

type ImageData struct {
	ID                 string   `json:"id"`
	URL                string   `json:"url"`
	Title              string   `json:"title"`
	RelatedImageTags   []string `json:"related_image_tags"`
	AnnotatedImageTags []string `json:"annotated_image_tags"`
}

type InvertedIndex struct {
	TF map[string]int `json:"tf"`
	DF int            `json:"df"`
}

type AppDataKey struct{}

type DocumentVector map[string]float64

type JsonData struct {
	DocumentInfoMap   map[string]ImageData
	InvertedIndexMap  map[string]InvertedIndex
	TotalDocs         int
	DocumentFrequency map[string]int
	DocumentVectors   map[string]DocumentVector
	AveraageDocLength float64
}
