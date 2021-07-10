package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func main () {
	
	a := Page (
		Title: "Test Page"
		Body: "This is a test page"
	)
	fmt.Println(a.Body)
}