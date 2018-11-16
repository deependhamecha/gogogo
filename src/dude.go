package main

import "fmt"


func add(a, d int) int {
	return a+d
}

func main() {
	fmt.Print("Hello")
	{
		fmt.Println(" World")
	}

	fmt.Println(add(5, 1))

	var dude = func(x,y int) {
		fmt.Println(x+y)
	}

// If you are returning something then you need to specify return type
	dude1 := func(x int,y int) int {
		fmt.Println(x+y)
		return x+y
	}

	dude2 := dude1(1,2)

	// Share closure with another variable
	dude3 := dude1


	dude(7,2)
	dude1(3,4)
	fmt.Println(dude2)

	dude1 = nil

	dude3(8,2)

	// It has shit expressions of nested functions
	// function can return a function
	// Type : func() TYPENAME
	okay := func() func() string {
		return func() string {
			return "yes"
		}
	}

	fmt.Println(okay()());


	// Blank Space, declared bt not used solved
	var ab int = 1
	var cd = 2
	
	// Multiple assignments also work in go lang
	_, _ = ab, cd

//	_ = cd
//	_ = ab

	// For Import
	// import _ "fmt"

	
	// For Functions	
	var abc func() = func() {
		fmt.Println("Hello");		
	}

	_ = abc

	const gl string = "Harry Potter"

	const (
		A = iota
		B = iota
		C = iota
	)

	fmt.Println("A+B+C : ", A+B+C)
	s := int(49)
	_ = s

	var scanit int

	fmt.Print("\n Enter a Number:")
	// fmt.Scan(&scanit)

	fmt.Println("scanit: ",scanit)

	fmt.Println("----------------")

	u:=7
	v:=u

	// w stores address of u
	var w *int = &u

	fmt.Println("u : ", u)
	fmt.Println("v : ", v)

	fmt.Println("w : ", w)

	// *w points to value of the address
	fmt.Println("*w : ", *w)

	u=8
	fmt.Println("u : ", u)
	fmt.Println("v : ", v)

	fmt.Println("----------------")

	fmt.Println("w : ", w)
	fmt.Println("*w : ", *w)

	fmt.Println("----------------")

	*w = 10
	fmt.Println("*w : ", *w)
	fmt.Println("u : ", u)

	fmt.Println("----------------")
	g:=5
	with_params(&g)
	fmt.Println(g)

	ra:='A'

	fmt.Println(ra)

	r:=`aaad
	dasfafaf
	dfd`

	fmt.Println(r)

	fmt.Println(rune(ra))

	fmt.Printf("%T\n", rune(ra))
	fmt.Printf("%T\n", ra)

	// It converts rune into String
	fmt.Println("Hello"+string(65));

	fmt.Println("----------------")
}




func dud() {
	fmt.Println("Harry Potter")
}

func with_params(x *int) {
	*x=0
}


// When a package is imported, the package name becomes an accessor for the contents.