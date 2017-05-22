package main

import (
	"fmt"
)

func main() {

	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		defer close(naturals)
		for x := 1; x <= 100; x++ {
			fmt.Printf("Sending: %d\n", x)
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		defer close(squares)
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			fmt.Printf("Squaring: %d -> %d\n", x, x*x)
			squares <- x * x
		}
	}()

	// Printer
	for i := range squares {
		fmt.Printf("End of pipeline: %d\n", i)
	}
}
