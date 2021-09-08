package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	ans  int64
	n    int
	m    int
	save = make([]int, 3)
	r    = bufio.NewReader(os.Stdin)
	w    = bufio.NewWriter(os.Stdout)
)

func solve() {
	save[0] = 0 // save even num
	save[1] = 0 // save odd num
	ans = 0
	fmt.Fscan(r, &n)
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(r, &x)
		save[x%2]++
		// count even and odd nums in array
	}
	fmt.Fscan(r, &m)
	for i := 1; i <= m; i++ {
		var x int
		fmt.Fscan(r, &x)
		ans += (int64(save[x%2]))
		// ans is number of even pair and odd pair
	}
	fmt.Println(int64(ans))
}
func main() {
	defer w.Flush()
	var te int
	fmt.Fscan(r, &te)
	for i := 1; i <= te; i++ {
		solve()
	}
}
