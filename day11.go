package main

import (
	"fmt"
)

type Coord struct {
	X int
	Y int
}

func main() {
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

	grid := make(map[Coord]int)

	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			grid[Coord{x, y}] = calcPower(x, y, 9810)
		}
	}

	for y := 1; y <= 300; y++ {
		for x := 0; x <= 300; x++ {
			fmt.Printf(" %2d ", grid[Coord{x, y}])
		}
		fmt.Printf("\n")
	}
	var highestAnswer Answer
	channel := make(chan Answer)

	for gridSize := 1; gridSize <= 300; gridSize++ {
		go processGrid(grid, gridSize, channel)
	}

	for gridSize := 1; gridSize <= 300; gridSize++ {
		answer := <-channel
		if answer.Value > highestAnswer.Value {
			highestAnswer = answer
		}
	}

	fmt.Printf("Highest Number %d\n", highestAnswer.Value)
	fmt.Printf("Highest answer %d,%d,%d\n", highestAnswer.X, highestAnswer.Y, highestAnswer.GridSize)
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
