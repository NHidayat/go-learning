package main

import "fmt"

func sumAll (numbers ...int) int{
	result := 0

	for _, number := range numbers {
		result += number
	}

	return result
}

func main () {
	value1 := sumAll(1,2,3,23,10,20)
	fmt.Println(value1)

	// menggunakan slice
	slice := []int{29,101,12,2}
	value2 := sumAll(slice...)
	fmt.Println(value2)

}