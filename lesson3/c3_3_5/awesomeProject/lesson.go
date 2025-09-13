package main

import "fmt"

type Vertext struct {
	X int
	Y int
	S string
}

func main() {
	v2 := Vertext{X: 1}
	fmt.Println(v2)
}
