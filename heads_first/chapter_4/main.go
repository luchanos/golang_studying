package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	fmt.Println("Enter the temperature in Fahreiheit: ")
	num, err := keyboard.GetFloat()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%0.2f degrees in Celsius\n\n", (num-32)*5/9)
}
