package main

import (
	"fmt"
	"strings"
)

func main(){
	var n int
	fmt.Scanln(&n)
	var result string
	for i:=0;i<n;i++{
		var length int
		fmt.Scanln(&length)
		var phone string
		fmt.Scanln(&phone)
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
