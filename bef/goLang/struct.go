package main

import "fmt"

type person struct {
	name string
	age int
}

func (p person) sayHello(){
	fmt.Println("hi, I'm %s and %d", p.name, p.age)
}


func structFunc() {

	person1 := person{"nico", 12}
	person1.sayHello()


}