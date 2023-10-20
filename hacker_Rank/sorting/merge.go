package main

import "fmt"

type Merge []int

func (m Merge) sort() []int {
	size := len(m)
	if size < 2 {
		return m
	}

	left := m[0 : size/2]
	right := m[size/2 : size]
	fmt.Println(left, right)
	//left.sort()
	//right.sort() // why calling separatly is not working???
	return merge(left.sort(), right.sort())

}

func merge(l, r []int) (merged []int) {
	merged = make([]int, len(l)+len(r))
	i := 0
	for len(l) > 0 && len(r) > 0 {
		if l[0] < r[0] {
			merged[i] = l[0]
			l = l[1:]
		} else {
			merged[i] = r[0]
			r = r[1:]
		}
		i++

	}

	for j := 0; j < len(l); j++ {
		merged[i] = l[j]
		i++
	}
	for j := 0; j < len(r); j++ {
		merged[i] = r[j]
		i++
	}
	fmt.Println(merged)

	return
}
