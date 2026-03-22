package main

import "fmt"

func main() {
	n := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
}
