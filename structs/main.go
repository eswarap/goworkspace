package main

import "fmt"

type contactInfo struct {
	email      string
	postalcode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	kuppan := person{
		firstName: "Kuppu",
		lastName:  "Swamy",
		contact: contactInfo{
			email:      "test@test.com",
			postalcode: 560001,
		},
	}
	// Turn value into address
	// kuppanPointer := &kuppan
	kuppan.updateName("subban")
	kuppan.print()

	name := "bill"

	namePointer := &name

	fmt.Println(&namePointer)
	printPointer(namePointer)
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}

func (p *person) updateName(newFirstName string) {
	//turn address into value
	(*p).firstName = newFirstName
}
