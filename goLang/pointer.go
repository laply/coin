package main

import "fmt"

func pointer () {
	a := 2
	b := &a
	a = 12
	fmt.Println(*b)
}