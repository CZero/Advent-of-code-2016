package main

import (
	"fmt"
	"strings"

	"github.com/CZero/gofuncy/lfs"
)

func main() {
	// triangles, err := lfs.ReadLines("example.txt")
	lines, err := lfs.ReadLines("input.txt")
	lfs.PanErr(err)

	triangles := getTriangles(lines)
	var unvalids, valids int
	for _, triangle := range triangles {
		valid := validateTriangle(triangle)
		fmt.Printf("%d is a valid triangle? %t\n", triangle, valid)
		if !valid {
			unvalids++
		} else {
			valids++
		}
	}
	fmt.Printf("%d unvalid traingles, %d valid triangles..", unvalids, valids)

}

func getTriangles(lines []string) [][]int {
	var triangles [][]int
	var heap1, heap2, heap3 []int
	for _, line := range lines {
		nums := strings.Fields(line)
		heap1 = append(heap1, lfs.SilentAtoi(nums[0]))
		heap2 = append(heap2, lfs.SilentAtoi(nums[1]))
		heap3 = append(heap3, lfs.SilentAtoi(nums[2]))
		if len(heap1) == 3 {
			triangles = append(triangles, heap1)
			triangles = append(triangles, heap2)
			triangles = append(triangles, heap3)
			heap1, heap2, heap3 = nil, nil, nil
		}
	}
	fmt.Println(triangles)
	return triangles
}

func validateTriangle(traingle []int) bool {
	a, b, c := traingle[0], traingle[1], traingle[2]
	return a+b > c && b+c > a && a+c > b
}
