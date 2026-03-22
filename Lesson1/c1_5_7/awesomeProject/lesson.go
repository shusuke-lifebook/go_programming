package main

import "fmt"

func main() {
	func(x int) {
		fmt.Println("inner func", x)
	}(1)
}
