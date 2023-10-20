package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
	// Write your code here
	sortedArr := sortArr(arr)
	fmt.Printf("%d %d", sumArr(sortedArr[:4]), sumArr(sortedArr[1:]))
}

func sumArr(arr []int32) (s int64) {
	for _, v := range arr {
		s += int64(v)
	}
	return
}

func sortArr(arr []int32) []int32 {
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}

	return arr
}

func main() {
	//reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	f, err := os.Open("./data.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f)
	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
