package utils

import "strings"

func ParseHashtags(content string) []string {
	var hashtags []string
	words := strings.Fields(content)
	for _, word := range words {
		if strings.HasPrefix(word, "#") && len(word) > 1 {
			hashtags = append(hashtags, strings.ToLower(word))
		}
	}
	return hashtags
}
