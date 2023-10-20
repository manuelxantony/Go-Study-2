// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func lonelyinteger(arr []int) int {
// 	var index []int
// 	for _, v := range arr {
// 		if !contains(index, v) {
// 			// add to index
// 			index = append(index, v)
// 		} else {
// 			// remove it from index
// 			index = remove(index, v)
// 		}
// 	}

// 	return index[0]
// }

// func remove(arr []int, v int) []int {
// 	arrLen := len(arr)
// 	var index int
// 	for i := 0; i < arrLen; i++ {
// 		if v == arr[i] {
// 			index = i
// 			break
// 		}
// 	}
// 	//assign
// 	arr[index] = arr[arrLen-1]
// 	//nullify / nil value
// 	arr[arrLen-1] = 0
// 	// truncate
// 	arr = arr[:arrLen-1]
// 	return arr
// }

// func contains(arr []int, v int) bool {
// 	if len(arr) == 0 {
// 		return false
// 	}
// 	for _, val := range arr {
// 		if val == v {
// 			return true
// 		}
// 	}
// 	return false
// }

// func main() {
// 	//reader
// 	//from file
// 	fReader := fileReader()
// 	reader := bufio.NewReader(fReader)
// 	//writer
// 	writer := bufio.NewWriter(os.Stdout)
// 	// read
// 	// array size from line one
// 	_, err := strconv.Atoi(readLine(reader))
// 	errOut(err)
// 	arr := make([]int, 0)
// 	// array data from line two
// 	arrData := strings.Split(readLine(reader), " ")
// 	for _, v := range arrData {
// 		vInt, err := strconv.Atoi(v)
// 		errOut(err)
// 		arr = append(arr, vInt)
// 	}

// 	result := lonelyinteger(arr)

// 	fmt.Fprintf(writer, "%d\n", result)
// 	writer.Flush()
// }

// func readLine(r *bufio.Reader) string {
// 	data, _, err := r.ReadLine()
// 	errOut(err)
// 	return strings.TrimRight(string(data), "\r\n")

// }

// func errOut(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func fileReader() *os.File {
// 	f, err := os.Open("./data.txt")
// 	errOut(err)
// 	return f
// }
