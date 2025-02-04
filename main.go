package main

import (
	"bytes"
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

	// HTTPクライアントを作成
	var client *http.Client = &http.Client{}
	// リクエストを送信
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

	// POST request
	ref, err = url.Parse("/status")
	if err != nil {
		fmt.Println(err)
		return
	}
	endpoint = base.ResolveReference(ref).String()
	fmt.Println(endpoint)
	req, err = http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte("test")))
	if err != nil {
		fmt.Println(err)
		return
	}

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
