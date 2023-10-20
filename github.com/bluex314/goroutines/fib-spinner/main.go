package main

import "fmt"

func fib(ch chan int, quite chan bool) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quite: // when a data passed to this channel it will pass
			return
		}
	}
}

func main() {
	ch := make(chan int)
	quite := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			result := <-ch
			fmt.Println(result)
		}
		// when we put data into this channel at fib, select will select the case which <-quite
		quite <- true

	}()
	fib(ch, quite)
}
