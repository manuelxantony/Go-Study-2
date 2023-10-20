/*
// Sample code to perform I/O:

fmt.Scanf("%s", &myname)            // Reading input from STDIN
fmt.Println("Hello", myname)        // Writing output to STDOUT

// Warning: Printing unwanted or ill-formatted data to output will cause the test cases to fail
*/

// Write your code here
package main

import "fmt"

func main() {
	var test_case_num int
	fmt.Scanf("%d\n", &test_case_num)
	fmt.Println("number of test", test_case_num)

	for i := 0; i < test_case_num; i++ {
		var num_of_elements int
		var A []int
		var R []int
		fmt.Scanf("%d\n", &num_of_elements)
		fmt.Println("number or elements", num_of_elements)
		for j := 0; j < num_of_elements; j++ {
			var num int
			fmt.Scanf("%d", &num)
			A = append(A, num)
		}
		fmt.Println(A)
		var x int
		fmt.Scanf("%d/n", &x)
		for j := 0; j < num_of_elements; j++ {
			var num int
			fmt.Scanf("%d", &num)
			R = append(R, num)
		}
		fmt.Scanf("%d/n", &x)
		fmt.Println(R)
	}

}
