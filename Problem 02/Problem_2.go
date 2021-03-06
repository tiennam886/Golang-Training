package main

import "fmt"

// n is number of employees
// l is number of leaders
// then n-l is number of non-leader employees
// hence, n-l is divisible by l

func main() {
	var num int
	fmt.Scanln(&num)
	if num < 1 {
		return
	}
	fmt.Println(checkDivisible(num))
	return

}
func checkDivisible(num int) int {
	way := 0
	for leader := 1; leader <= num/2; leader++ {
		if (num-leader)%leader == 0 {
			way += 1
		}
	}
	return way
}
