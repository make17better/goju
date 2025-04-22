package goju

import (
	"testing"
)

func TestGetCharacterByHiragana(t *testing.T) {
	tests := []struct {
		name      string
		hiragana  string
		wantFound bool
	}{
		{"Valid hiragana", "あ", true},
		{"Valid hiragana", "が", true},
		{"Valid hiragana", "きゃ", true},
		{"Invalid hiragana", "ああ", false},
		{"Empty string", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, found := GetCharacterByHiragana(tt.hiragana)
			if found != tt.wantFound {
				t.Errorf("GetCharacterByHiragana() found = %v, want %v", found, tt.wantFound)
			}
		})
	}
}

func TestGetCharacterByKatakana(t *testing.T) {
	tests := []struct {
		name      string
		katakana  string
		wantFound bool
	}{
		{"Valid katakana", "ア", true},
		{"Valid katakana", "ガ", true},
		{"Valid katakana", "キャ", true},
		{"Invalid katakana", "アア", false},
		{"Empty string", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, found := GetCharacterByKatakana(tt.katakana)
			if found != tt.wantFound {
				t.Errorf("GetCharacterByKatakana() found = %v, want %v", found, tt.wantFound)
			}
		})
	}
}

func TestGetCharacterByRomaji(t *testing.T) {
	tests := []struct {
		name      string
		romaji    string
		wantFound bool
	}{
		{"Valid romaji", "a", true},
		{"Valid romaji", "ga", true},
		{"Valid romaji", "kya", true},
		{"Invalid romaji", "aa", false},
		{"Empty string", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, found := GetCharacterByRomaji(tt.romaji)
			if found != tt.wantFound {
				t.Errorf("GetCharacterByRomaji() found = %v, want %v", found, tt.wantFound)
			}
		})
	}
}
