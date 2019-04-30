package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	charles := person{
		firstName: "Charles",
		lastName:  "Qiu",
		contactInfo: contactInfo{
			email:   "xxx@mail",
			zipCode: 221,
		},
	}

	charles.setFirstName("CH")
	charles.print()

	name := "XXX"
	fmt.Println(*&name)
}

func (p *person) setFirstName(firstName string) {
	// fmt.Printf("%p\n", p)
	p.firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
