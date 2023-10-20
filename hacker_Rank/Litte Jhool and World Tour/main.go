package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputArray []string

func ProcessData(N uint64, M uint64, k uint64) string {
	mm := make(map[uint64]uint64)

	for M > 0 {
		tmp := strings.Fields(inputArray[k])
		a, _ := strconv.ParseUint(tmp[0], 10, 64)
		b, _ := strconv.ParseUint(tmp[1], 10, 64)

		test1 := false
		if a <= b {
			for i := a; i <= b; i++ {
				if _, ok := mm[i]; !ok {
					mm[i] = 1
					test1 = true
					break
				}
			}
			for i := a + N; i < b+N; i++ {
				if _, ok := mm[i]; !ok {
					mm[i] = 1
					test1 = true
					break
				}
			}
		} else {
			for i := a; i < b+N-1; i++ {
				if _, ok := mm[i]; !ok {
					mm[i] = 1
					test1 = true
					break
				}
			}
		}
		if test1 == false {
			return "NO"
		}
		M--
		k++
	}
	return "YES"
}

func main() {

	f, _ := os.Open("data.txt")
	in := bufio.NewReader(f)

	b := make([]byte, 100000*10000)
	//n, _ := os.Stdin.Read(b)
	n, _ := in.Read(b)
	inputArray = strings.Split(string(b[:n]), "\n")
	testCases, _ := strconv.Atoi(strings.ReplaceAll(inputArray[0], "\r", ""))
	k := uint64(1)
	for testCases > 0 {
		tmp := strings.Fields(strings.ReplaceAll(inputArray[k], "\r", ""))
		N, _ := strconv.ParseUint(tmp[0], 10, 64)
		M, _ := strconv.ParseUint(tmp[1], 10, 64)
		k++
		fmt.Println(ProcessData(N, M, k))
		k = k + M
		testCases--
	}
}
