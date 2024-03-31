package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func CleanImageURL(url string, width int) string {
	// Implement the logic of cleanImageUrl function in Go
	// Here's an example of the equivalent Go code:
	cleanedURL := regexp.MustCompile(`(\?|\&)(q|auto|fit|ixlib|ixid)=[^&]+`).ReplaceAllString(url, "")
	replacedURL := regexp.MustCompile(`(\?|\&)w=[^&]+`).ReplaceAllString(cleanedURL, "")
	modifiedURL := fmt.Sprintf("%s%sw=%d", replacedURL, func() string {
		if strings.Contains(replacedURL, "?") {
			return "&"
		}
		return "?"
	}(), width)
	return modifiedURL
}
