package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := []int{}
	sa := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Println(sa)
	s = append(s, 1, 2, 3)
	sa = (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Println(sa)
	s = append(s, 5, 6, 7)
	sa = (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Println(sa)

	// // sc is same as s
	// sc := s
	// fmt.Println(sc)

	// e := []int{}
	// she := (*reflect.SliceHeader)(unsafe.Pointer(&e))
	// fmt.Printf("sh of e before append: %+v\n", she)

	// e = append(e, 1, 2, 3)
	// she = (*reflect.SliceHeader)(unsafe.Pointer(&e))
	// fmt.Printf("sh of e: %+v\n", she)
	// m0 := make([]int, 0)
	// shm0 := (*reflect.SliceHeader)(unsafe.Pointer(&m0))
	// fmt.Printf("sh of m0: %+v\n", shm0)
	// var n []int
	// fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&n)))
	// n = append(n, 0)
	// fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&n)))

}

// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type Person interface {
// 	Age() int
// 	Name() string
// 	Place() string
// }

// type Empty interface{}

// type programmer struct {
// 	age    int
// 	name   string
// 	gender string
// }

// type athlete struct {
// 	age    int
// 	name   string
// 	gender string
// }

// func (p programmer) Age() int {
// 	return p.age
// }

// func (p programmer) Name() string {
// 	return p.name
// }

// func (p programmer) Place() string {
// 	return "A Place"
// }

// // scope out of interface
// func (p programmer) Gender() string {
// 	return ""
// }

// func (p athlete) Age() int {
// 	return p.age
// }

// func (p athlete) Name() string {
// 	return p.name
// }

// func (p athlete) Place() string {
// 	return "A Place"
// }

// // scope out of interface
// func (p athlete) Gender() string {
// 	return ""
// }

// func NewProgrammer(age int, name string) programmer {
// 	return programmer{age: age, name: name}
// }

// func NewAthlete(age int, name string) athlete {
// 	return athlete{age: age, name: name}
// }

// func getAllPeople(p, a interface{}) {
// 	//???????????

// }

// func main() {
// 	var p Person = NewProgrammer(30, "Fahad")
// 	var a Person = NewAthlete(20, "Mich")
// 	fmt.Println(p.Age())
// 	fmt.Println(a.Name())
// 	fmt.Println("-------------------")
// 	var e Empty = NewProgrammer(1, "opoppo")
// 	fmt.Println(e)
// 	fmt.Println(reflect.TypeOf(e))
// 	var s programmer = NewProgrammer(23, "dfdfdf")
// 	fmt.Println(reflect.TypeOf(s))
// }

// import "fmt"

// type Movie struct {
// 	id   string
// 	name string
// }

// var movie []Movie

// func main() {

// 	movie = append(movie, Movie{id: "1", name: "one"})
// 	fmt.Println(movie[0])
// 	for i, it := range movie {
// 		movie[i].id = "2"
// 		it.name = "two"
// 	}
// 	fmt.Println(movie[0])
// 	// 	m := make([][]int, 5)
// 	// 	x := []int{0, 0, 0, 0, 0}
// 	// 	for i := range m {
// 	// 		m[i] = x
// 	// 	}

// 	// 	n := 0
// 	// 	for i := 0; i < 5; i++ {
// 	// 		for j := 0; j < 5; j++ {
// 	// 			m[i][j] = n
// 	// 			n++
// 	// 		}
// 	// 	}
// 	// 	r := make([][]int, 5)
// 	// 	// 2d matrix
// 	// 	for i := range r {
// 	// 		r[i] = make([]int, 5)
// 	// 	}
// 	// 	fmt.Println(m)
// 	// 	fmt.Println("=====================")
// 	// 	for _ = range []int{1, 2, 3} {
// 	// 		m = modi(m, r)
// 	// 		fmt.Println("m from main:", m)
// 	// 		fmt.Println("-------------------")
// 	// 	}

// 	// }

// 	// func modi(m, r [][]int) [][]int {
// 	// 	for i := range m {
// 	// 		for j := range m {
// 	// 			r[i][j] = m[i][j] + 3
// 	// 		}
// 	// 	}
// 	// 	fmt.Println("m from modi:", m)
// 	// 	mr := make([][]int, 5)
// 	// 	for i := range mr {
// 	// 		mr[i] = make([]int, 5)
// 	// 		copy(mr[i], r[i])
// 	// 	}

// 	// 	return mr
// }
