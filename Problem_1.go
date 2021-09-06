package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    var tableCard string
	fmt.Scanln(&tableCard)
	consoleReader := bufio.NewReader(os.Stdin)
	input, _ := consoleReader.ReadString('\n')
	handCards := strings.Split(input, " ")
	var result string
	result = checkCard(tableCard, handCards)
	fmt.Println(result)
}

func checkCard(tableCard string, handCards []string) string{
	for _, handCard := range handCards{
		if tableCard[0]==handCard[0] || tableCard[1]==handCard[1] {
			return "YES"
		}
	}
	return "NO"
}