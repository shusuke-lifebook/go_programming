package main

import "fmt"

func main() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	v, ok := m["apple"]
	fmt.Println(v, ok)
	v2, ok2 := m["orange"]
	fmt.Print(v2, ok2)
}
