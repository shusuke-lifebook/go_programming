package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)
}
