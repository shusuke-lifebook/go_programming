package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	base, _ := url.Parse("http://example.com")
	reference, _ := url.Parse("test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	q := req.URL.Query()
	q.Add("c", "3&%")
	fmt.Println(q.Encode())
	req.URL.RawQuery = q.Encode()

	var client *http.Client = &http.Client{} // クライアントを作成
	resp, _ := client.Do(req)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
