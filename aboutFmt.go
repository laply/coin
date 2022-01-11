package main

import "fmt"

func aboutFmt() {
	x := 31415926535

	fmt.Printf("%o\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%x\n", x)
	fmt.Printf("%U\n", x)


	xAsBinary := fmt.Sprintf("%b\n", x) // 포멧팅해서 스트링으로 반납함
	fmt.Println(xAsBinary)

}