package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Coord struct {
	X int
	Y int
}

func manhattanDistance(c1, c2 Coord) int {
	distance := 0
	if c2.X > c1.X {
		distance = c2.X - c1.X
	} else {
		distance = c1.X - c2.X
	}

	if c2.Y > c1.Y {
		distance += c2.Y - c1.Y
	} else {
		distance += c1.Y - c2.Y
	}
	return distance
}

func TestDay6Part1(t *testing.T) {
	areaFinder := func(input string) int {
		coordMap := make([]Coord, 0)
		for _, line := range strings.Split(input, "\n") {
			parts := strings.Split(line, ", ")
			x, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(fmt.Sprintf("Failed for process x from %v for line %v. Err %v", parts[0], line, err))
			}
			y, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(fmt.Sprintf("Failed for process y from %v for line %v. Err %v", parts[1], line, err))
			}
			coordMap = append(coordMap, Coord{x, y})
		}
		sort.SliceStable(coordMap, func(i, j int) bool {
			return coordMap[i].X < coordMap[j].X
		})

		bottomX := coordMap[0].X
		topX := coordMap[len(coordMap)-1].X

		sort.SliceStable(coordMap, func(i, j int) bool {
			return coordMap[i].Y < coordMap[j].Y
		})

		bottomY := coordMap[0].Y
		topY := coordMap[len(coordMap)-1].Y

		assignedValue := make(map[Coord]int)
		for x := bottomX; x <= topX; x++ {
			for y := bottomY; y <= topY; y++ {
				// who owns this square?
				manhattanMatches := make([]Coord, 0)
				lowestManhattanDistance := topY + topX + 2
				for _, coord := range coordMap {

					distance := manhattanDistance(coord, Coord{x, y})
					if distance == lowestManhattanDistance {
						manhattanMatches = append(manhattanMatches, coord)
					} else if distance < lowestManhattanDistance {
						manhattanMatches = make([]Coord, 1)
						manhattanMatches[0] = coord
						lowestManhattanDistance = distance
					}

				}
				if len(manhattanMatches) == 1 {
					assignedValue[manhattanMatches[0]] += 1

				}
			}
		}
		highighestOwnership := -1
		for coord, ownership := range assignedValue {
			if ownership > highighestOwnership {
				var bl, tl, br, tr *Coord
				// check there is a bottom left, a bottom right, a top left and a top right
				for _, potentialCorner := range coordMap {
					if potentialCorner.X < coord.X && potentialCorner.Y < coord.Y {
						tl = &potentialCorner
					} else if potentialCorner.X < coord.X && potentialCorner.Y > coord.Y {
						bl = &potentialCorner
					} else if potentialCorner.X > coord.X && potentialCorner.Y < coord.Y {
						tr = &potentialCorner
					} else if potentialCorner.X > coord.X && potentialCorner.Y > coord.Y {
						br = &potentialCorner
					}
				}
				if tl != nil && tr != nil && bl != nil && br != nil {
					highighestOwnership = ownership
				}
			}
		}

		return highighestOwnership
	}

	testInput := `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`

	assert.Equal(t, 17, areaFinder(testInput))
	assert.Equal(t, 3989, areaFinder(day6Input))

}

func TestDay6Part2(t *testing.T) {
	regionFinder := func(input string, maxDistance int) int {
		coordMap := make([]Coord, 0)
		for _, line := range strings.Split(input, "\n") {
			parts := strings.Split(line, ", ")
			x, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(fmt.Sprintf("Failed for process x from %v for line %v. Err %v", parts[0], line, err))
			}
			y, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(fmt.Sprintf("Failed for process y from %v for line %v. Err %v", parts[1], line, err))
			}
			coordMap = append(coordMap, Coord{x, y})
		}
		sort.SliceStable(coordMap, func(i, j int) bool {
			return coordMap[i].X < coordMap[j].X
		})

		bottomX := coordMap[0].X
		topX := coordMap[len(coordMap)-1].X

		sort.SliceStable(coordMap, func(i, j int) bool {
			return coordMap[i].Y < coordMap[j].Y
		})

		acceptableSquares := make([]Coord, 0)
		bottomY := coordMap[0].Y
		topY := coordMap[len(coordMap)-1].Y

		for x := bottomX; x <= topX; x++ {
			for y := bottomY; y <= topY; y++ {
				totalManhattanDistance := 0
				for _, coord := range coordMap {
					totalManhattanDistance += manhattanDistance(coord, Coord{x, y})

				}
				if totalManhattanDistance < maxDistance {
					acceptableSquares = append(acceptableSquares, Coord{x, y})
				}
			}
		}

		return len(acceptableSquares)
	}
	testInput := `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`
	assert.Equal(t, 16, regionFinder(testInput, 32))
	assert.Equal(t, 49715, regionFinder(day6Input, 10000))
}

var day6Input = `224, 153
176, 350
353, 241
207, 59
145, 203
123, 210
113, 203
191, 241
172, 196
209, 249
260, 229
98, 231
305, 215
258, 141
337, 282
156, 140
325, 197
179, 279
283, 233
317, 150
305, 245
67, 109
251, 140
245, 59
173, 105
59, 173
257, 70
269, 110
102, 162
179, 180
324, 112
357, 311
317, 245
239, 112
321, 220
133, 97
334, 99
117, 102
133, 112
222, 316
68, 296
150, 287
263, 263
66, 347
128, 118
63, 202
68, 236
264, 122
77, 243
92, 110`
