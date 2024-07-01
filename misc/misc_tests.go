package main

import (
	"fmt"
)

func reverse(s []int) {
	
}

func main() {
	var x, y []int
	for i:=0; i<10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}