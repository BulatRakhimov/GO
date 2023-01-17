package main

import (
	"fmt"
	"math/rand"
)

func main() {
	i := 1
	var tab1 string = "========================================"
	fmt.Printf("%-18s%3s %-11s Price\n%v\n", "Spaceline", "Days", "Trip Type", tab1)
	for i < 11 {
		x := rand.Intn(3) + 1
		var company string
		switch x {
		case 1:
			company = "Virgin Galactic"
		case 2:
			company = "SpaceX"
		case 3:
			company = "Space Adventures"
		}
		y := rand.Intn(2) + 1
		var trip string
		switch y {
		case 1:
			trip = "One-way"
		case 2:
			trip = "Round-trip"
		}
		z := rand.Intn(15)
		skorost := 16 + z                             //km/sec
		stoimost := 36 + z                            //mln dollarrov
		vremya := (62100000 / skorost) / 60 / 60 / 24 //days
		if y == 2 {
			vremya *= 2
			stoimost *= 2
		}
		fmt.Printf("%-19v%3v %-11v $%4v\n", company, vremya, trip, stoimost)
		i++
	}
}
