package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i:=0;i<3;i++ {
		fmt.Println("i: "+strconv.Itoa(i))
	}

	a:=4;
	for a<7 {
		fmt.Println(a)
		a++
	}

	for {
		if(a==10) {
			a++
			continue
		}

		fmt.Println(a)

		a++

		if(a>12) {
			break
		}
	}

	// In Go Switch doesnt go to another block if not specified break
	// But if you want then use fallthrough
	switch "Dude1" {
		case "Dude1":
			fmt.Println("Dude1")
			fallthrough
		case "Dude":
			fmt.Println("Dude")
		case "Dude2":
			fmt.Println("Dude2")
		default:
			fmt.Println("Default")
	}

	// If you do not specify switch identifier then it will test condition in case
	d := "Dude"
	switch {
		case d=="Dude1":
			fmt.Println("Dude1")
			fallthrough
		case d=="Dude":
			fmt.Println("Dude")
		case d=="Dude2":
			fmt.Println("Dude2")
		default:
			fmt.Println("Default")
	}

	// Lecture 260 5.00

	// Check type
	// var d interface
	// switch d.(type) {
	// 	case int:
	// 		fmt.Println("Dude1")
	// 		fallthrough
	// 	case string:
	// 		fmt.Println("Dude")
	// 	default:
	// 		fmt.Println("Default")

	// }

	if true {
		fmt.Println(true);
	}

	// You can also initialize in conditional statements
	// mother variable is scoped
	if mother:=5;true {
		fmt.Println(mother)
	}

}