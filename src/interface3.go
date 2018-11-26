package main

import (
	"fmt"
)

type animal interface{}

type dog struct {
	animal
	Name string
}

type cat struct {
	animal
	Name string
}

func main() {
	fmt.Println("Hello, playground")

	var dude = 3

	fmt.Printf("dude: %T", dude)
	fmt.Println()

	//	a := new([]animal)
	a := []animal{dog{Name: "Wow"}, cat{Name: "Meow"}}

	fmt.Println(a)
	
	for index, item:= range a {
		fmt.Println(index, " - ", item)
	}
}
