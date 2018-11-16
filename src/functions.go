package main

import "fmt"

func main() {
	fmt.Println(name("Deepen", "Dhamecha"))

	fmt.Println(new_thing_in_go("Deepen"))

	fmt.Println(multiple_returns("Deepen", "Dhamecha"))
}


// Last parameter must be defined with a type
func name(fname, lname string) string {
	return fmt.Sprint(fname, lname, "Dude", 1, true)
}

// You can define a variable to return and just write return in the last line without variable
// It will throw s for you.
// But its useless as you can also return anything besides even if you specify it
func new_thing_in_go(name string) (s string) {
	s = name
	return
}

// Multiple returns

func multiple_returns(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}