package main

import "fmt"

func main () {
	name := "Dayat"

	switch name {
	case "Dayat":
		fmt.Println("Hello, Dayat")
	case "Nur":
		fmt.Println("Hello, Nur")
	default:
		fmt.Println("Hellow there!")
	}
}