package main

import "fmt"

func thirdPartyConnectionDB() {
	panic("Unable to connect database")
}

func save() {
	thirdPartyConnectionDB()
}

func main() {
	save()
	fmt.Println("OK?")
}
