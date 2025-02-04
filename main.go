package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	base, err := url.Parse("https://httpbin.org/")
	if err != nil {
		fmt.Println(err)
		return
	}
	ref, err := url.Parse("/test?a=1&b=2")
	if err != nil {
		fmt.Println(err)
		return
	}
	endpoint := base.ResolveReference(ref).String()
	fmt.Println(endpoint)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	q := req.URL.Query()
	fmt.Println(q)
	// クエリに追加
	q.Add("c", "3&%")
	fmt.Println(q)
	// & is escaped to %26, % is escaped to %25
	fmt.Println(q.Encode())
}
