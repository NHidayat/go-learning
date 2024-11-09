package main

import "fmt"

func main () {

	value := 100

	if value < 30 {
		fmt.Println("D")
	} else if value < 50 {
		fmt.Println("C")
	} else {
		fmt.Println("A")
	}

}