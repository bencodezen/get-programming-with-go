package main

import (
	"fmt"
)

func main() {
	const distance = 56000000
	const requiredDays = 28
	const requiredHours = requiredDays * 24

	var minSpeed = distance / requiredHours

	fmt.Println(minSpeed)
}
