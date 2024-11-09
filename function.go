package main

import "fmt"



func sayHeloo (name string) {
	fmt.Println("Hello ", name)
}

func sum (value1 int, value2 int) int {
	return value1 + value2
}

func getFullName () (string, string) {
	return "Nur", "Hidayat"
}


func main () {
	sayHeloo("dayat")
	fmt.Println(sum(10, 29))
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}