package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type myWriter struct{}

func (myWriter) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	fmt.Println("My Writer")
	return len(p), nil
}
func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		os.Exit(1)
	}
	io.Copy(new(myWriter), resp.Body)
}
