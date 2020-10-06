package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello benchmark test!")
	num := RandInt()
	fmt.Println(num)
}

func RandInt() (int) {
	//return(9.4)
	return(rand.Int())
}