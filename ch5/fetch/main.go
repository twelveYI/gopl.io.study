package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	// defer func ()  {
	// 	if closeErr := f.Close(); err == nil {
	// 		err = closeErr
	// 	}
	// }()
	n, err = io.Copy(f, resp.Body)
	if closeErr := f.Close(); err == nil {
		return local, n, closeErr
	}

	return local, n, err

}

// go run main.go http://gopl.io
func main() {
	for _, url := range os.Args[1:] {
		filename, n, err := fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %s %v\n", url, err)
		}
		fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", url, filename, n)
	}
}
