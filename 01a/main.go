// Fix for: https://adventofcode.com/2016/day/1
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
	// lines, err := lfs.ReadLines("example2.txt") // Read the input
	// lines, err := lfs.ReadLines("example3.txt") // Read the input
	if err != nil {
		panic(err)
	}
	var pos coord // Where we are
	instructions := strings.Fields(lines[0])
	direction := "U" // default start

	for _, step := range instructions {
		step = strings.Trim(step, ",") // Get wrid of ,
		// fmt.Printf("%4v %s %s\n", pos, direction, step)
		direction = changedirection(string(step[0]), direction)
		pos = move(pos, direction, lfs.SilentAtoi(string(step[1:len(step)])))
	}
	fmt.Println(calculateDistance(pos), pos)
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

func move(pos coord, direction string, i int) coord {
	switch direction {
	case "U":
		pos.r += i
	case "R":
		pos.c += i
	case "D":
		pos.r -= i
	case "L":
		pos.c -= i
	}
	return pos
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
