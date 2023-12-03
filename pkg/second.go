package pkg

import (
	"log"
	"strconv"
	"strings"
)

var inputData = `
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`

func CubeGame() {
	log.Println("-------Advent of Code 2023 #2-------")
	result := processGames(inputData)
	log.Println("The sum of the IDs of the possible games")
	log.Println("Result: ", result)
	log.Println("---------")
	log.Println()
}

func isGamePossible(sets []map[string]int, maxCubes map[string]int) bool {
	for _, set := range sets {
		for color, count := range set {
			if count > maxCubes[color] {
				return false
			}
		}
	}
	return true
}

func processGames(inputData string) int {
	var possibleGames []int
	maxCubes := map[string]int{"red": 12, "green": 13, "blue": 14}

	for _, game := range strings.Split(inputData, "\n") {
		game = strings.TrimSpace(game)
		if len(game) == 0 {
			continue
		}
		gameID, setsData := strings.Split(game, ":")[0], strings.Split(game, ":")[1]
		gameID = strings.TrimSpace(strings.Replace(gameID, "Game", "", -1))
		gameID = strings.TrimSpace(gameID)
		setsData = strings.TrimSpace(setsData)
		var sets []map[string]int
		for _, set := range strings.Split(setsData, ";") {
			set = strings.TrimSpace(set)
			setData := map[string]int{}
			for _, item := range strings.Split(set, ",") {
				item = strings.TrimSpace(item)
				count, color := strings.Split(item, " ")[0], strings.Split(item, " ")[1]
				countColor, err := strconv.Atoi(count)
				if err != nil {
					log.Println(err.Error())
					continue
				}
				setData[color] = countColor
			}
			sets = append(sets, setData)
		}

		if isGamePossible(sets, maxCubes) {
			possibleGame, err := strconv.Atoi(gameID)
			if err != nil {
				log.Println(err.Error())
				continue
			}
			possibleGames = append(possibleGames, possibleGame)
		}
	}

	sum := 0
	for _, game := range possibleGames {
		sum += game
	}

	return sum
}
