package main

import "fmt"



func plus (a, b int, name string ) ( int, string ){
	return a + b, name
}


func manyPlus (a ...int) int {
	var sum int
	for _, item := range a{
		sum += item 
	}

	return sum
}


func function (){
	result, name := plus(20, 8, "laply")
	fmt.Println(result, name)

	result2 := manyPlus(1, 2, 3, 4, 5, 6, 7)

	fmt.Println(result2)


	name2 := "abcdefghijklmnop"
	for index, letter := range name2 {
		fmt.Println(index, string(letter))
	}


}