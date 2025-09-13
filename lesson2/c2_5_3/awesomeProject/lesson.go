package main

import "log"

func main() {
	log.Println("logging!")
	log.Printf("%T %v", "test", "test")

	log.Fatalln("error!")
}
