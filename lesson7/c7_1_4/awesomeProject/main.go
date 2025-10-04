package main

import (
	"fmt"
	"net/url"
)

func main() {
	base, _ := url.Parse("http://example.com")
	reference, _ := url.Parse("/test?a=1&b=2")
	endopoint := base.ResolveReference(reference).String()
	fmt.Println(endopoint)
}
