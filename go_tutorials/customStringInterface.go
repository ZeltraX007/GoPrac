package main

import "fmt"

type Stringer interface {
    String() string
}

type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

func main() {
    person := Person{Name: "John", Age: 30}

    fmt.Println(person)
}
