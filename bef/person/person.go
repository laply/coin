package person

import "fmt"

type Person struct {
	name string
	age int
}

func (p* Person) SetDetails(name string, age int) {
	p.name = name
	p.age = age
}

func (p Person) PrintDetails(){
	fmt.Println(p)
}