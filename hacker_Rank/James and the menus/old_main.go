package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main1() {
// 	values := readTextFromConsole()
// 	//values := "3 4\n1 2 1 10\n3 2 3 4\n1 3 3 2"
// 	all_input := strings.Split(values, "\n")
// 	// line 1
// 	num_of_menues, _ := strconv.Atoi(strings.Split(all_input[0], " ")[0])
// 	num_of_items, _ := strconv.Atoi(strings.Split(all_input[0], " ")[1])

// 	menues := [][]int{}
// 	for i := 0; i < num_of_menues; i++ {
// 		menu_string := strings.Split(all_input[i+1], " ")
// 		menu := sliceToIntString(menu_string)
// 		menues = append(menues, menu)

// 	}
// 	//fmt.Println(menues)
// 	answer := logic(menues, num_of_menues, num_of_items)
// 	fmt.Println(answer)
// }

// func logic(menues [][]int, r int, c int) int {
// 	ansH := make([][]int, 0)
// 	f := make([]int, r)
// 	for i := 0; i < c; i++ {
// 		res := menues[0][i]
// 		pos := []int{}
// 		pos = append(pos, 0)
// 		a := make([]int, r)
// 		for j := 1; j < r; j++ {
// 			if res < menues[j][i] {
// 				pos = []int{}
// 			}
// 			if res <= menues[j][i] {
// 				res = menues[j][i]
// 				pos = append(pos, j)
// 			}
// 		}
// 		for _, p := range pos {
// 			a[p] += 1
// 		}
// 		ansH = append(ansH, a)
// 	}
// 	for _, item := range ansH {
// 		for it, ite := range item {
// 			f[it] += ite
// 		}
// 	}
// 	// in f. find the higest
// 	ri := f[0]
// 	hp := make([]int, r)
// 	for i, fi := range f {
// 		if ri < fi {
// 			hp = make([]int, r)
// 		}
// 		if ri <= fi {
// 			ri = fi
// 			hp[i] = 1
// 		}

// 	}
// 	//fmt.Println(hp)

// 	// finding the average in the higest menues
// 	avg := 0

// 	for ixix, hpp := range hp {
// 		t := 0
// 		if hpp != 0 {
// 			//fmt.Println(menues[ixix])
// 			for _, e := range menues[ixix] {
// 				t += e
// 			}
// 			if avg < t/c {
// 				avg = t / c
// 			}

// 			//fmt.Println(avg)
// 		}

// 	}
// 	return avg
// }

// func sliceToIntString(menu_string []string) []int {
// 	var arr []int
// 	for _, val := range menu_string {
// 		i_val, _ := strconv.Atoi(val)
// 		arr = append(arr, i_val)
// 	}
// 	return arr
// }

// func readTextFromConsole() string {
// 	readText, err := ioutil.ReadAll(os.Stdin)
// 	if err != nil {
// 		log.Printf("failed to read from console %s", err)
// 	}
// 	return string(readText)
// }
