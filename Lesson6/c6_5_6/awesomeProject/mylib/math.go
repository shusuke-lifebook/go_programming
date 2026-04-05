// Package mylib provides math utilities.
package mylib

import "fmt"

// Average returns the average of a series of numbers
func Average(s []int) int {
	total := 0
	for _, i := range s {
		total += i
	}
	return total / len(s)
}

type Person2 struct {
	// Name
	Name string
	// Age
	Age int
}

func (p *Person2) Say() {
	fmt.Println("Person2")
}
