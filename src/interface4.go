package main

import (
	"fmt"
)

type animal struct {
	sound string
	dude string
}

type dog struct {
	animal
	isFriendly bool
}

type cat struct {
	animal
	isFriendly bool
}


func data(d interface{}) {
	fmt.Println(d)
}

func main() {

	dog1 := dog{animal{sound: "Woof!"}, false}
	cat1 := cat{animal{sound: "Meow!"}, true}

	dog2 := dog{animal{dude: "WoofX!"}, false}
	cat2 := cat{animal{dude: "MeowX!"}, true}

	data(dog1)
	data(cat1)
	
	data(dog2)
	data(cat2)
}
