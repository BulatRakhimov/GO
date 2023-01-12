package main

import (
	"fmt"
)

func main() {
	var god = 2000
	if god%400 == 0 || (god%4 == 0 && god%100 != 0) {
		fmt.Println("wash god visokosniy suka")
	} else {
		fmt.Println("god nevisokosny suka")
	}
}
