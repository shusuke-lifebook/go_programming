package main

import "fmt"

func main() {
	a := make([]int, 3)
	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)
}
