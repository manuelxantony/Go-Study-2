package main

import "fmt"

func main() {
	n := 3
	fmt.Println(fact(n))
}

func fact(n int) int {
	if n == 1 {
		return 1
	}

	return n * fact(n-1)
}
