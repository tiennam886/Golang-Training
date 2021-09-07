package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n   int
	ans int
	min int = 1e9 + 10
	r       = bufio.NewReader(os.Stdin)
	w       = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	const N int = 1e5 + 10

	_, _ = fmt.Fscan(r, &n)
	if n < 2 {
		return
	}
	var queues [N]int
	for i := 1; i <= n; i++ {
		var k int
		_, _ = fmt.Fscan(r, &k)
		queues[i] = k
	}
	for i := 1; i <= n; i++ {
		t := (queues[i] - i + n) / n
		if min > t {
			min = t
			ans = i

		}
	}
    // count and find the minimum rounds that a queue can be 0, prior to initial queue
	fmt.Println(ans)

}
