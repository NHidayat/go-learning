package main

import "fmt"

func main () {

	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Loop ", counter)
	// 	counter++
	// }

	// for count := 1; count <= 10; count++ {
	// 	fmt.Println("count ", count)
	// }

	// names := []string{"Muhammad", "Nur", "Hidayat"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for i, name := range names {
	// 	fmt.Println("Index ", i, " ", name)
	// }

	person := make(map[string]string)
	person["name"] = "dayat"
	person["age"] = "22"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

}