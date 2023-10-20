package main

import (
	"fmt"
	"graph/internal"
)

// hmmm, see why
type Graph internal.Graph

type Stack struct {
	internal.Stack
}

type Queue struct {
	internal.Queue
}

// type QueueNode internal.QueueNode

type VDQueue struct {
	internal.VDQueue
}

func main() {
	// g := &Graph{
	// 	Map: map[string][]string{
	// 		"a": {"b", "c"},
	// 		"b": {"d"},
	// 		"c": {"e"},
	// 		"d": {"f"},
	// 		"e": {},
	// 		"f": {},
	// 	},
	// }
	// g.DFS("a")
	// g.BFS("a") // abcdef

	// problems
	// g := &Graph{
	// 	Map: map[string][]string{
	// 		"f": {"g", "i"},
	// 		"g": {"h"},
	// 		"h": {},
	// 		"i": {"g", "k"},
	// 		"j": {"i"},
	// 		"k": {},
	// 	},
	// }
	// fmt.Println(hasPath(g, "f", "j"))

	// undirected
	// edges := [][]string{
	// 	{"i", "j"},
	// 	{"k", "i"},
	// 	{"m", "k"},
	// 	{"k", "l"},
	// 	{"o", "n"},
	// }
	// fmt.Println(undirectedPathDitection(edges, "i", "o"))

	// connected components count
	// graph := Graph{
	// 	Map: map[string][]string{
	// 		"0": {"8", "1", "5"},
	// 		"1": {"0"},
	// 		"5": {"0", "8"},
	// 		"8": {"0", "5"},
	// 		"2": {"3", "4"},
	// 		"3": {"2", "4"},
	// 		"4": {"3", "2"},
	// 	},
	// }
	// fmt.Println(connectedComponents(graph))

	// largest component
	// graph = Graph{
	// 	Map: map[string][]string{
	// 		"0": {"8", "1", "5"},
	// 		"1": {"0"},
	// 		"5": {"0", "8"},
	// 		"8": {"0", "5"},
	// 		"2": {"3", "4"},
	// 		"3": {"2", "4"},
	// 		"4": {"3", "2"},
	// 	},
	// }
	// fmt.Println(largestComponent(graph))

	// shortest path
	// edges := [][]string{
	// 	{"w", "x"},
	// 	{"x", "y"},
	// 	{"z", "y"},
	// 	{"z", "v"},
	// 	{"w", "v"},
	// 	{"w", "i"},
	// 	{"i", "j"},
	// 	{"j", "k"},
	// 	{"k", "l"},
	// 	{"l", "m"},
	// }
	// fmt.Println(shortedPath(edges, "w", "p"))

	// grid := [][]string{
	// 	{"W", "L", "W", "W", "W"},
	// 	{"W", "L", "W", "W", "W"},
	// 	{"W", "W", "W", "L", "W"},
	// 	{"W", "W", "L", "L", "W"},
	// 	{"L", "W", "W", "L", "L"},
	// 	{"L", "L", "W", "W", "W"},
	// }
	grid2 := [][]string{{"1", "0", "0", "1", "1", "1", "0", "1", "1", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
		{"1", "0", "0", "1", "1", "0", "0", "1", "0", "0", "0", "1", "0", "1", "0", "1", "0", "0", "1", "0"},
		{"0", "0", "0", "1", "1", "1", "1", "0", "1", "0", "1", "1", "0", "0", "0", "0", "1", "0", "1", "0"},
		{"0", "0", "0", "1", "1", "0", "0", "1", "0", "0", "0", "1", "1", "1", "0", "0", "1", "0", "0", "1"},
		{"0", "0", "0", "0", "0", "0", "0", "1", "1", "1", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
		{"1", "0", "0", "0", "0", "1", "0", "1", "0", "1", "1", "0", "0", "0", "0", "0", "0", "1", "0", "1"},
		{"0", "0", "0", "1", "0", "0", "0", "1", "0", "1", "0", "1", "0", "1", "0", "1", "0", "1", "0", "1"},
		{"0", "0", "0", "1", "0", "1", "0", "0", "1", "1", "0", "1", "0", "1", "1", "0", "1", "1", "1", "0"},
		{"0", "0", "0", "0", "1", "0", "0", "1", "1", "0", "0", "0", "0", "1", "0", "0", "0", "1", "0", "1"},
		{"0", "0", "1", "0", "0", "1", "0", "0", "0", "0", "0", "1", "0", "0", "1", "0", "0", "0", "1", "0"},
		{"1", "0", "0", "1", "0", "0", "0", "0", "0", "0", "0", "1", "0", "0", "1", "0", "1", "0", "1", "0"},
		{"0", "1", "0", "0", "0", "1", "0", "1", "0", "1", "1", "0", "1", "1", "1", "0", "1", "1", "0", "0"},
		{"1", "1", "0", "1", "0", "0", "0", "0", "1", "0", "0", "0", "0", "0", "0", "1", "0", "0", "0", "1"},
		{"0", "1", "0", "0", "1", "1", "1", "0", "0", "0", "1", "1", "1", "1", "1", "0", "1", "0", "0", "0"},
		{"0", "0", "1", "1", "1", "0", "0", "0", "1", "1", "0", "0", "0", "1", "0", "1", "0", "0", "0", "0"},
		{"1", "0", "0", "1", "0", "1", "0", "0", "0", "0", "1", "0", "0", "0", "1", "0", "1", "0", "1", "1"},
		{"1", "0", "1", "0", "0", "0", "0", "0", "0", "1", "0", "0", "0", "1", "0", "1", "0", "0", "0", "0"},
		{"0", "1", "1", "0", "0", "0", "1", "1", "1", "0", "1", "0", "1", "0", "1", "1", "1", "1", "0", "0"},
		{"0", "1", "0", "0", "0", "0", "1", "1", "0", "0", "1", "0", "1", "0", "0", "1", "0", "0", "1", "1"},
		{"0", "0", "0", "0", "0", "0", "1", "1", "1", "1", "0", "1", "0", "0", "0", "1", "1", "0", "0", "0"}}
	fmt.Println(islandCount(grid2))
}

/*
{{"1","0","0","1","1","1","0","1","1","0","0","0","0","0","0","0","0","0","0","0"},
 {"1","0","0","1","1","0","0","1","0","0","0","1","0","1","0","1","0","0","1","0"},
 {"0","0","0","1","1","1","1","0","1","0","1","1","0","0","0","0","1","0","1","0"},
 {"0","0","0","1","1","0","0","1","0","0","0","1","1","1","0","0","1","0","0","1"},
 {"0","0","0","0","0","0","0","1","1","1","0","0","0","0","0","0","0","0","0","0"},
 {"1","0","0","0","0","1","0","1","0","1","1","0","0","0","0","0","0","1","0","1"},
 {"0","0","0","1","0","0","0","1","0","1","0","1","0","1","0","1","0","1","0","1"},
 {"0","0","0","1","0","1","0","0","1","1","0","1","0","1","1","0","1","1","1","0"},
 {"0","0","0","0","1","0","0","1","1","0","0","0","0","1","0","0","0","1","0","1"},
 {"0","0","1","0","0","1","0","0","0","0","0","1","0","0","1","0","0","0","1","0"},
 {"1","0","0","1","0","0","0","0","0","0","0","1","0","0","1","0","1","0","1","0"},
 {"0","1","0","0","0","1","0","1","0","1","1","0","1","1","1","0","1","1","0","0"},
 {"1","1","0","1","0","0","0","0","1","0","0","0","0","0","0","1","0","0","0","1"},
 {"0","1","0","0","1","1","1","0","0","0","1","1","1","1","1","0","1","0","0","0"},
 {"0","0","1","1","1","0","0","0","1","1","0","0","0","1","0","1","0","0","0","0"},
 {"1","0","0","1","0","1","0","0","0","0","1","0","0","0","1","0","1","0","1","1"},
 {"1","0","1","0","0","0","0","0","0","1","0","0","0","1","0","1","0","0","0","0"},
 {"0","1","1","0","0","0","1","1","1","0","1","0","1","0","1","1","1","1","0","0"},
 {"0","1","0","0","0","0","1","1","0","0","1","0","1","0","0","1","0","0","1","1"},
 {"0","0","0","0","0","0","1","1","1","1","0","1","0","0","0","1","1","0","0","0"}}
*/
