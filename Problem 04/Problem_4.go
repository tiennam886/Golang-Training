package main

import (
	"fmt"
	"sort"
)

// Assume c is the largest num, then it must be c<=b+a-1 to make a triangle
// So if a,b,c do not make a triangle, the minimum num must be added is c+1-b-a

var (
	a int
	b int
	c int
)

func main() {
	fmt.Scanf("%d %d %d", &a, &b, &c)
	arr := [3]int{a, b, c}
	sort.Ints(arr[:])
	fmt.Println(solve(arr[0], arr[1], arr[2]))
}
func solve(a int, b int, c int) int {
	if c >= (a + b) {
		return c + 1 - b - a
	}
	return 0

}
