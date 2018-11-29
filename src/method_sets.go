package main

import "fmt"

type User struct {
	name string
}

func (u User) getName() string {
	return u.name
}

func (u *User) getNameViaPointer() string {
	return (*u).name
}

type Person interface {
	getName() string
}

func data(s Person) {
	fmt.Println("via interface:",s.getName())
}

// func info(s *Person) {
// 	fmt.Println((*s).getNameViaPointer())
// }

func main() {
	user := User{"Deepen"}

	data(user)
	data(&user)

	fmt.Println("using concrete object:", user.getName())
	fmt.Println("using address:", (&user).getName())

	fmt.Println("using pointer:", (&user).getNameViaPointer())
	// fmt.Println((*user).info())

	// p := &user

	// fmt.Println((*p).getName())

	// q := new(User)
	// q.name = "Dhamecha"
	// fmt.Println(q.getName())
}
