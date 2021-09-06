package main

import (
	"fmt"
	"strings"
)

func main(){
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println("1")
	}
	var result string
	for i:=0;i<n;i++{
		var length int
		_, err2 := fmt.Scanln(&length)
		if err2 != nil {
			fmt.Println("2")
		}
		var phone string
		_, err3 := fmt.Scanln(&phone)
		if err3 != nil {
			fmt.Println("3")
		}
		result += "\n"+checkPhoneNum(length,phone)
	}
	fmt.Println(result)
}
func checkPhoneNum(length int, phone string) string{
	first8Index := strings.Index(phone, "8")
	if first8Index!=-1 && (length-first8Index) >=11{
		return "YES"
	}

	return "NO"
}

