package main

import "fmt"

func main ()  {
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name string = "Dayat"
	var e = name[0]
	var eString = string(e)

	println(name)
	println(e)
	println(eString)
}