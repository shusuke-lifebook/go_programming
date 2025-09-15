package main

import "fmt"

func do(i interface{}) {
	i *= 2
	fmt.Println(i)
}

func main() {
	var i interface{} = 10
	do(i)
}
