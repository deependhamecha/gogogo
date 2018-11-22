package main

import "fmt"

import "encoding/json"

// Variables need to be exported
// json:keyname is optional and is alias for variable name in json/marshal
type Person struct {
  Name string `json:"myname"`
  Age int `json:"myage"`
}

func createObj() Person {

  return Person{
    Name: "Deepen",
    Age: 27,
  }
}

func main() {

  person1 := createObj();

  person2 := Person{
    Name: "Dhamecha",
    Age: 27,
  }

  fmt.Println("Object: ", person1)
  marshal1, _ := json.Marshal(person1);
  fmt.Println("Marshal:", marshal1)
  fmt.Printf("Type %T\n", marshal1)
  fmt.Println("String: ", string(marshal1))

  fmt.Println()
  fmt.Println("Object: ", person2)
  marshal2, _ := json.Marshal(person2);
  fmt.Println("Marshal:", marshal2)
  fmt.Printf("Type %T\n", marshal2)
  fmt.Println("String: ", string(marshal2))

  // Same Representation of marshal, but marshal is a slice of unit8 where this is byte
  bytes:= []byte(`{"Name":"Dhamecha","Age":27}`)


}
