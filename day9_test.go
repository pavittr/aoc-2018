package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func nextTurn(gameBoard []int, currentMarbleIndex int, newMarble int) ([]int, int, int) {
	playersScore := 0
	if len(gameBoard) == 1 {
		newGameBoard := append(gameBoard, 1)
		return newGameBoard, 1, 0
	}

	if newMarble%23 == 0 {
		playersScore += newMarble
		breakPoint := currentMarbleIndex - 7
		if breakPoint < 0 {
			breakPoint += len(gameBoard) - 1
		}
		playersScore += gameBoard[breakPoint]

		newGameBoard := make([]int, 0)
		newGameBoard = append(newGameBoard, gameBoard[0:breakPoint]...)
		newGameBoard = append(newGameBoard, gameBoard[breakPoint+1:]...)
		return newGameBoard, breakPoint, playersScore
	}
	breakPoint := currentMarbleIndex + 2

	if breakPoint > len(gameBoard) {
		breakPoint = 1
	}
	newGameBoard := make([]int, 0)
	newGameBoard = append(newGameBoard, gameBoard[0:breakPoint]...)
	newGameBoard = append(newGameBoard, newMarble)
	newGameBoard = append(newGameBoard, gameBoard[breakPoint:]...)
	return newGameBoard, breakPoint, 0
}

func TestTurn(t *testing.T) {
	givenBoard := []int{0}
	andCurrentIndex := 0
	whenGameBoard, whenMarbleIndex, roundScore := nextTurn(givenBoard, andCurrentIndex, 1)
	assert.Equal(t, []int{0, 1}, whenGameBoard)
	assert.Equal(t, 1, whenMarbleIndex)
	assert.Equal(t, 0, roundScore)

	givenBoard = []int{0, 8, 4, 2, 5, 1, 6, 3, 7}
	andCurrentIndex = 1
	whenGameBoard, whenMarbleIndex, roundScore = nextTurn(givenBoard, andCurrentIndex, 9)
	assert.Equal(t, []int{0, 8, 4, 9, 2, 5, 1, 6, 3, 7}, whenGameBoard)
	assert.Equal(t, 3, whenMarbleIndex)
	assert.Equal(t, 0, roundScore)

	// [6]  0  4  2  5  1 (6) 3
	// [7]  0  4  2  5  1  6  3 (7)
	givenBoard = []int{0, 4, 2, 5, 1, 6, 3}
	andCurrentIndex = 5
	whenGameBoard, whenMarbleIndex, roundScore = nextTurn(givenBoard, andCurrentIndex, 7)
	assert.Equal(t, []int{0, 4, 2, 5, 1, 6, 3, 7}, whenGameBoard)
	assert.Equal(t, 7, whenMarbleIndex)
	assert.Equal(t, 0, roundScore)

	// [7]  0  4  2  5  1  6  3 (7)
	// [8]  0 (8) 4  2  5  1  6  3  7
	givenBoard = []int{0, 4, 2, 5, 1, 6, 3, 7}
	andCurrentIndex = 7
	whenGameBoard, whenMarbleIndex, roundScore = nextTurn(givenBoard, andCurrentIndex, 8)
	assert.Equal(t, []int{0, 8, 4, 2, 5, 1, 6, 3, 7}, whenGameBoard)
	assert.Equal(t, 1, whenMarbleIndex)
	assert.Equal(t, 0, roundScore)

	//[22]  0 16  8 17  4 18  9 19  2 20 10 21  5(22)11  1 12  6 13  3 14  7 15
	//[23]  0 16  8 17  4 18(19) 2 20 10 21  5 22 11  1 12  6 13  3 14  7 15
	givenBoard = []int{0, 16, 8, 17, 4, 18, 9, 19, 2, 20, 10, 21, 5, 22, 11, 1, 12, 6, 13, 3, 14, 7, 15}
	andCurrentIndex = 13
	whenGameBoard, whenMarbleIndex, roundScore = nextTurn(givenBoard, andCurrentIndex, 23)

	assert.Equal(t, []int{0, 16, 8, 17, 4, 18, 19, 2, 20, 10, 21, 5, 22, 11, 1, 12, 6, 13, 3, 14, 7, 15}, whenGameBoard)
	assert.Equal(t, 6, whenMarbleIndex)
	assert.Equal(t, 32, roundScore)

}

func TestDay9Part1(t *testing.T) {

	gameScoreCalc := func(playerCount, turns int) int {
		players := make([]int, playerCount)
		currentPlayer := 0
		board := []int{0}
		currentIndex := 1
		for i := 1; i <= turns; i++ {
			fmt.Printf("Looking at player %v\n", currentPlayer)

			board, currentIndex, players[currentPlayer] = nextTurn(board, currentIndex, i)
			if currentPlayer >= playerCount-1 {
				currentPlayer = 0
			} else {
				currentPlayer++
			}
		}

		sort.Ints(players)

		return players[playerCount-1]
	}

	assert.Equal(t, 32, gameScoreCalc(9, 25))
	//assert.Equal(t, 8317, gameScoreCalc(10, 1618))
	//	assert.Equal(t, 146373, gameScoreCalc(13, 7999))
	//	assert.Equal(t, 2764, gameScoreCalc(17, 1104))
	//	assert.Equal(t, 54718, gameScoreCalc(21, 6111))
	//	assert.Equal(t, 37305, gameScoreCalc(30, 5807))
}

func TestDay9Part2(t *testing.T) {

}

var day9Input = `458 players; last marble is worth 71307 points`
