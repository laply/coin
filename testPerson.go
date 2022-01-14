package main

import (
	"fmt"
	"github.com/laply/coin/person"
)

func testPerson() {
	person1 := person.Person{}
	person1.SetDetails("hi", 3)
	fmt.Println("main", person1)
	person1.PrintDetails()
}