package main

type Bubble []int

// using swaped for while loop
func (b Bubble) sort() []int {
	lb := len(b)

	for i := 0; i < lb-1; i++ {
		swaped := false
		//for swaped {
		//swaped = false
		for j := 0; j < lb-1-i; j++ {
			if b[j] > b[j+1] {
				b[j], b[j+1] = b[j+1], b[j]
				swaped = true
			}
		}
		if !swaped {
			break
		}
		//lb -= 1
	}
	return b
}
