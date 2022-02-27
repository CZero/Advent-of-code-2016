package main

import (
	"fmt"

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
	keypad := InitKeypad()
	code := followInstructions(keypad, codeinstructions)
	fmt.Println(code)

}

// followInstructions walks through al the instructions, getting all the digits for a code.
func followInstructions(keypad map[coord]string, codeinstructions []string) (code string) {
	pos := coord{2, 0}
	for _, instruction := range codeinstructions {
		for _, direction := range instruction {
			pos = followDirection(string(direction), pos, keypad)
		}
		code += keypad[pos]
		// fmt.Println(pos, " = ", keypad[pos])
	}
	return code
}

func followDirection(direction string, pos coord, keypad map[coord]string) coord {
	switch direction {
	case "U":
		if _, ok := keypad[coord{pos.c, pos.r + 1}]; ok {
			pos.r++
		}
	case "L":
		if _, ok := keypad[coord{pos.c - 1, pos.r}]; ok {
			pos.c--
		}
	case "R":
		if _, ok := keypad[coord{pos.c + 1, pos.r}]; ok {
			pos.c++
		}
	case "D":
		if _, ok := keypad[coord{pos.c, pos.r - 1}]; ok {
			pos.r--
		}
	}
	return pos
}

// InitKeypad builds the keypad map
func InitKeypad() map[coord]string {
	var panel = make(map[coord]string)
	panellines, err := lfs.ReadLines("panel.txt")
	lfs.PanErr(err)
	fmt.Printf("Panel:\n")
	revlen := len(panellines) - 1 // used for reversing the keypad, this will help with ease of reading up and down.
	for i, panelline := range panellines {
		for j, digit := range panelline {
			if string(digit) != " " {
				panel[coord{revlen - i, j}] = string(digit)
			}
			fmt.Printf("%d,%d=%s   ", revlen-i, j, string(digit))
		}
		fmt.Println()
	}
	return panel
}
