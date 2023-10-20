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
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here
	generalTimeArr := strings.Split(s, ":")
	// am pm
	amPm := AMOrPM(s)
	if amPm == "AM" {
		if generalTimeArr[0] == "12" {
			generalTimeArr[0] = "00"
		}

	}

	if amPm == "PM" {

		tInt := stringToInt(generalTimeArr[0])
		if tInt < 12 {
			tInt += 12
			generalTimeArr[0] = intToString(tInt)
		}
	}
	generalTimeArr[2] = strings.Replace(generalTimeArr[2], amPm, "", 1)
	return joinStringArr(generalTimeArr)
}

func stringToInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func intToString(i int) string {
	return strconv.Itoa(i)
}

func joinStringArr(s []string) string {
	return strings.Join(s, ":")
}

func AMOrPM(t string) string {
	if strings.Contains(t, "AM") {
		return "AM"
	}
	if strings.Contains(t, "PM") {
		return "PM"
	}
	return "not found AM or PM"
}

func main() {
	//reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	f, err := os.Open("./data.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f)

	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

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
