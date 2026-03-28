package main

import "fmt"

func do(i interface{}) {
	i *= 2 // invalid operation: i *= 2 (mismatched types interface{} and untyped int)
	fmt.Println(i)
}

func main() {
	var i interface{} = 10
	do(i)
}
