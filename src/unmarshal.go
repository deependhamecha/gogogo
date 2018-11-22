package main

import (
  "fmt"
  "encoding/json"
)

type Person struct {
  Name string
  Age int
}

func main() {
  person := Person{
    Name: "Deepen",
    Age: 27,
  }

  fmt.Println("Person:", person)

  bytes:= []byte(`{"Name":"Dhamecha","Age":27}`)

  marshal, _ := json.Marshal(person)

  json.Unmarshal(bytes, &person)

  fmt.Println("---------------------")

  fmt.Println("Person:", person)

  fmt.Println("Marshal:", marshal)

  json.Unmarshal(marshal, &person)

  fmt.Println("Umarshal:", unmarshalCode)
}
