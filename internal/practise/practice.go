package practise

import (
	"math/rand"
	"time"

	"github.com/make17better/goju/pkg/goju"
)

// PracticeResult represents the result of a practice session
type PracticeResult struct {
	Date       time.Time
	Total      int
	Correct    int
	Incorrect  int
	Mistakes   []Mistake
	Categories []string
	Duration   time.Duration
}

// Mistake represents a mistake made during practice
type Mistake struct {
	Character goju.Character
	Input     string
	Attempts  int
	Correct   bool
	TimeSpent time.Duration
}

// PracticeSession represents an ongoing practice session
type PracticeSession struct {
	Count      int
	Categories []string
	Results    []PracticeResult
	StartTime  time.Time
	Current    struct {
		Character goju.Character
		Input     string
		Attempts  int
		StartTime time.Time
	}
}

// NewPracticeSession creates a new practice session
func NewPracticeSession(count int, categories []string) *PracticeSession {
	return &PracticeSession{
		Count:      count,
		Categories: categories,
		StartTime:  time.Now(),
	}
}

// GetNextCharacter returns a random character from the specified categories
func (p *PracticeSession) GetNextCharacter() goju.Character {
	var availableChars []goju.Character
	for _, category := range p.Categories {
		if chars, ok := goju.Characters[goju.Category(category)]; ok {
			availableChars = append(availableChars, chars...)
		}
	}

	if len(availableChars) == 0 {
		return goju.Character{}
	}

	rand.Seed(time.Now().UnixNano())
	return availableChars[rand.Intn(len(availableChars))]
}

// CheckAnswer checks if the provided answer is correct
func (p *PracticeSession) CheckAnswer(input string) bool {
	return input == p.Current.Character.Romaji
}

// RecordMistake records a mistake in the current practice session
func (p *PracticeSession) RecordMistake(input string) {
	p.Current.Attempts++
	p.Current.Input = input
}

// CompleteQuestion marks the current question as complete
func (p *PracticeSession) CompleteQuestion(correct bool) {
	duration := time.Since(p.Current.StartTime)
	mistake := Mistake{
		Character: p.Current.Character,
		Input:     p.Current.Input,
		Attempts:  p.Current.Attempts,
		Correct:   correct,
		TimeSpent: duration,
	}

	if len(p.Results) == 0 {
		p.Results = append(p.Results, PracticeResult{
			Date:       time.Now(),
			Categories: p.Categories,
		})
	}

	currentResult := &p.Results[len(p.Results)-1]
	currentResult.Total++
	if correct {
		currentResult.Correct++
	} else {
		currentResult.Incorrect++
		currentResult.Mistakes = append(currentResult.Mistakes, mistake)
	}
}

// GetAccuracy returns the accuracy percentage of the current session
func (p *PracticeSession) GetAccuracy() float64 {
	if len(p.Results) == 0 {
		return 0
	}
	currentResult := p.Results[len(p.Results)-1]
	if currentResult.Total == 0 {
		return 0
	}
	return float64(currentResult.Correct) / float64(currentResult.Total) * 100
}

// GetWeaknesses returns the most frequently missed characters
func (p *PracticeSession) GetWeaknesses(limit int) []Mistake {
	if len(p.Results) == 0 {
		return nil
	}

	mistakeCount := make(map[string]int)
	for _, result := range p.Results {
		for _, mistake := range result.Mistakes {
			mistakeCount[mistake.Character.Romaji]++
		}
	}

	var weaknesses []Mistake
	for _, result := range p.Results {
		for _, mistake := range result.Mistakes {
			if mistakeCount[mistake.Character.Romaji] >= 2 {
				weaknesses = append(weaknesses, mistake)
			}
		}
	}

	if len(weaknesses) > limit {
		return weaknesses[:limit]
	}
	return weaknesses
}
