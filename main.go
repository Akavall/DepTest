package main

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

func main() {
	fmt.Println("Hello")

	resp, body, errs := gorequest.New().
		Get("http://localhost:8000").
		End()

	fmt.Println(errs)
	fmt.Println(resp)
	fmt.Println(body)
}
