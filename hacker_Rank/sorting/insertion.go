package main

import "fmt"

type Insertion []int

// 6 3 2 1 5

func (in Insertion) sort() []int {
	len := len(in)
	for i := 1; i < len; i++ {
		keyIn := i
		key := in[i]
		for j := i - 1; j >= 0; j-- {
			if key < in[j] {
				in[keyIn] = in[j]
				keyIn = j
			}
			fmt.Println("while sorting:", in)
		}
		in[keyIn] = key
		fmt.Println("while:", in)
	}

	return in
}
