package main

import "fmt"

func do(i interface{}) {
	ii := i * 2 // invalid operation: i * 2 (mismatched types interface{} and untyped int)
	fmt.Println(ii)
}

func main() {
	do(10)
}
