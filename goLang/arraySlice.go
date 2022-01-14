package main

import "fmt"

func arraySlice () {
	foods := [3]string{"pizza", "pasta", "banana"} // 길이가 정해저 있는 array

	for _, dish := range foods {
		fmt.Println(dish)
	}

	for i := 0; i < len(foods); i++ {
		fmt.Println(foods[i])
	}


	sliceFoods := []string{"pizza", "pasta", "banana"} // 길이가 정해져 있지 않는 slice
	sliceFoods = append(sliceFoods, "apple" )

	fmt.Printf("%v\n", sliceFoods)

}