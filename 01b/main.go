// Fix for: https://adventofcode.com/2016/day/1b
package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/CZero/gofuncy/lfs"
)

type coord struct {
	r int
	c int
}

func main() {
	lines, err := lfs.ReadLines("input.txt") // Read the input
	// lines, err := lfs.ReadLines("example1.txt") // Read the input
	if err != nil {
		panic(err)
	}

	instructions := strings.Fields(lines[0])
	var pos coord                      // Where we are
	var visited = make(map[coord]bool) // The places we have been
	direction := "U"                   // default start
	visited[coord{0, 0}] = true

	for _, step := range instructions {
		var succes bool
		step = strings.Trim(step, ",") // Get wrid of ,
		direction = changedirection(string(step[0]), direction)
		pos, visited, succes = move(pos, direction, lfs.SilentAtoi(string(step[1:len(step)])), visited)
		if succes {
			fmt.Printf("Done!\n")
			break
		}
	}
	fmt.Printf("We are standing %v, for the second time. The distance to start is: %d\n", pos, calculateDistance(pos))
}

func move(pos coord, direction string, i int, visited map[coord]bool) (coord, map[coord]bool, bool) {
	var succes bool
	switch direction {
	case "U":
		for j := 0; j < i; j++ {
			pos.r++
			// fmt.Printf("%v Now standing: %v\n", visited, pos)
			if visited[pos] {
				fmt.Printf("The second time we visit %v!\n", pos)
				succes = true
				break
			} else {
				visited[pos] = true
			}
		}
	case "R":
		for j := 0; j < i; j++ {
			pos.c++
			// fmt.Printf("%v Now standing: %v\n", visited, pos)
			if visited[pos] {
				fmt.Printf("The second time we visit %v!\n", pos)
				succes = true
				break
			} else {
				visited[pos] = true
			}
		}
	case "D":
		for j := 0; j < i; j++ {
			pos.r--
			// fmt.Printf("%v Now standing: %v\n", visited, pos)
			if visited[pos] {
				fmt.Printf("The second time we visit %v!\n", pos)
				succes = true
				break
			} else {
				visited[pos] = true
			}
		}
	case "L":
		for j := 0; j < i; j++ {
			pos.c--
			// fmt.Printf("%v Now standing: %v\n", visited, pos)
			if visited[pos] {
				fmt.Printf("The second time we visit %v!\n", pos)
				succes = true
				break
			} else {
				visited[pos] = true
			}
		}
	}
	return pos, visited, succes
}

// changedirection changes direction based on the instruction and the current direction
//
//        U
//       L R
//      R   L
//     L     R
//      L   R
//       R L
//        D
//
func changedirection(instruction, direction string) string {
	switch direction {
	case "U":
		if instruction == "L" {
			return "L"
		} else {
			return "R"
		}
	case "R":
		if instruction == "L" {
			return "U"
		} else {
			return "D"
		}
	case "D":
		if instruction == "L" {
			return "R"
		} else {
			return "L"
		}
	case "L":
		if instruction == "L" {
			return "D"
		} else {
			return "U"
		}
	}
	panic("Wrong direction")
}

func calculateDistance(pos coord) int {
	if pos.c < 0 {
		pos.c = int(math.Abs(float64(pos.c)))
	}
	if pos.r < 0 {
		pos.r = int(math.Abs(float64(pos.r)))
	}
	return pos.c + pos.r
}
