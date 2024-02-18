package main

import (
	"fmt"
	"net/http"
)

func listing(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "<html><body>"+
		"<a href=\"/example1\">Example 1</a><br/>"+
		"<a href=\"/example2\">Example 2</a><br/>"+
		"</body></html>")
}

type data struct {
	Name string
}

func main() {
	http.HandleFunc("/", listing)

	http.HandleFunc("/example1", example1)
	http.HandleFunc("/example2", example2)

	fmt.Println("http://127.0.01:8090")
	http.ListenAndServe("127.0.0.1:8090", nil)
}
