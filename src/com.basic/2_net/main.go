package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//Creating this custum writer by implementing Writer interface (by creating a Write Function)
type logWriter struct {
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(resp)

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

//implementing by Write function to logWriter
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
