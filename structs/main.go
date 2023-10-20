package main

import "fmt"

type personInfo struct {
	job    string
	height float32
}

type person struct {
	firstName string
	lastName  string
	pInfo     personInfo
}

type in int

func main() {

	p := person{
		firstName: "Michael",
		lastName:  "Jordan",
		pInfo: personInfo{
			job:    "Former NBA player",
			height: 6.7,
		},
	}
	(&p).updateName("Mike")
	//p.updateName("Mike")
	p.print()

}

func (p *person) updateName(newName string) {
	(p).firstName = newName
	fmt.Println(p)
	fmt.Println(*p)
}

func (p person) print() {
	fmt.Println(p)
}
