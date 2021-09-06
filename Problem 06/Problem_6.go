package main

import "fmt"

// y = x + p and y = -x + q has intersection point is ( (q-p)/2;(p+q)/2)
// hence the point is int when p+q is divisible by 2
// => result is number of pair p,q has even sum
func main(){
	var t int
	const N int =1e5
	var a [N]int
	fmt.Scanln(&t)
	for i:=0; i<t; i++{
		var nP int
		fmt.Scanln(&nP)
		evenP :=0
		for i:=0; i<nP;i++{
			var x int
			fmt.Scan(&x)
			if x%2==0{
				evenP +=1
			}
		}
		var nQ int
		fmt.Scanln(&nQ)
		evenQ :=0
		for i:=0; i<nQ;i++{
			var x int
			fmt.Scan(&x)
			if x%2==0{
				evenQ +=1
			}
		}
		a[i] = evenQ*evenP+(nP-evenP)*(nQ-evenQ)

	}
	fmt.Println("\nOutput:")
	for i:=0; i<t;i++{
		fmt.Println(a[i])
	}
}
