package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var num = 0
// var mu sync.Mutex

// func add3(n int) {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	num = n + 3
// 	fmt.Println(num)
// }

// func main() {
// 	for i := 0; i < 9; i++ {
// 		go add3(i)
// 	}
// 	time.Sleep(time.Second)
// }
