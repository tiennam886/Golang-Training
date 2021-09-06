package main

import (
	"fmt"
)

// Assume c is the largest num, then it must be c<=b+a-1 to make a triangle
// So if a,b,c do not make a triangle, the minimum num must be added is c+1-b-a
func main()  {
	var a int
	var b int
	var c int
	_, err := fmt.Scanf("%d %d %d", &a,&b, &c)
	if err != nil {
		return 
	}
	a,b,c = findMax(a,b,c)
	fmt.Println(solve(a,b,c))
}
func findMax(a int, b int, c int) (int, int, int){
	if a>=b && a>=c{
		return a, b, c
	}
	if b>=c && b>=a{
		return b, a, c
	}
	if c>=b && c>=a{
		return c,b,a
	}
	return 0,0,0
}
func solve(a int, b int, c int) int{
	if a >= (b+c) {
		return a+1-b-c
	}else{
		return 0
	}
}