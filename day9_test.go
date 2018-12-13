package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Ring struct {
	GameBoard    *Marble
	CurrentIndex *Marble
	CurrentSize  int
}

type Marble struct {
	Value    int
	Next     *Marble
	Previous *Marble
}

func (r *Ring) Remove(removableMarble *Marble) *Marble {
	rear := removableMarble.Previous
	nextMarble := removableMarble.Next
	rear.Next = nextMarble
	nextMarble.Previous = rear
	r.CurrentSize--
	return removableMarble

}

func (r *Ring) asArray() []int {
	array := make([]int, 0)
	element := r.GameBoard
	for i := 0; i < r.CurrentSize; i++ {
		array = append(array, element.Value)
		element = element.Next
	}

	return array
}

type Board interface {
	nextTurn(turnNumber int) int
}

func (r *Ring) nextTurn(turnId int) int {
	if turnId == 1 {
		r.GameBoard = &Marble{0, nil, nil}
		r.GameBoard.Next = &Marble{Value: 1, Next: r.GameBoard, Previous: r.GameBoard}
		r.CurrentIndex = r.GameBoard.Next
		r.CurrentSize = 2
		fmt.Printf("Done \n")
		return 0
	}

	if turnId%23 == 0 {
		playersScore := turnId
		element := r.CurrentIndex
		for i := 0; i < 7; i++ {
			element = element.Previous
		}
		removedMarble := r.Remove(element)
		playersScore += removedMarble.Value
		r.CurrentIndex = removedMarble.Next

		return playersScore
	}
	newMarbleNext := r.CurrentIndex.Next.Next
	newMarblePrevious := r.CurrentIndex.Next
	newMarble := &Marble{Value: turnId, Next: newMarbleNext, Previous: newMarblePrevious}
	newMarblePrevious.Next = newMarble
	newMarbleNext.Previous = newMarble
	r.CurrentIndex = newMarble
	r.CurrentSize++
	return 0
}

func spinOutBoard(size int) *Ring {
	gameBoard := &Marble{}
	return &Ring{gameBoard, nil, size}
}

func TestTurn(t *testing.T) {

	fmt.Printf("Stage 1\n")
	givenBoard := spinOutBoard(2)
	fmt.Printf("Board spun up\n")
	assert.Equal(t, 0, givenBoard.nextTurn(1))
	assert.Equal(t, []int{0, 1}, givenBoard.asArray())
	assert.Equal(t, 1, givenBoard.CurrentIndex.Value)
	assert.Equal(t, 2, givenBoard.CurrentSize)

	fmt.Printf("Stage 2\n")
	givenBoard = spinOutBoard(1)
	for i := 1; i <= 9; i++ {
		givenBoard.nextTurn(i)
	}
	assert.Equal(t, []int{0, 8, 4, 9, 2, 5, 1, 6, 3, 7}, givenBoard.asArray())
	assert.Equal(t, 9, givenBoard.CurrentIndex.Value)
	assert.Equal(t, 4, givenBoard.CurrentIndex.Previous.Value)
	assert.Equal(t, 10, givenBoard.CurrentSize)
	//
	//	// [6]  0  4  2  5  1 (6) 3
	//	// [7]  0  4  2  5  1  6  3 (7)
	//	givenBoard = []int{0, 4, 2, 5, 1, 6, 3}
	//	andCurrentIndex = 5
	//	whenGameBoard, whenMarbleIndex, roundScore = nextTurn(givenBoard, andCurrentIndex, 7)
	//	assert.Equal(t, []int{0, 4, 2, 5, 1, 6, 3, 7}, whenGameBoard)
	//	assert.Equal(t, 7, whenMarbleIndex)
	//	assert.Equal(t, 0, roundScore)
	//
	//	// [7]  0  4  2  5  1  6  3 (7)
	//	// [8]  0 (8) 4  2  5  1  6  3  7
	//	givenBoard = []int{0, 4, 2, 5, 1, 6, 3, 7}
	//	andCurrentIndex = 7
	//	whenGameBoard, whenMarbleIndex, roundScore = nextTurn(givenBoard, andCurrentIndex, 8)
	//	assert.Equal(t, []int{0, 8, 4, 2, 5, 1, 6, 3, 7}, whenGameBoard)
	//	assert.Equal(t, 1, whenMarbleIndex)
	//	assert.Equal(t, 0, roundScore)
	//
	//[22]  0 16  8 17  4 18   9  19  2 20 10 21  5 (22) 11  1 12  6 13  3 14  7 15
	//[23]  0 16  8 17  4 18 (19)  2 20 10 21  5 22  11   1 12  6 13  3 14  7 15
	givenBoard = spinOutBoard(1)
	for i := 1; i <= 22; i++ {
		givenBoard.nextTurn(i)
	}
	assert.Equal(t, []int{0, 16, 8, 17, 4, 18, 9, 19, 2, 20, 10, 21, 5, 22, 11, 1, 12, 6, 13, 3, 14, 7, 15}, givenBoard.asArray())
	assert.Equal(t, 32, givenBoard.nextTurn(23))

	assert.Equal(t, []int{0, 16, 8, 17, 4, 18, 19, 2, 20, 10, 21, 5, 22, 11, 1, 12, 6, 13, 3, 14, 7, 15}, givenBoard.asArray())
	assert.Equal(t, 19, givenBoard.CurrentIndex.Value)

	fmt.Printf("Stage 3\n")
	//	givenBoard = spinOutBoard(1)
	//	for i := 1; i < 92; i++ {
	//		givenBoard.nextTurn(i)
	//	}
	//	assert.Equal(t, 107, givenBoard.nextTurn(92))
	//
	//	assert.Equal(t, []int{0, 90, 39, 91, 16, 40, 8, 41, 42, 4, 47, 43, 48, 18, 49, 44, 50, 19, 51, 45, 52, 2, 53, 24, 54, 20, 55, 25, 56, 10, 57, 26, 58, 21, 59, 27, 60, 5, 61, 28, 62, 22, 63, 29, 64, 65, 30, 70, 66, 71, 1, 72, 67, 73, 31, 74, 68, 75, 12, 76, 32, 77, 6, 78, 33, 79, 13, 80, 34, 81, 3, 82, 35, 83, 14, 84, 36, 85, 7, 86, 37, 87, 88, 38, 89, 89, 0}, givenBoard.asArray())
	//	assert.Equal(t, 82, givenBoard.CurrentIndex)
	//	assert.Equal(t, 88, givenBoard.CurrentIndex.Value)
	//	assert.Equal(t, 85, givenBoard.CurrentSize)

}

//func printfBoard(round int, board []int, currentIndex int, roundScore int, currentPlayer int) {
//	fmt.Printf("[%5d] [%2d] ", round, currentPlayer)
//	for index, value := range board {
//		if index == currentIndex {
//			fmt.Printf("(%5d)", value)
//		} else {
//			fmt.Printf(" %5d ", value)
//		}
//	}
//	fmt.Printf("\n")
//}
//
func TestDay9Part1(t *testing.T) {

	gameScoreCalc := func(playerCount, turns int) int {
		players := make([]int, playerCount)
		currentPlayer := 0
		board := spinOutBoard(turns)

		for i := 1; i <= turns; i++ {
			if i%10000 == 0 {
				fmt.Printf("Turn %d of %d\n", i, turns)
			}
			players[currentPlayer] += board.nextTurn(i)
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
	assert.Equal(t, 8317, gameScoreCalc(10, 1618))
	assert.Equal(t, 63, gameScoreCalc(3, 47))
	assert.Equal(t, 95, gameScoreCalc(23, 47))
	assert.Equal(t, 165, gameScoreCalc(10, 118))
	assert.Equal(t, 146373, gameScoreCalc(13, 7999))
	assert.Equal(t, 2764, gameScoreCalc(17, 1104))
	assert.Equal(t, 54718, gameScoreCalc(21, 6111))
	assert.Equal(t, 37305, gameScoreCalc(30, 5807))
	assert.Equal(t, 398048, gameScoreCalc(458, 71307))
	assert.Equal(t, 398048, gameScoreCalc(458, 7130700))
}

func TestDay9Part2(t *testing.T) {

}

var day9Input = `458 players; last marble is worth 71307 points`
