package main

import (
	"fmt"
)

func main() {
	var nestedSize int
	fmt.Scanln(&nestedSize)
	if nestedSize < 1 {
		return
	} 
	nestedSlice := initMatrix(nestedSize)
	fmt.Println(Solve(nestedSlice, nestedSize))
	return
}

func initMatrix(nestedSize int) [][]int{
	var nestedSlice [][]int
	for i:=0;i<nestedSize; i++ {
		var jSlice []int
		for j:=0;j<nestedSize; j++ {
			jSlice = append(jSlice, 1)
		}
		nestedSlice = append(nestedSlice, jSlice)
	}
	return nestedSlice
}


func Solve(nestedSlice [][]int, nestedSize int) int{
	for i:=1;i<nestedSize;i++ {
		for j:=1;j<nestedSize;j++{
			nestedSlice[i][j] =nestedSlice[i][j-1]+nestedSlice[i-1][j]
		}
	}
	return nestedSlice[nestedSize-1][nestedSize-1]
}
