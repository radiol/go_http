package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	base, err := url.Parse("https://httpbin.org/")
	if err != nil {
		fmt.Println(err)
		return
	}
	ref, err := url.Parse("/ip")
	if err != nil {
		fmt.Println(err)
		return
	}
	endpoint := base.ResolveReference(ref).String()
	fmt.Println(endpoint)

	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
