package main

import "fmt"

func foo(params ...int) {
	fmt.Println(params)
	for _, param := range params {
		fmt.Println(param)
	}
}

func main() {
	s := []int{1, 2, 3}
	fmt.Println(s)
	foo(s...)
}
