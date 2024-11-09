package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main () {
	address1 := Address{"Jakarta", "DKI Jakarta", "Indonesua"}
	address2 := &address1
	address3 := &address1

	address2.City = "Makassar"

	*address2 = Address{"Bali", "Bali", "Indonesia"} // Merubah semua variable yg menggunakan ref yg sama (address1)

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	address4 := new(Address) // Address baru
	address4.City = "Papua"
	fmt.Println(address4)

	address5 := Address{
		City: "Bandung",
		Province: "West Java",
		Country: "US",
	}

	changeCountryToIndonesia(&address5)
	fmt.Println(address5)

}