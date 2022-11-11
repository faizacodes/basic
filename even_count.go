package main

import (
	"fmt"
)

func count_even(index int) int{
	event := 0
	for i := 0; i < index; i++{
		if i % 2 == 0{
			event += 1
		}
	}
	return event
}

func main() {
	fmt.Println(count_even(100))
}