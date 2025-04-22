package learn

import (
	"fmt"
	"strings"

	"github.com/make17better/goju/pkg/goju"
)

// Difficulty represents the learning difficulty level
type Difficulty string

const (
	Easy   Difficulty = "easy"   // Seion only
	Normal Difficulty = "normal" // Seion + Dakuon
	Hard   Difficulty = "hard"   // All categories
)

// LearningContent represents the content for learning mode
type LearningContent struct {
	Title      string
	Categories []string
	Characters []goju.Character
}

// GetLearningContent returns learning content based on difficulty and script type
func GetLearningContent(difficulty Difficulty, scriptType string) LearningContent {
	var content LearningContent
	var categories []string

	// Set categories based on difficulty
	switch difficulty {
	case Easy:
		categories = []string{string(goju.Seion)}
		content.Title = "Basic Sounds (清音)"
	case Normal:
		categories = []string{string(goju.Seion), string(goju.Dakuon)}
		content.Title = "Basic and Voiced Sounds (清音・浊音)"
	case Hard:
		categories = []string{string(goju.Seion), string(goju.Dakuon), string(goju.Handaku), string(goju.Yoon)}
		content.Title = "All Sounds (五十音)"
	}

	content.Categories = categories

	// Collect characters based on categories and script type
	for _, category := range categories {
		if chars, ok := goju.Characters[goju.Category(category)]; ok {
			content.Characters = append(content.Characters, chars...)
		}
	}

	return content
}

// FormatCharacter formats a character for display
func FormatCharacter(char goju.Character, scriptType string) string {
	var display string
	switch strings.ToLower(scriptType) {
	case "hiragana":
		display = char.Hiragana
	case "katakana":
		display = char.Katakana
	case "both":
		display = fmt.Sprintf("%s (%s)", char.Hiragana, char.Katakana)
	default:
		display = char.Hiragana
	}

	return fmt.Sprintf("%s - %s", display, char.Romaji)
}

// FormatCategory formats a category for display
func FormatCategory(category string) string {
	switch category {
	case string(goju.Seion):
		return "清音 (Basic Sounds)"
	case string(goju.Dakuon):
		return "浊音 (Voiced Sounds)"
	case string(goju.Handaku):
		return "半浊音 (Half-Voiced Sounds)"
	case string(goju.Yoon):
		return "拗音 (Contracted Sounds)"
	default:
		return category
	}
}

// FormatLearningContent formats the learning content for display
func FormatLearningContent(content LearningContent, scriptType string) string {
	var sb strings.Builder

	// Add title
	sb.WriteString(fmt.Sprintf("%s\n\n", content.Title))

	// Group characters by category
	categoryChars := make(map[string][]goju.Character)
	for _, char := range content.Characters {
		categoryChars[string(char.Category)] = append(categoryChars[string(char.Category)], char)
	}

	// Format each category
	for _, category := range content.Categories {
		sb.WriteString(fmt.Sprintf("%s:\n", FormatCategory(category)))
		if chars, ok := categoryChars[category]; ok {
			for _, char := range chars {
				sb.WriteString(fmt.Sprintf("  %s\n", FormatCharacter(char, scriptType)))
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
