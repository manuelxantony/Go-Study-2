package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

func main() {
	readFromFile()
	k := readFromConsole()[0]
	ls := readFromConsole()[0]
	arr := readFromConsole()
	fmt.Println(findTheSequence(k, ls, arr))
}

func findTheSequence(k, ls int, arr []int) int {
	lnSeq := []int{}
	j := 1
	i := 0
	count := 1
	lnSeq = append(lnSeq, arr[i])
	for _ = range arr {
		if arr[j]-arr[i] >= k {
			lnSeq = append(lnSeq, arr[j])
			count++
			i = j
		}
		j++
		if j == ls {
			break
		}
	}
	//fmt.Println(lnSeq)
	return count
}

func readFromFile() {
	f, _ := os.Open("data.txt")
	in = bufio.NewReader(f)
}

// read line by line
func readFromConsole() []int {
	line, _ := in.ReadString('\n')
	return stringToIntArr(filterData(line))
}

func filterData(line string) string {
	line = strings.Replace(line, "\n", "", -1)
	line = strings.ReplaceAll(line, "\r", "")
	return line
}

func stringToIntArr(line string) (sTiA []int) {
	sTiA = []int{}
	for _, v := range strings.Fields(line) {
		csToI, _ := strconv.Atoi(v)
		sTiA = append(sTiA, csToI)
	}
	return
}
