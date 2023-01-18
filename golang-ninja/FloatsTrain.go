package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var cents float32
	for cents < 2000 {
		plus := rand.Intn(3) + 1
		switch plus {
		case 1:
			cents += 5
		case 2:
			cents += 10
		case 3:
			cents += 25
		}
		var balance float32
		balance = cents / 100
		fmt.Printf("%.2f\n", balance)
	}

}
