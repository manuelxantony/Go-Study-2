package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	//fmt.Println(scanner.Text())
	//cases1, _ := strconv.ParseInt(scanner.Text(), 0, 64)
	cases, _ := strconv.Atoi(scanner.Text())
	fmt.Println("cases: ", cases)
	for i := 0; i < cases; i++ {
		var IA []int
		var IR []int
		fmt.Println("ar:", IA, IR)

		scanner.Scan()
		length_of_seq_A, _ := strconv.Atoi(scanner.Text())
		fmt.Println("Seq Length:", length_of_seq_A)

		scanner.Scan()
		seq_string_A := strings.Split(scanner.Text(), " ")

		//creting Slice of int for A
		for j := 0; j < length_of_seq_A; j++ {
			val, _ := strconv.Atoi(seq_string_A[j])
			IA = append(IA, val)
		}
		fmt.Println("A int Strign:", IA)

		scanner.Scan()
		length_of_seq_R, _ := strconv.Atoi(scanner.Text())
		fmt.Println("Seq Length:", length_of_seq_R)

		scanner.Scan()
		seq_string_R := strings.Split(scanner.Text(), " ")
		for j := 0; j < length_of_seq_R; j++ {
			val, _ := strconv.Atoi(seq_string_R[j])
			IR = append(IR, val)
		}
		fmt.Println("R int String", IR)
		fmt.Println(IA, IR)

	}

}
