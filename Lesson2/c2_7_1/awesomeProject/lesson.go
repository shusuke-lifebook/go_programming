package main

import "fmt"

func thirdPartyConnectDB() {
	panic("Unable to connect database")
}

func save() {
	thirdPartyConnectDB()
}

func main() {
	save()
	fmt.Println("OK?")
}
