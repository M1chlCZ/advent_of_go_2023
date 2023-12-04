package pkg

import (
	"log"
	"strconv"
	"strings"
)

func WinningCards() {
	input := `
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

	cardsData := parseCards(input)

	log.Println("-------Advent of Code 2023 #4-------")
	result := calculatePoints(cardsData)
	log.Println("The sum of winning cards points")
	log.Println("Result first part: ", result)
	result = winningCards(cardsData)
	log.Println("Result second part: ", result)
	log.Println("---------")
	log.Println()

}

func parseCards(input string) []map[string][]int {
	lines := strings.Split(input, "\n")
	cardsData := make([]map[string][]int, 0)

	for i, line := range lines {
		if len(line) == 0 && i == 0 {
			continue
		}
		parts := strings.Split(line, "|")
		winningNumbers := strings.Fields(strings.Split(parts[0], ":")[1])
		playerNumbers := strings.Fields(parts[1])

		winning := make([]int, len(winningNumbers))
		player := make([]int, len(playerNumbers))

		for i, w := range winningNumbers {
			winning[i], _ = strconv.Atoi(w)
		}

		for i, p := range playerNumbers {
			player[i], _ = strconv.Atoi(p)
		}

		card := map[string][]int{
			"winning": winning,
			"player":  player,
		}

		cardsData = append(cardsData, card)
	}
	return cardsData
}

func calculatePoints(cards []map[string][]int) int {
	totalPoints := 0
	for _, card := range cards {
		winning := card["winning"]
		player := card["player"]
		matches := map[int]bool{}
		for _, winningNumber := range winning {
			for _, playerNumber := range player {
				if winningNumber == playerNumber {
					matches[winningNumber] = true
				}
			}
		}
		if len(matches) > 0 {
			cardPoints := 1
			for i := 0; i < len(matches)-1; i++ {
				cardPoints *= 2
			}
			totalPoints += cardPoints
		}
	}
	return totalPoints
}

func winningCards(cardsData []map[string][]int) int {
	cardCopies := make([]int, len(cardsData))
	for i := range cardCopies {
		cardCopies[i] = 1 // Each card starts with 1 copy (the original)
	}

	totalCards := len(cardsData) // Initial count of cards

	for i, card := range cardsData {
		matches := countMatches(card["winning"], card["player"])

		// Distribute copies to subsequent cards
		for j := 1; j <= matches && i+j < len(cardsData); j++ {
			cardCopies[i+j] += cardCopies[i]
			totalCards += cardCopies[i] // Add the copies to the total count
		}
	}
	return totalCards
}

func countMatches(winning, player []int) int {
	matchCount := 0
	for _, w := range winning {
		for _, p := range player {
			if w == p {
				matchCount++
				break
			}
		}
	}
	return matchCount
}
