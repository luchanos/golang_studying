package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var status string
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	input = strings.TrimSpace(input)
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatalln(err)
	}
	if num >= 60 {
		status = "Passed!"
	} else {
		status = "Not passed!"
	}
	fmt.Println(status)
}
