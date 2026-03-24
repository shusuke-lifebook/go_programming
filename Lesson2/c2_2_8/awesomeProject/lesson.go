package main

import "fmt"

func main() {
	l := []string{"Python", "go", "java"}
	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}
}
