package pkg

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

func isSymbol(r rune) bool {
	return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '.'
}

func sumPartNumbers(schematic []string) int {
	sum := 0
	for y, line := range schematic {
		for x := 0; x < len(line); x++ {
			char := rune(line[x])
			if unicode.IsDigit(char) {
				number, end := extractNumber(line, x)
				if isAdjacentToSymbol(schematic, x, y, end) {
					log.Println("Found number at: ", number, "at", x, y)
					//log.Println("Found number at: ", number, "at", x, y)
					value, _ := strconv.Atoi(number)
					sum += value
					x = end - 1 // Skip to the end of the current number
				}
			}
		}
	}
	return sum
}

func extractNumber(line string, start int) (string, int) {
	end := start
	for end < len(line) && unicode.IsDigit(rune(line[end])) {
		end++
	}
	return line[start:end], end
}

func isAdjacentToSymbol(schematic []string, x, y, end int) bool {
	for xi := x; xi < end; xi++ {
		if isSymbolInAdjacentCells(schematic, xi, y) {
			return true
		}
	}
	return false
}

func isSymbolInAdjacentCells(schematic []string, x, y int) bool {
	directions := []struct{ dx, dy int }{
		{-1, -1}, {0, -1}, {1, -1}, // top row
		{-1, 0}, {1, 0}, // middle row
		{-1, 1}, {0, 1}, {1, 1}, // bottom row
	}

	for _, dir := range directions {
		nx, ny := x+dir.dx, y+dir.dy
		if nx >= 0 && ny >= 0 && ny < len(schematic) && nx < len(schematic[ny]) {
			if isSymbol(rune(schematic[ny][nx])) {
				return true
			}
		}
	}
	return false
}

func Schematic() {
	//BEWARE OF THE WHITE SPACE, IT WILL BREAK YOUR MOTIVATION TO LIVE
	inputString := `467..114..
					...*......
					..35..633.
					......#...
					617*......
					.....+.58.
					..592.....
					......755.
					...$.*....
					.664.598..`

	var schematic []string
	for _, line := range strings.Split(inputString, "\n") {
		schematic = append(schematic, strings.TrimSpace(line))
	}
	log.Println("-------Advent of Code 2023 #3-------")
	sum := sumPartNumbers(schematic)
	log.Println("The sum of all part numbers")
	log.Println("Result: ", sum)
	log.Println()
}
