package main

import (
	"errors"
	"fmt"
)

func bagi (value1 int, value2 int) (int, error) {
	if value1 * value2 == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		return value1 / value2, nil
	}
} 

func main() {
	var errContoh error = errors.New("Error dude")

	fmt.Println(errContoh.Error())

	hasil, err := bagi(10, 0)
	
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println(err.Error())
	}
}