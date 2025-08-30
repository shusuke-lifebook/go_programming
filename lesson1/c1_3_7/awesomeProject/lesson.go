package main

import "fmt"

func main() {
	fmt.Println("Hello World"[0]) // ASCIIコードで取得、表示される。
	fmt.Println(string("Hello World"[0]))
}
