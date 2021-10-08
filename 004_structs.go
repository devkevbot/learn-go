package main

import "fmt"

type address struct {
	street      string
	houseNumber string
	postalCode  string
	city        string
	state       string
	country     string
}

func (a address) String() string {
	return fmt.Sprintf("%s %s\n%s\n%s, %s\n%s",
		a.houseNumber, a.street, a.postalCode, a.city, a.state, a.country)
}

/* Pointer receivers */
func (a *address) moveToCanada() {
	a.street = "Louis Lane"
	a.houseNumber = "456"
	a.postalCode = "V3K 8M6"
	a.city = "Vancouver"
	a.state = "British Columbia"
	a.country = "Canada"
}

func _004Structs() {
	_address := address{
		street:      "Candy Cane Lane",
		houseNumber: "123",
		postalCode:  "123 ABC",
		city:        "New York",
		state:       "New York",
		country:     "USA",
	}

	fmt.Printf("Address is: \n%s\n", _address)

	_address.moveToCanada()
	fmt.Printf("After moving to Canada, address is: \n%s\n", _address)
}
