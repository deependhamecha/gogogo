package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("Dude")

	fmt.Println(average(5,6,4,3,2,1,9))

	fmt.Println(average(7))

	data:=[]float64{45,55}
	fmt.Println(average(data...))

	fmt.Println("dude"+strconv.Itoa(19))

	// Its a function written any where bt executes at the end, notice the order of function call

	fmt.Println("-------------------------")
	defer hello1()
	defer hello2()
	hello3()
	defer hello4()
}

func average(s ...float64) float64 {
	sum := 0.0

	for _, v := range s {
		sum += v
	}

	return sum/float64(len(s))
}


func hello1() {
	fmt.Println("Hello 1")
}


func hello2() {
	fmt.Println("Hello 2")
}


func hello3() {
	fmt.Println("Hello 3")
}


func hello4() {
	fmt.Println("Hello 4")
}
