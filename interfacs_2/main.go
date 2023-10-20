package main

import "fmt"

type Currency float32

type Stringer interface {
	String() string
	Just()
}

func main() {
	var c = new(Currency) // creats a pointer
	*c = 1.1
	fmt.Println(c.String()) //explicit
	fmt.Printf("%v\n", c)   //not implicit

	var c1 Currency
	c1 = 2.1
	fmt.Println(c1.String()) //explicit
	fmt.Printf("%v", c1)     //implicit (Here the String() will be called from fmt, the following String implements a pointer Currency)

	// v := 1.444
	// fmt.Println(v)

	// s := Stringer(c)
	// fmt.Printf("%v, %T", s, s)

}

func (c *Currency) String() string {
	return fmt.Sprintf("$%0.2f", float32(*c))
}

func (c Currency) Just() {}
