package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N int64 = 1e5 + 10

var (
	n     int
	m     int
	nShop [N]int
	r     = bufio.NewReader(os.Stdin)
	w     = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	solved()
}
func solved() {
	fmt.Fscan(r, &n)
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(r, &x)
		nShop[i] = x
	}
	sort.Ints(nShop[0:n])
	// sort min->max
	fmt.Fscan(r, &m)
	for i := 0; i < m; i++ {
		var x int
		fmt.Fscan(r, &x)
		l := 0
		r := n

		// find the minimum position of x that x >= sorted_nShop[i]
		fmt.Println(upperbound(l, r, x, nShop[0:n]))

	}
}
func upperbound(l int, r int, x int, nShop []int) int {
	var mid int
	for l < r {
		mid = (l + r) / 2
		if nShop[mid] <= x {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
