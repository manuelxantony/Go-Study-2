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
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	// Write your code here
	dataArr := make([]int32, 3)
	for _, v := range arr {
		if v > 0 {
			dataArr[0]++
		}
		if v < 0 {
			dataArr[1]++
		}
		if v == 0 {
			dataArr[2]++
		}
	}

	lenArr := len(arr)
	for _, v := range dataArr {
		fmt.Printf("%.6f\n", float32(v)/float32(lenArr))
	}

}

func main() {

	f, err := os.Open("./data.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f)

	//reader := bufio.NewReaderSize(br, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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
