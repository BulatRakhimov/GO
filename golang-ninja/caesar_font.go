package main

import "fmt"

func main() {
	message := "L fdph, L vdz, L frqtxhuhg"
	for i := 0; i < len(message); i++ {
		c := message[i]
		e := c - 3
		if c == 'a' {
			e = 'z'
		}
		if c == 'A' {
			e = 'Z'
		}
		if c == ' ' {
			e = c
		}

		fmt.Printf("%c", e)
	}
}
