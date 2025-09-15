package main

import "fmt"

func do(i interface{}) {
	ii := i * 2
	fmt.Println(ii)
}

func main() {
	do(10)
}
