package main

import "fmt"

func main(){
	const N int = 1e5
	var n int
	fmt.Scanln(&n)
	if n<2{
		return
	}

	var queues [N]int
	for i:=0;i<n;i++{
		var k int
		fmt.Scan(&k)
		queues[i]=k
	}
	remain := queues[0]
	currentPos :=0
	timer := 0
	for remain > 0{
		timer +=1
		remain = queues[currentPos+1]-timer
		currentPos = (currentPos+1)%n
	}
	fmt.Println("\nOutput:")
	fmt.Println(currentPos+1)

}