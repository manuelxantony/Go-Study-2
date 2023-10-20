package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func mainn() {
	if len(os.Args) >= 2 && os.Args[1] == "FF" {
		readFromFile()
	}
	//exclude while submitting
	//readFromFile()
	//test cases
	tc := readFromConsole()
	for i := 0; i < tc[0]; i++ {
		logic()
	}

}

func logic() {
	c := readFromConsole()
	n, m := c[0], c[1]
	rArr := make([]int, n)
	for i := 0; i < m; i++ {
		r := readFromConsole()
		if r[0] >= n || r[1] >= n {
			continue
		}

		// == conditon have to check, now if its == then adding 2
		if r[0] <= r[1] {
			for i := r[0]; i <= r[1]; i++ {
				rArr[i] += 1
			}
		}
		if r[0] > r[1] {

			for k := r[0]; k < n; k++ {
				rArr[k] += 1
			}
			for j := 0; j <= r[1]; j++ {
				rArr[j] += 1
			}
		}
	}
	count1 := 0
	for _, v := range rArr {
		if v == 1 {
			count1++
		}
	}
	if count1 == 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

var in = bufio.NewReader(os.Stdin)

func readFromConsole() []int {
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "\n", "")
	line = strings.ReplaceAll(line, "\r", "")
	return stringToIntArray(line)
}

func stringToIntArray(l string) []int {
	sArr := stringToArray(l)
	len := len(sArr)
	arr := make([]int, len)
	for i := 0; i < len; i++ {
		arr[i], _ = strconv.Atoi(sArr[i])
	}
	return arr
}

func stringToArray(l string) []string {
	return strings.Split(l, " ")
}

func readFromFile() {
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	in = bufio.NewReader(f)
}
