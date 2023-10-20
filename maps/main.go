package main

import "fmt"

func main() {
	m := make(map[uint64]uint64)
	m[0] = 1
	m[1] = 2
	x, ok := m[0]
	fmt.Println(x, ok)

	colors := make(map[string]int)
	colors["red"] = 3
	colors["blue"] = 4

	music := map[string]string{
		"Rock":        "Black Sabbat",
		"POP":         "MJ",
		"Trash Metal": "Metalica",
	}

	printMap(colors)
	printMap(music)

	printM(1, 2)
	printM("dfd", "fdfdfd")

}

func printMap[K comparable, V any](m map[K]V) {

	for x, y := range m {
		fmt.Println("Here goes", x, "and", y)
	}
}

func printM[K comparable, V any](k K, v V) {

	fmt.Println(k, v)
}
