package lookup

import (
	"fmt"
	"strings"

	"github.com/make17better/goju/pkg/goju"
)

// LookupResult represents the result of a character lookup
type LookupResult struct {
	Character goju.Character
	Found     bool
}

// Lookup performs a character lookup based on the input type and value
func Lookup(inputType, value string) LookupResult {
	var result LookupResult
	var found bool

	switch strings.ToLower(inputType) {
	case "hiragana":
		result.Character, found = goju.GetCharacterByHiragana(value)
	case "katakana":
		result.Character, found = goju.GetCharacterByKatakana(value)
	case "romaji":
		result.Character, found = goju.GetCharacterByRomaji(value)
	default:
		return LookupResult{Found: false}
	}

	result.Found = found
	return result
}

// FormatLookupResult formats the lookup result for display
func FormatLookupResult(result LookupResult) string {
	if !result.Found {
		return "Character not found"
	}

	return fmt.Sprintf(
		"Hiragana: %s\nKatakana: %s\nRomaji: %s\nCategory: %s",
		result.Character.Hiragana,
		result.Character.Katakana,
		result.Character.Romaji,
		result.Character.Category,
	)
}

// BatchLookup performs multiple character lookups
func BatchLookup(inputType string, values []string) []LookupResult {
	results := make([]LookupResult, len(values))
	for i, value := range values {
		results[i] = Lookup(inputType, value)
	}
	return results
}

// FormatBatchLookup formats multiple lookup results for display
func FormatBatchLookup(results []LookupResult) string {
	var sb strings.Builder
	for i, result := range results {
		if i > 0 {
			sb.WriteString("\n\n")
		}
		sb.WriteString(FormatLookupResult(result))
	}
	return sb.String()
}
