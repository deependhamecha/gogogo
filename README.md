# Go Lang

$GOROOT = Should point to directory of go executable
$GOPATH = must include your directory with src, bin directory
$GOHOME = same as $GOPATH

## Packages

There are no classes in go lang but packages surely are. Now, functions and variables are visible outside a package if its name starts with a capital letter.
If not then it is private in that **package**. It has nothing to do with filename either.


name.go
```go
package dude

import "fmt"

func Dude() {
	fmt.Println("Dude")
}
```


main.go
```go
package main

import name "1_package_example/dude"

func main() {
	name.Dude()
}
```

## Commands

```sh
go build
```

Creates an executable beside the main function file but if it doesnt contain main then it wont create an executable file.

```sh
go run
```

```sh
go install
```

Puts the binary to bin folder. If main isnt present then also it will create an archive.

## Functions
Note : Last parameter of a function parameter must declare its type

## Init Block in function

```go
package main

import "fmt"

func main() {
	fmt.Print("Hello")
	{
		fmt.Println(" World")
	}	
}
```

## Anonymous Block

```go
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
```
