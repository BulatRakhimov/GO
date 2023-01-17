package main

import "fmt"

func main() {
	var summa int = 1
	var prib int = 1
	var skoka int = 1
	for summa <= 4000000000 {
		prib = prib * 2
		summa = summa + prib
		fmt.Println(prib)
		skoka++
	}
	fmt.Println(skoka)

}
