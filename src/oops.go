package main

import "fmt"

/**
*	Encapsulation
*/
//------------------1---------------
type foo int



//-------------Inheritance-----------
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



//-------------Polymorphism------------

func main() {

	//----------------1-----------------
	var myType foo = 5
	m:= 2

	fmt.Println("myType: ", myType)

	// Wont work
	// fmt.Println(myType+m)
	fmt.Println(int(myType) + m)


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

	fmt.Println("------------------")
	fmt.Println("Function Called: ", dude.Person.printPersonName())


}

// Only Dude type variable can call this 
func (d Person) printPersonName() string {
	return fmt.Sprint(d.firstName+" "+d.lastName)
}