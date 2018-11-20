package main

import "fmt"

func main() {
	// Arrays
	// var a [5]int // Working
	// var a [3]int{1,2,3} // Not Working
	a := [3]int{1,2,3}

	fmt.Println(a)

	fmt.Println(len(a))

	fmt.Println(a[0])

	fmt.Println("a[1:2] : ",a[1:3])

	// If you do not define array size, then its a slice(list)
	var mySlice = []int{3,1,2}

	fmt.Println(mySlice)

	fmt.Println(string("Dude"[2]))

	fmt.Println("------------------")

	for i:= range mySlice {
		fmt.Print("len : ", len(mySlice))
		fmt.Print(" cap : ", cap(mySlice))
		fmt.Println(" value : ", mySlice[i])
	}


	// Now there are 2 things: length which we regularly use in other languages and there is another thing called capacity
	// By Default if you dont specify capacity then its same as length
	// However, if you want to specify then you do it via make() method
	// make() declares and initializes as well

	s:=make([]int, 5, 100)

	fmt.Println("------------------")
	for i:= range s {
		fmt.Print("len : ", len(s))
		fmt.Print(" cap : ", cap(s))
		fmt.Println(" value : ", s[i])
	}

	// So, the concept of capacity is that your slice will have no problem while increasing until 100, but if it increases above
	// 100 then it will recreate whole old array into new one with increased capacity to 101 then 102, etc, see below output of append()

	t:=make([]int, 2, 3)

	t[0]=7
	t[1]=8
	// t[2]=9 // You cannot just point to next index and add element as index doesnt exist due to length
	// So you add element into slice via append() method

	t=append(t, 9)
	t=append(t, 10)	// Check Capacity after this statement execution

	fmt.Println("-----------------")
	for i:= range t {
		fmt.Print("len : ", len(t))
		fmt.Print(" cap : ", cap(t))
		fmt.Println(" value : ", t[i])
	}

	// Now to Delete an element is very lame
	n := 3
	t = append(t[:n], t[n+1:]...) // This is kind of a formula

	fmt.Println("-----------------")

	fmt.Println("Deleted t[3] element: ", t)


	// Append an Array

	r := []int{4,5,6}

	t = append(t, r...)

	fmt.Println("------------------")
	fmt.Println("Appended an Array : ", t)


	fmt.Println("------------------")
	// Maps
	m := make(map[string]int)

	m["k1"] = 7;
	m["k2"] = 9;

	fmt.Println("Map : ", m)

	// it returns two values, 1st one of dumped other is stored in p
	_, p := m["k1"]

	fmt.Println(p);

	fmt.Println("--------------")

	_,e := return2Values()
	fmt.Println(e)

	// Another way of declaring and initializing map
	k := map[string]int{"key1": 5, "key2": 6}
	fmt.Println("Another way : ", k);

	// Delete something from map
	// Notice map is a reference so assigning isnt necessary
	delete(k, "key2")
	fmt.Println("After deleting : ", k);

	var h1 map[string]int
	// same as
	h2 := map[string]int{}
	fmt.Println("h1: ", h1 == nil)
	fmt.Println("h2: ", h2 == nil)

	// when you fetch the element, it returns 2 return values
	if val, exists := k["key1"]; exists {
		fmt.Println(val)
	}

	fmt.Println("---------------")

	// Multi dim maps
	y:= map[int]map[int]string{
		0: map[int]string{
			0: "Dude",
		}, // Comma is important
	}

	fmt.Println(y)

}


func return2Values() (string, string) {
	return "Deepen", "Dhamecha"
}
