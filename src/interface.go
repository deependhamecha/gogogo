package main

import "fmt"


type NormalUser struct {
	name string
}

func (n NormalUser) dude() {
	fmt.Println("Dude")
}



type User interface {
	dude()
}

func printName(z User) {
	fmt.Println(z)
}


func main() {
	u := NormalUser{name: "Deepen"}
	printName(u)
}