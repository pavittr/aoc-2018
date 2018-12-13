package main

import (
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

type Board interface {
	nextTurn(turnNumber int) int
}

func (r *Ring) nextTurn(turnId int) int {
	if turnId == 1 {
		r.GameBoard = &Marble{0, nil, nil}
		r.GameBoard.Next = &Marble{Value: 1, Next: r.GameBoard, Previous: r.GameBoard}
		r.CurrentIndex = r.GameBoard.Next
		r.CurrentSize = 2
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

func TestDay9Part1(t *testing.T) {

	gameScoreCalc := func(playerCount, turns int) int {
		players := make([]int, playerCount)
		currentPlayer := 0
		board := &Ring{nil, nil, 0}

		for i := 1; i <= turns; i++ {
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
	assert.Equal(t, 3180373421, gameScoreCalc(458, 7130700))
}
