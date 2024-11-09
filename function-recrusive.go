package main

import "fmt"

func factorialLoop (number int) int {
	result := 1

	for i := number; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialRecrusive (number int) int {
	if number == 1 {
		return 1
	} else {
		return number * factorialRecrusive(number -1)
	}
}

func main () {
	loop := factorialLoop(5)
	fmt.Println(loop)

	loop2 := factorialRecrusive(5)
	fmt.Println(loop2)
}