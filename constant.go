package main

func main () {
	// Cara 1
	const firstName string = "Dayat"
	const lastName string = "Nur"

	// Cara 2
	const (
		age int = 24
		active string = "y"
	)

	println(firstName)
	println(lastName)
	println(age)
	println(active)
}