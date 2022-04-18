package main

import "fmt"

func main() {
	a := 1
	aPointer := &a
	fmt.Println(a, aPointer)
	*aPointer = 154
	fmt.Println(a, aPointer)
}
