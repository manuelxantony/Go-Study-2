package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println(cfs(readFromConsole()))
}

func cfs(n1, n2 int) (count int) {
	count = 0
	if n1 == 0 || n2 == 0 {
		return
	} else {
		if n1 < n2 {
			n1, n2 = n2, n1
		}
		count++
		sq := int(math.Sqrt(float64(n1)))
		for i := 2; i <= sq; i++ {
			if n1%i == 0 {
				if n2%i == 0 {
					count++
				}
				if n2%(n1/i) == 0 {
					count++
				}
			}
		}
	}
	return
}

func readFromConsole() (int, int) {
	l, _ := in.ReadString('\n')
	line := strings.Fields(l)
	n1, _ := strconv.Atoi(line[0])
	n2, _ := strconv.Atoi(line[1])
	return n1, n2
}
