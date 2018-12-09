package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay7Part1(t *testing.T) {

	stepOrder := func(input string) string {

		orderedList := make([]string, 0)
		stepParser := regexp.MustCompile(`Step (?P<dependency>[[:upper:]]) must be finished before step (?P<stepLetter>[[:upper:]]) can begin.`)
		dependencies := make(map[string][]string)

		for _, line := range strings.Split(input, "\n") {

			if !stepParser.MatchString(line) {
				panic(fmt.Sprintf("Message is unreadable. line %+v", line))
			}

			groups := stepParser.FindStringSubmatch(line)
			dependency := groups[1]
			stepLetter := groups[2]
			stepDeps, alreadySeen := dependencies[stepLetter]
			if !alreadySeen {
				stepDeps = make([]string, 0)
			}
			if _, seenDepToo := dependencies[dependency]; !seenDepToo {
				dependencies[dependency] = make([]string, 0)
			}

			stepDeps = append(stepDeps, dependency)
			dependencies[stepLetter] = stepDeps
		}

		// find what is currently available
		actioned := make(map[string]bool)
		for {
			availableSteps := make([]string, 0)

			for step, deps := range dependencies {
				allDepsActioned := true
				for _, dep := range deps {
					if !actioned[dep] {
						allDepsActioned = false
						break
					}
				}

				if allDepsActioned && !actioned[step] {
					availableSteps = append(availableSteps, step)
				}
			}

			if len(availableSteps) > 0 {
				sort.Strings(availableSteps)
				stepToAction := availableSteps[0]
				actioned[stepToAction] = true
				orderedList = append(orderedList, stepToAction)
			} else {
				break
			}

		}

		return strings.Join(orderedList, "")
	}
	testInput := `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`

	assert.Equal(t, "CABDFE", stepOrder(testInput))
	assert.Equal(t, "CQSWKZFJONPBEUMXADLYIGVRHT", stepOrder(day7Input))
}

func TestDay7Part2(t *testing.T) {
	stepOrder := func(input string, workerCount, durationoffset int) int {

		stepParser := regexp.MustCompile(`Step (?P<dependency>[[:upper:]]) must be finished before step (?P<stepLetter>[[:upper:]]) can begin.`)
		dependencies := make(map[string][]string)
		timeToDone := make(map[string]int)

		for _, line := range strings.Split(input, "\n") {

			if !stepParser.MatchString(line) {
				panic(fmt.Sprintf("Message is unreadable. line %+v", line))
			}

			groups := stepParser.FindStringSubmatch(line)
			dependency := groups[1]
			stepLetter := groups[2]
			stepDeps, alreadySeen := dependencies[stepLetter]
			if !alreadySeen {
				stepDeps = make([]string, 0)
			}
			if _, seenDepToo := dependencies[dependency]; !seenDepToo {
				dependencies[dependency] = make([]string, 0)
			}

			stepDeps = append(stepDeps, dependency)
			dependencies[stepLetter] = stepDeps
		}

		for step := range dependencies {
			switch step {
			case "A":
				timeToDone[step] = (durationoffset + 1)
			case "B":
				timeToDone[step] = (durationoffset + 2)
			case "C":
				timeToDone[step] = (durationoffset + 3)
			case "D":
				timeToDone[step] = (durationoffset + 4)
			case "E":
				timeToDone[step] = (durationoffset + 5)
			case "F":
				timeToDone[step] = (durationoffset + 6)
			case "G":
				timeToDone[step] = (durationoffset + 7)
			case "H":
				timeToDone[step] = (durationoffset + 8)
			case "I":
				timeToDone[step] = (durationoffset + 9)
			case "J":
				timeToDone[step] = (durationoffset + 10)
			case "K":
				timeToDone[step] = (durationoffset + 11)
			case "L":
				timeToDone[step] = (durationoffset + 12)
			case "M":
				timeToDone[step] = (durationoffset + 13)
			case "N":
				timeToDone[step] = (durationoffset + 14)
			case "O":
				timeToDone[step] = (durationoffset + 15)
			case "P":
				timeToDone[step] = (durationoffset + 16)
			case "Q":
				timeToDone[step] = (durationoffset + 17)
			case "R":
				timeToDone[step] = (durationoffset + 18)
			case "S":
				timeToDone[step] = (durationoffset + 19)
			case "T":
				timeToDone[step] = (durationoffset + 20)
			case "U":
				timeToDone[step] = (durationoffset + 21)
			case "V":
				timeToDone[step] = (durationoffset + 22)
			case "W":
				timeToDone[step] = (durationoffset + 23)
			case "X":
				timeToDone[step] = (durationoffset + 24)
			case "Y":
				timeToDone[step] = (durationoffset + 25)
			case "Z":
				timeToDone[step] = (durationoffset + 26)
			default:
				panic("Not looking good for step " + step)
			}
		}

		tick := 0
		workedOnPreviousTick := make(map[string]bool)
		for {
			availableSteps := make([]string, 0)

			for step, deps := range dependencies {
				allDepsActioned := true
				for _, dep := range deps {
					if timeToDone[dep] > 0 {
						allDepsActioned = false
						break
					}
				}

				if allDepsActioned && timeToDone[step] > 0 {
					availableSteps = append(availableSteps, step)
				}
			}

			if len(availableSteps) < 1 {
				break
			}

			sort.Strings(availableSteps)
			freeWorkers := workerCount
			for _, worker := range workedOnPreviousTick {
				if worker {
					freeWorkers--
				}
			}
			for _, step := range availableSteps {
				if workedOnPreviousTick[step] {
					timeToDone[step] -= 1
					if timeToDone[step] == 0 {
						workedOnPreviousTick[step] = false
					}
				} else if freeWorkers > 0 {
					freeWorkers--
					timeToDone[step] -= 1
					if timeToDone[step] > 0 {
						workedOnPreviousTick[step] = true
					}
				}
			}

			tick++
		}

		return tick
	}
	testInput := `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`

	assert.Equal(t, 15, stepOrder(testInput, 2, 0))
	assert.Equal(t, 914, stepOrder(day7Input, 5, 60))
}

var day7Input = `Step Q must be finished before step O can begin.
Step Z must be finished before step G can begin.
Step W must be finished before step V can begin.
Step C must be finished before step X can begin.
Step O must be finished before step E can begin.
Step K must be finished before step N can begin.
Step P must be finished before step I can begin.
Step X must be finished before step D can begin.
Step N must be finished before step E can begin.
Step F must be finished before step A can begin.
Step U must be finished before step Y can begin.
Step M must be finished before step H can begin.
Step J must be finished before step B can begin.
Step B must be finished before step E can begin.
Step S must be finished before step L can begin.
Step A must be finished before step L can begin.
Step E must be finished before step L can begin.
Step L must be finished before step G can begin.
Step D must be finished before step I can begin.
Step Y must be finished before step I can begin.
Step I must be finished before step G can begin.
Step G must be finished before step R can begin.
Step V must be finished before step T can begin.
Step R must be finished before step H can begin.
Step H must be finished before step T can begin.
Step S must be finished before step E can begin.
Step C must be finished before step E can begin.
Step P must be finished before step T can begin.
Step I must be finished before step H can begin.
Step O must be finished before step P can begin.
Step M must be finished before step L can begin.
Step S must be finished before step D can begin.
Step P must be finished before step D can begin.
Step P must be finished before step R can begin.
Step I must be finished before step R can begin.
Step Y must be finished before step G can begin.
Step Q must be finished before step L can begin.
Step N must be finished before step R can begin.
Step J must be finished before step E can begin.
Step N must be finished before step T can begin.
Step B must be finished before step V can begin.
Step Q must be finished before step B can begin.
Step J must be finished before step H can begin.
Step F must be finished before step B can begin.
Step W must be finished before step X can begin.
Step S must be finished before step T can begin.
Step J must be finished before step G can begin.
Step O must be finished before step R can begin.
Step K must be finished before step B can begin.
Step Z must be finished before step O can begin.
Step Q must be finished before step S can begin.
Step K must be finished before step V can begin.
Step B must be finished before step R can begin.
Step J must be finished before step T can begin.
Step E must be finished before step T can begin.
Step G must be finished before step V can begin.
Step D must be finished before step Y can begin.
Step M must be finished before step Y can begin.
Step F must be finished before step G can begin.
Step C must be finished before step P can begin.
Step V must be finished before step R can begin.
Step R must be finished before step T can begin.
Step J must be finished before step Y can begin.
Step U must be finished before step R can begin.
Step Z must be finished before step F can begin.
Step Q must be finished before step V can begin.
Step U must be finished before step M can begin.
Step J must be finished before step R can begin.
Step L must be finished before step V can begin.
Step W must be finished before step K can begin.
Step B must be finished before step Y can begin.
Step O must be finished before step N can begin.
Step D must be finished before step V can begin.
Step P must be finished before step B can begin.
Step U must be finished before step I can begin.
Step O must be finished before step T can begin.
Step S must be finished before step G can begin.
Step X must be finished before step A can begin.
Step U must be finished before step T can begin.
Step A must be finished before step I can begin.
Step B must be finished before step G can begin.
Step N must be finished before step Y can begin.
Step Z must be finished before step J can begin.
Step M must be finished before step D can begin.
Step U must be finished before step A can begin.
Step S must be finished before step R can begin.
Step Z must be finished before step A can begin.
Step Y must be finished before step R can begin.
Step E must be finished before step Y can begin.
Step N must be finished before step G can begin.
Step Z must be finished before step X can begin.
Step P must be finished before step X can begin.
Step Z must be finished before step T can begin.
Step Z must be finished before step P can begin.
Step V must be finished before step H can begin.
Step P must be finished before step L can begin.
Step L must be finished before step H can begin.
Step X must be finished before step V can begin.
Step W must be finished before step G can begin.
Step N must be finished before step D can begin.
Step Z must be finished before step U can begin.`
