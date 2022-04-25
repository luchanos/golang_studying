package main

import (
	"dates"
	"fmt"
)

func main() {
	days := 3
	fmt.Printf("Your appointment in %d days\n", days)
	fmt.Printf("with a follow-up in %d days\n", days+dates.DaysInWeek)
}
