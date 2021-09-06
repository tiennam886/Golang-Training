package main

import "fmt"

func main(){
	var nShop int
	const N int =1e5
	var priceShops [N]int
	var nShopCanBuy [N]int
	fmt.Scanln(&nShop)
	for i:=0; i<nShop;i++{
		var x int
		fmt.Scan(&x)
		priceShops[i]=x
	}
	var nQ int
	fmt.Scanln(&nQ)
	for i:=0; i<nQ;i++{
		var spendCoin int
		var k int
		k=0
		fmt.Scan(&spendCoin)
		for j:=0; j<nShop;j++{
			if priceShops[j]<=spendCoin{
				k+=1
			}
		}
		nShopCanBuy[i]=k
	}
	fmt.Println("\nOutput:")
	for i:=0; i<nQ;i++{
		fmt.Println(nShopCanBuy[i])
	}
}
