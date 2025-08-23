package main

import "fmt"

// var big int = 9223372036854775807
// var big int = 9223372036854775807 + 1
const big = 9223372036854775807 + 1

func main() {
	fmt.Println(big - 1)
}
