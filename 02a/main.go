package main

import (
	"fmt"
	"strings"

	"github.com/CZero/gofuncy/lfs"
)

type coord struct {
	r int // Row
	c int // Col
}

func main() {
	codeinstructions, err := lfs.ReadLines("input.txt")
	// codeinstructions, err := lfs.ReadLines("example1.txt")
	lfs.PanErr(err)
	_ = codeinstructions
	keypad, upperbound, rightbound := InitKeypad()
	// fmt.Printf("%v\n", keypad)
	code := followInstructions(keypad, codeinstructions, upperbound, rightbound)
	fmt.Println(code)

}

// followInstructions walks through al the instructions, getting all the digits for a code.
func followInstructions(keypad map[coord]int, codeinstructions []string, upperbound, rightbound int) (code []int) {
	pos := coord{1, 1}
	for _, instruction := range codeinstructions {
		for _, direction := range instruction {
			pos = followDirection(string(direction), pos, upperbound, rightbound)
		}
		code = append(code, keypad[pos])
		fmt.Println(pos)
	}
	return code
}

func followDirection(direction string, pos coord, upperbound, rightbound int) coord {
	switch direction {
	case "U":
		if pos.r < upperbound {
			pos.r++
		}
	case "L":
		if pos.c > 0 {
			pos.c--
		}
	case "R":
		if pos.c < rightbound {
			pos.c++
		}
	case "D":
		if pos.r > 0 {
			pos.r--
		}
	}
	return pos
}

// InitKeypad builds the keypad map, I expect 1b to be a different layout as he imagined it while walking
func InitKeypad() (map[coord]int, int, int) {
	var panel = make(map[coord]int)
	panellines, err := lfs.ReadLines("panel.txt")
	lfs.PanErr(err)
	fmt.Printf("Panel:\n")
	revlen := len(panellines) - 1 // used for reversing the keypad, this will help with ease of reading up and down.
	rightbound := len(panellines[0]) / 2
	for i, line := range panellines {
		digits := strings.Fields(line)
		for j, digit := range digits {
			panel[coord{revlen - i, j}] = lfs.SilentAtoi(digit)
			fmt.Printf("%d,%d = %s    ", revlen-i, j, digit)
		}
		fmt.Printf("\n")
	}
	fmt.Println()
	fmt.Println(rightbound)
	return panel, revlen, rightbound
}
