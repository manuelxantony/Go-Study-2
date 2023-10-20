package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func lonelyinteger(arr []int) int {
	var x int
	for _, val := range arr {
		x ^= val
	}

	return x
}

func main() {
	//reader
	//from file
	fReader := fileReader()
	reader := bufio.NewReader(fReader)
	//writer
	writer := bufio.NewWriter(os.Stdout)
	// read
	// array size from line one
	_, err := strconv.Atoi(readLine(reader))
	errOut(err)
	arr := make([]int, 0)
	// array data from line two
	arrData := strings.Split(readLine(reader), " ")
	for _, v := range arrData {
		vInt, err := strconv.Atoi(v)
		errOut(err)
		arr = append(arr, vInt)
	}

	result := lonelyinteger(arr)

	fmt.Fprintf(writer, "%d\n", result)
	writer.Flush()
}

func readLine(r *bufio.Reader) string {
	data, _, err := r.ReadLine()
	errOut(err)
	return strings.TrimRight(string(data), "\r\n")

}

func errOut(err error) {
	if err != nil {
		panic(err)
	}
}

func fileReader() *os.File {
	f, err := os.Open("./data.txt")
	errOut(err)
	return f
}
