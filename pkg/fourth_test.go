package pkg

import (
	"io"
	"log"
	"reflect"
	"testing"
)

func TestParseCards(t *testing.T) {
	input := `
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

	expected := []map[string][]int{
		{"winning": {41, 48, 83, 86, 17}, "player": {83, 86, 6, 31, 17, 9, 48, 53}},
		{"winning": {13, 32, 20, 16, 61}, "player": {61, 30, 68, 82, 17, 32, 24, 19}},
		{"winning": {1, 21, 53, 59, 44}, "player": {69, 82, 63, 72, 16, 21, 14, 1}},
		{"winning": {41, 92, 73, 84, 69}, "player": {59, 84, 76, 51, 58, 5, 54, 83}},
		{"winning": {87, 83, 26, 28, 32}, "player": {88, 30, 70, 12, 93, 22, 82, 36}},
		{"winning": {31, 18, 13, 56, 72}, "player": {74, 77, 10, 23, 35, 67, 36, 11}},
	}
	currOutput := log.Writer()
	log.SetOutput(io.Discard)
	if result := parseCards(input); !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
	log.SetOutput(currOutput)
}

func TestCalculatePoints(t *testing.T) {
	cards := []map[string][]int{
		{"winning": {41, 48, 83, 86, 17}, "player": {83, 86, 6, 31, 17, 9, 48, 53}},
		{"winning": {13, 32, 20, 16, 61}, "player": {61, 30, 68, 82, 17, 32, 24, 19}},
		{"winning": {1, 21, 53, 59, 44}, "player": {69, 82, 63, 72, 16, 21, 14, 1}},
		{"winning": {41, 92, 73, 84, 69}, "player": {59, 84, 76, 51, 58, 5, 54, 83}},
		{"winning": {87, 83, 26, 28, 32}, "player": {88, 30, 70, 12, 93, 22, 82, 36}},
		{"winning": {31, 18, 13, 56, 72}, "player": {74, 77, 10, 23, 35, 67, 36, 11}},
	}

	currOutput := log.Writer()
	log.SetOutput(io.Discard)
	if result := calculatePoints(cards); result != 13 {
		t.Errorf("Expected 12 but got %d", result)
	}
	log.SetOutput(currOutput)
}

func TestWinningCards(t *testing.T) {
	cards := []map[string][]int{
		{"winning": {41, 48, 83, 86, 17}, "player": {83, 86, 6, 31, 17, 9, 48, 53}},
		{"winning": {13, 32, 20, 16, 61}, "player": {61, 30, 68, 82, 17, 32, 24, 19}},
		{"winning": {1, 21, 53, 59, 44}, "player": {69, 82, 63, 72, 16, 21, 14, 1}},
		{"winning": {41, 92, 73, 84, 69}, "player": {59, 84, 76, 51, 58, 5, 54, 83}},
		{"winning": {87, 83, 26, 28, 32}, "player": {88, 30, 70, 12, 93, 22, 82, 36}},
		{"winning": {31, 18, 13, 56, 72}, "player": {74, 77, 10, 23, 35, 67, 36, 11}},
	}

	currOutput := log.Writer()
	log.SetOutput(io.Discard)
	if result := winningCards(cards); result != 30 {
		t.Errorf("Expected 11 but got %d", result)
	}
	log.SetOutput(currOutput)
}

func BenchmarkParseCards(b *testing.B) {
	input := `
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

	currOutput := log.Writer()
	log.SetOutput(io.Discard)

	for i := 0; i < b.N; i++ {
		_ = parseCards(input)
	}

	// Restore output back
	log.SetOutput(currOutput)
}

func BenchmarkCalculatePoints(b *testing.B) {
	input := `
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

	currOutput := log.Writer()
	log.SetOutput(io.Discard)

	for i := 0; i < b.N; i++ {
		parsed := parseCards(input)
		calculatePoints(parsed)
	}

	// Restore output back
	log.SetOutput(currOutput)
}

func BenchmarkWinningCards(b *testing.B) {
	input := `
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

	currOutput := log.Writer()
	log.SetOutput(io.Discard)

	for i := 0; i < b.N; i++ {
		parsed := parseCards(input)
		winningCards(parsed)
	}

	// Restore output back
	log.SetOutput(currOutput)
}
