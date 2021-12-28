package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func handleErr(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

func main() {
	const Url = "https://www.google.com/robots.txt"

	// GET

	resp, err := http.Get(Url)
	handleErr(err)
	// Print HTTP status
	fmt.Println(resp.Status)

	// Read and display response body
	body, err := ioutil.ReadAll(resp.Body)
	handleErr(err)
	fmt.Println(string(body))
	_ = resp.Body.Close()

	// HEAD

	resp, err = http.Head(Url)
	handleErr(err)
	_ = resp.Body.Close()
	fmt.Println(resp.Status)

	// POST

	form := url.Values{}
	form.Add("foo", "bar")
	resp, err = http.Post(
		Url,
		"application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		log.Panicln(err)
	}
	_ = resp.Body.Close()

	// DELETE

	req, err := http.NewRequest("DELETE", Url, nil)
	handleErr(err)
	var client http.Client
	resp, err = client.Do(req)
	handleErr(err)
	_ = resp.Body.Close()
	fmt.Println(resp.Status)

	// PUT

	req, err = http.NewRequest("PUT", Url, strings.NewReader(form.Encode()))
	handleErr(err)
	resp, err = client.Do(req)
	handleErr(err)
	_ = resp.Body.Close()
	fmt.Println(resp.Status)
}
