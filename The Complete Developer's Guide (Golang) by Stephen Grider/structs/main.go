package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // Use `contactInfo` as a field name that is the same as the type name
}

func main() {
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	// fmt.Println(alex)       // output: {Alex Anderson}
	// fmt.Printf("%+v", alex) // output: {firstName:Alex lastName:Anderson}%

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@mail.com",
			zipCode: 94000,
		},
	}

	// fmt.Printf("%+v", jim) // output: {firstName:Jim lastName:Party contact:{email:jim@mailcom zipCode:94000}}
	jim.updateName("Jimmy")
	jim.print()

	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")
	// jimPointer.print()
}

// with `func (p person) updateName(newFirstName string)` we are passing a copy of the person struct
// so the updateName function will not update the original person struct.
// Change the receiver to a pointer to the person struct to update the original person struct.
// `func (p *person) updateName(newFirstName string)` now we should see the updated name
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
