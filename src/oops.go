package main

import "fmt"

/**
*	Encapsulation
*/
//------------------1---------------
type foo int



//------------------2----------------
type Person struct {
	firstName string
	lastName string
}

type Dude struct {
	Person
	age int
}



//------------------3---------------
type UserDetail struct {
	person Person
	email string
}



//-------------Inheritance-------------
//-------------Polymorphism------------

func main() {


	//----------------1-----------------
	var myType foo = 5

	fmt.Println("myType",myType)



//---------------2------------------
	dude := Dude{
		Person: Person{
			firstName: "Deepen",
			lastName: "Dhamecha",
		},
		age: 27,
	}

	fmt.Println("Dude:", dude)

//-------------3---------------------
	ud := UserDetail{
		person: Person{
			firstName: "Harry",
			lastName: "Potter",
		},
		email: "deependhamecha@gmail.com",
	}

	fmt.Println("UserDetail:", ud)


}
