package main

type Selection []int

func (s Selection) sort() []int {
	ls := len(s)
	for i := 0; i < ls; i++ {
		minIdx := i
		for j := i + 1; j < ls; j++ {
			if s[j] < s[minIdx] {
				minIdx = j
			}
		}
		s[minIdx], s[i] = s[i], s[minIdx]
	}
	return s
}
