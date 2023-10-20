package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	numbs := []int{6, 3, 2, 1, 5}
	//numbs := generateSlice(38)
	start := time.Now()
	//fmt.Println(Bubble(numbs).sort())
	//fmt.Println(Selection(numbs).sort())
	//fmt.Println(Insertion(numbs).sort())
	fmt.Println(Merge(numbs).sort())
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
