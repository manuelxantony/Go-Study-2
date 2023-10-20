package main

import (
	"bytes"
	"fmt"

	"io/ioutil"

	"log"

	"strconv"

	"strings"

	"os"
)

func main() {

	// reading all file

	file, err := ioutil.ReadAll(os.Stdin)

	// if error occurred meanwhile reading, loging Fatal

	if err != nil {

		log.Fatal(err)

	}

	// creating buffer of read file

	buf := bytes.NewBuffer(file)

	line1, _ := buf.ReadString('\n')

	cases, _ := strconv.Atoi(strings.TrimSpace(line1))
	if cases < 1 || cases > 10 {

		return

	}

	for ; cases > 0; cases-- {

		// reading length n

		line2, _ := buf.ReadString('\n')

		// reading line of A

		line3, _ := buf.ReadString('\n')

		// reading line of R, but used B in programming

		line4, _ := buf.ReadString('\n')

		n, _ := strconv.Atoi(strings.TrimSpace(line2))

		a, b := []int{}, []int{}

		// reading line and triming the line value for rid of left and right whitespace & remove new line from string value and converting array of string values

		A := strings.Split(strings.TrimSpace(line3), " ")

		B := strings.Split(strings.TrimSpace(line4), " ")

		for i := 0; i < n; i++ {

			// converting the A indexed value to number

			num1, _ := strconv.Atoi(A[i])

			num2, _ := strconv.Atoi(B[i])

			// checking and appending the slice by checking Question given constraint for both "a" and "b" slice

			if num1 >= 0 || num1 < n {

				a = append(a, num1)

			}

			if num2 >= 0 || num2 < n {

				b = append(b, num2)

			}

		}

		// initializing map and "set" named slice, here map will insert key and value and simply check, it had been already inserted or not

		// but at the same time "set" slice has been used to manage and keep the sequence of number insert consistent, otherwise map reading is not sequential, as you know, and numbers will fluctuate.

		map1 := map[int32]uint8{}

		set := []int32{}

		for i := 0; i < n; i++ {

			// here suppose for i index is zero, so it will add first beginning item of "a" index and last item of "b" index and thereafter subtract after "n" length

			val := n - (a[i] + b[n-1-i])

			_, exists := map1[int32(val)]

			// will only took the number once in set, to keep distinct and avoid duplicasy.

			if !exists {

				map1[int32(val)] = 1

				set = append(set, int32(val))

			}

		}

		// check if "set" length is equal to "n" and "set" first item is not equal to 1 and also "set" last item is not equal to "n"

		if len(set) == n && set[0] != 1 && set[len(set)-1] != int32(n) {

			for i := 0; i < n; i++ {

				// converting number to explicit string to printing for better performance, at least it will buy some more seconds or milliseconds compared printing directly numbers

				fmt.Print(strconv.Itoa(int(set[i])), " ")

			}

			fmt.Println()

		} else {

			fmt.Println("-1")

		}

	}

}
