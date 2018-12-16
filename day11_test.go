package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Coord struct {
	X int
	Y int
}

func TestDay11Part1(t *testing.T) {
	calcPower := func(cellX, cellY, gridSerialNumber int) int {
		rackID := cellX + 10
		powerLevelStart := rackID * cellY
		powerLevelStart += gridSerialNumber
		powerLevelStart *= rackID
		powerLevel := 0
		if powerLevelStart > 99 {
			powerLevel = powerLevelStart / 100
			powerLevel = powerLevel % 10
		}
		powerLevel = powerLevel - 5
		return powerLevel
	}

	assert.Equal(t, 4, calcPower(3, 5, 8))
	assert.Equal(t, -5, calcPower(122, 79, 57))
	assert.Equal(t, 0, calcPower(217, 196, 39))
	assert.Equal(t, 4, calcPower(101, 153, 71))

	grid := make(map[Coord]int)

	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			grid[Coord{x, y}] = calcPower(x, y, 9810)
		}
	}
	highestCell := 0
	var newCoord Coord

	for x := 1; x < 299; x++ {
		for y := 1; y < 299; y++ {
			value := grid[Coord{x, y}]
			value += grid[Coord{x + 1, y}]
			value += grid[Coord{x + 2, y}]
			value += grid[Coord{x, y + 1}]
			value += grid[Coord{x + 1, y + 1}]
			value += grid[Coord{x + 2, y + 1}]
			value += grid[Coord{x, y + 2}]
			value += grid[Coord{x + 1, y + 2}]
			value += grid[Coord{x + 2, y + 2}]
			if value > highestCell {
				newCoord = Coord{x, y}
				highestCell = value
			}
		}
	}

	assert.Equal(t, 29, highestCell)
	assert.Equal(t, Coord{245, 14}, newCoord)

}

func TestDay11Part2(t *testing.T) {
	//Yeeeeaaaaaah, about that....
}

type Answer struct {
	X        int
	Y        int
	GridSize int
	Value    int
}

func processGrid(grid map[Coord]int, gridSize int, channel chan Answer) {
	highestCell := 0
	var newCoord Coord

	for x := 1; x <= (301 - gridSize); x++ {
		for y := 1; y <= (301 - gridSize); y++ {
			value := 0
			for xwidth := 0; xwidth < gridSize; xwidth++ {
				for ywidth := 0; ywidth < gridSize; ywidth++ {
					value += grid[Coord{x + xwidth, y + ywidth}]
				}
			}
			if value > highestCell {
				newCoord = Coord{x, y}
				highestCell = value
			}
		}
	}
	channel <- Answer{newCoord.X, newCoord.Y, gridSize, highestCell}
	fmt.Printf("Processed grid %d\n", gridSize)

}

var day11Input = ``
