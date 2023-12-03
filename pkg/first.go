package pkg

import (
	"log"
	"strings"
	"unicode"
)

func Calibration() {
	inputString := `
	1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`

	log.Println("-------Advent of Code 2023 #1-------")
	lines := strings.Split(inputString, "\n")
	for _, line := range lines {
		log.Println(strings.TrimSpace(line))
	}
	log.Println("---------")
	log.Println()

	pairs, err := findNumberPairs(inputString)
	if err != nil {
		log.Println(err.Error())
		return
	}

	sum := 0
	for _, num := range pairs {
		sum += num
	}

	log.Println("Result: ", sum)
}

func findNumberPairs(inputString string) ([]int, error) {
	var numberPairs []int
	var firstDigit, lastDigit int
	foundFirst, foundLast := false, false

	lines := strings.Split(inputString, "\n")

	for _, line := range lines {
		for _, r := range line {
			if unicode.IsDigit(r) {
				firstDigit = int(r - '0')
				foundFirst = true
				log.Println("Found first digit: ", firstDigit)
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				lastDigit = int(line[i] - '0')
				foundLast = true
				log.Println("Found second digit: ", lastDigit)
				break
			}
		}

		if foundFirst && foundLast {
			numberPairs = append(numberPairs, firstDigit*10+lastDigit)
			log.Println("Found pair: ", firstDigit*10+lastDigit)
			log.Println("---------")
		}
	}

	return numberPairs, nil
}
