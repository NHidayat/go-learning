package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Maried () {
	man.Name = "Mr " + man.Name
}

func main () {
	dayat := Man{"Dayat"}
	dayat.Maried()


	fmt.Println(dayat.Name)
}