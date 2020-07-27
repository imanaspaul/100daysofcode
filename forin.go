package main

import "fmt"

func main() {

	// for is the most powerful control logic in Go. It can read data in loops and iterative operations, just like while .
	var sum int = 0
	for sum < 100 {
		sum += 1
	}
	fmt.Println("Total sum is :", sum)
}
