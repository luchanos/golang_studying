package main

import (
	"fmt"
	"github.com/luchanos/golang_studying/heads_first/dates"
)

func main() {
	days := 3
	fmt.Printf("Your appointment in %d days\n", days)
	fmt.Printf("with a follow-up in %d days\n", days+dates.DaysInWeek)
}
