// Struct adalah template atau prototype data

package main

import "fmt"

type Customer struct {
	// name, address string
	Name string
	Age int
	Address string
}

func main () {

	var cust Customer
	cust.Name = "Dayat"
	cust.Age = 23
	cust.Address = "Tomang Raya"

	cust2 := Customer {
		Name: "John",
		Age: 30,
		Address: "New York",
	}

	cust3 := Customer {"Jenny", 28, "LA"}

	fmt.Println(cust)
	fmt.Println(cust2)
	fmt.Println(cust3)
}