package main

import "fmt"

import "sort"

type Dude struct {
	name string
}

type Interface1 interface {
	dude() string
	motherfucker() string
}


func (d Dude)dude() string {
	return d.name
}

func (d Dude) motherfucker() string {
	return d.name+" motherfucker"
}

// For Interface to implement, concrete classes needs to implement all methods
func interfaceMethod(i Interface1) string {
	return i.dude()
}


func main() {
	dude := Dude{
		name: "Deepen",
	}


	fmt.Println(interfaceMethod(dude))

	a := []string{"Zenila", "Dude", "Adiel"}
	// sort.Strings(a)
	// or
	sort.StringSlice(a).Sort()
	// either works
	// sort.Sort(sort.StringSlice(a))
	fmt.Println(a)
}