package main

import (
	"fmt"
	"strings"

	"github.com/CZero/gofuncy/lfs"
)

func main() {
	// triangles, err := lfs.ReadLines("example.txt")
	triangles, err := lfs.ReadLines("input.txt")

	lfs.PanErr(err)
	var unvalids, valids int
	for _, traingle := range triangles {
		valid := validateTriangle(traingle)
		fmt.Printf("%s is a valid triangle? %t\n", traingle, valid)
		if !valid {
			unvalids++
		} else {
			valids++
		}
	}
	fmt.Printf("%d unvalid traingles, %d valid triangles..", unvalids, valids)

}

func validateTriangle(traingle string) bool {
	sides := strings.Fields(traingle)
	a, b, c := lfs.SilentAtoi(sides[0]), lfs.SilentAtoi(sides[1]), lfs.SilentAtoi(sides[2])
	return a+b > c && b+c > a && a+c > b
}
