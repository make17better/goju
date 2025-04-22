package lookup

import (
	"testing"

	"github.com/make17better/goju/pkg/goju"
)

func TestLookup(t *testing.T) {
	tests := []struct {
		name      string
		inputType string
		value     string
		wantFound bool
	}{
		{"Valid hiragana", "hiragana", "あ", true},
		{"Valid katakana", "katakana", "ア", true},
		{"Valid romaji", "romaji", "a", true},
		{"Invalid type", "invalid", "あ", false},
		{"Invalid value", "hiragana", "ああ", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Lookup(tt.inputType, tt.value)
			if result.Found != tt.wantFound {
				t.Errorf("Lookup() found = %v, want %v", result.Found, tt.wantFound)
			}
		})
	}
}

func TestBatchLookup(t *testing.T) {
	tests := []struct {
		name      string
		inputType string
		values    []string
		wantCount int
	}{
		{"Valid hiragana batch", "hiragana", []string{"あ", "い", "う"}, 3},
		{"Valid katakana batch", "katakana", []string{"ア", "イ", "ウ"}, 3},
		{"Valid romaji batch", "romaji", []string{"a", "i", "u"}, 3},
		{"Mixed valid/invalid", "hiragana", []string{"あ", "ああ", "い"}, 3},
		{"Empty batch", "hiragana", []string{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := BatchLookup(tt.inputType, tt.values)
			if len(results) != tt.wantCount {
				t.Errorf("BatchLookup() count = %v, want %v", len(results), tt.wantCount)
			}
		})
	}
}

func TestFormatLookupResult(t *testing.T) {
	tests := []struct {
		name   string
		result LookupResult
		want   string
	}{
		{
			"Valid character",
			LookupResult{
				Character: goju.Character{
					Hiragana: "あ",
					Katakana: "ア",
					Romaji:   "a",
					Category: goju.Seion,
				},
				Found: true,
			},
			"Hiragana: あ\nKatakana: ア\nRomaji: a\nCategory: seion",
		},
		{
			"Not found",
			LookupResult{Found: false},
			"Character not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatLookupResult(tt.result)
			if got != tt.want {
				t.Errorf("FormatLookupResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

// CharacterArrayLookup looks up multiple characters at once
func CharacterArrayLookup(input []string) ([]*goju.Character, error) {
	results := make([]*goju.Character, len(input))
	var err error

	for i, char := range input {
		result := Lookup(char, "hiragana") // Default to hiragana lookup
		if result.Found {
			results[i] = &result.Character
		}
	}

	return results, err
}

func TestCharacterArrayLookup(t *testing.T) {
	tests := []struct {
		name        string
		input       []string
		expected    []*goju.Character
		expectError bool
	}{
		{
			name:  "Valid character array",
			input: []string{"あ", "い", "う"},
			expected: []*goju.Character{
				{Hiragana: "あ", Romaji: "a", Category: goju.Seion},
				{Hiragana: "い", Romaji: "i", Category: goju.Seion},
				{Hiragana: "う", Romaji: "u", Category: goju.Seion},
			},
			expectError: false,
		},
		{
			name:  "Invalid character array",
			input: []string{"あ", "invalid", "う"},
			expected: []*goju.Character{
				{Hiragana: "あ", Romaji: "a", Category: goju.Seion},
				nil,
				{Hiragana: "う", Romaji: "u", Category: goju.Seion},
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results, err := CharacterArrayLookup(tt.input)
			if err != nil && !tt.expectError {
				t.Errorf("CharacterArrayLookup() error = %v, expectError = %v", err, tt.expectError)
			}
			if len(results) != len(tt.expected) {
				t.Errorf("CharacterArrayLookup() results count = %v, want %v", len(results), len(tt.expected))
			}
			for i, result := range results {
				if result != nil && tt.expected[i] != nil {
					if result.Hiragana != tt.expected[i].Hiragana || result.Romaji != tt.expected[i].Romaji || result.Category != tt.expected[i].Category {
						t.Errorf("CharacterArrayLookup() result[%d] = %v, want %v", i, result, tt.expected[i])
					}
				}
			}
		})
	}
}
