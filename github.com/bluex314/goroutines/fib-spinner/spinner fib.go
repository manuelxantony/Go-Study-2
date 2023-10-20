package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	sch := make(chan int)
// 	go spinner(100*time.Millisecond, sch)
// 	n := 45
// 	ch := make(chan int)
// 	go fibN(n, ch)
// 	result := <-ch
// 	close(ch)
// 	close(sch)

// 	fmt.Printf("\rFibonaci at %d = %d", n, result)
// }

// func spinner(delay time.Duration, sch chan int) {
// 	for {
// 		for _, r := range `-\|/` {
// 			fmt.Printf("\r%c", r)
// 			time.Sleep(delay)
// 		}
// 	}
// }

// func fibN(nf int, ch chan int) {
// 	fn := fib(nf)
// 	ch <- fn
// }

// func fib(nf int) int {
// 	if nf < 2 {
// 		return nf
// 	}
// 	return fib(nf-1) + fib(nf-2)
// }
