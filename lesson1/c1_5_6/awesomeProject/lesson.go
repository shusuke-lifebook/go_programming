package main

import "fmt"

func main() {
	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)
}
