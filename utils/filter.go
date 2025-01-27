package utils

import (
	"strings"

	snowballeng "github.com/kljensen/snowball/english"
)

func lowercaseFiler(tokens []string) []string{
	r := make([]string, len(tokens))
	for i, token := range tokens{
		r[i] = strings.ToLower(token)
	}
	return r
}

func stopwordFilers (tokens []string) []string{
	r := make([]string, len(tokens))
	var stopwords = map[string]struct{}{
		"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
		"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
	}
	for _, token := range tokens{
		if _, exists := stopwords[token]; exists {
			continue
		}
		r = append(r, token)
	}
	return r
}

func stemmerFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = snowballeng.Stem(token, false)
	}
	return r
}