package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var balance float64
	for balance < 20.00 {
		skoka := rand.Intn(3) + 1
		switch skoka {
		case 1:
			balance += 0.05
		case 2:
			balance += 0.10
		case 3:
			balance += 0.25
		}
		fmt.Printf("%.2f\n", balance)
	}

}
