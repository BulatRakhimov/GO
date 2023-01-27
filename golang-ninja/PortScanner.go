package main

import (
	"fmt"
	"net"
	"sort"
)

func rab(porty, rezy chan int) {
	for p := range porty {
		adress := fmt.Sprintf("scanme.nmap.org:%v", p)
		podkl, err := net.Dial("tcp", adress)
		if err != nil {
			rezy <- 0
			continue
		}
		podkl.Close()
		rezy <- p
	}
}

func main() {
	porty := make(chan int, 100)
	rezy := make(chan int)
	var otkrporty []int
	for i := 0; i < cap(porty); i++ {
		go rab(porty, rezy)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			porty <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-rezy
		if port != 0 {
			otkrporty = append(otkrporty, port)
		}
	}

	close(porty)
	close(rezy)
	sort.Ints(otkrporty)
	for _, port := range otkrporty {
		fmt.Printf("%v otkrit\n", port)
	}
}
