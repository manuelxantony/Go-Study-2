package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFromFile()
	s := readFromConsole()[0]
	m := make([][]int, s)
	i := 0
	l := s
	for s > 0 {
		m[i] = readFromConsole()
		i++
		s--
	}
	r := make([][]int, l)
	// 2d matrix
	for i := range r {
		r[i] = make([]int, l)
	}
	rv := readRoationVales()
	for _, v := range rv {
		if v == "L" {
			m = rotate(m, r, l, true)
		}
		if v == "R" {
			m = rotate(m, r, l, false)

		}

	}
	printMatrixValues(r, l)
}

func printMatrixValues(m [][]int, s int) {
	for i := 0; i < s; i++ {
		for j := 0; j < s; j++ {
			fmt.Print(m[i][j], " ")
		}
		fmt.Println("")
	}
}

func readRoationVales() []string {
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "LR", "")
	line = strings.ReplaceAll(line, "RL", "")
	return strings.Split(line, "")

}

func rotate(m [][]int, r [][]int, l int, left bool) [][]int {
	si := l
	l -= 1

	for i := 0; i < si; i++ {
		for j := 0; j < si; j++ {
			// left rotation
			if left {
				r[i][j] = m[j][l]
			} else {
				// right rotation
				r[i][j] = m[l][i]
				l -= 1
			}
		}
		if left {
			l -= 1
		} else {
			l = si - 1
		}

	}
	mr := make([][]int, si)
	for i := range mr {
		mr[i] = make([]int, si)
		copy(mr[i], r[i])
	}

	return mr
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
