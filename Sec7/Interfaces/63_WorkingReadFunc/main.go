package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)

	resp.Body.Read(bs)

	fmt.Println(string(bs))

	ioutil.WriteFile("main.html", bs, 0666)
}
