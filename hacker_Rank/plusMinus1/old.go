package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'findSubstring' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */

const vowels string = "aeiou"

func findSubstring(s string, k int32) string {
	// Write your code here
	var vCount int
	var bigVowel string
	if numOfVowels(s) == 0 {
		return "Not found!"
	}

	// length of string
	lS := int32(len(s))
	for i := int32(0); i < lS-k+1; i++ {
		subStringOfKLen := s[i : i+k]
		nV := numOfVowels(subStringOfKLen)
		if vCount < nV {
			vCount = nV
			bigVowel = subStringOfKLen
		}
		if vCount == nV {
			bigVowel = bigVowelAmongCommon(bigVowel, subStringOfKLen, int(k))
		}
	}

	return bigVowel
}

func bigVowelAmongCommon(bigVowel, subStringOfKLen string, len int) string {
	for i := 0; i < len; i++ {
		if isAVowel(string(bigVowel[i])) {
			return bigVowel
		}
		if isAVowel(string(subStringOfKLen[i])) {
			return subStringOfKLen
		}

	}
	// never reach
	return ""
}

func numOfVowels(s string) int {
	var vCount int
	for _, v := range s {
		if isAVowel(string(v)) {
			vCount++
		}
	}
	return vCount
}

func isAVowel(c string) bool {
	for _, v := range vowels {
		if c == string(v) {
			return true
		}
	}
	return false
}

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	in := bufio.NewReader(f)
	reader := bufio.NewReaderSize(in, 16*1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := findSubstring(s, k)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
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
