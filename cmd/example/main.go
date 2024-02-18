package main

import (
	"fmt"
	"net/http"
)

func listing(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "<html><body>"+
		"<a href=\"/example1\">Example 1 - Simple</a><br/>"+
		"<a href=\"/example2\">Example 2 - with Errors</a><br/>"+
		"<a href=\"/example3\">Example 3 - Complex</a><br/>"+
		"</body></html>")
}

func main() {
	http.HandleFunc("/", listing)

	http.HandleFunc("/example1", example1)
	http.HandleFunc("/example2", example2)
	http.HandleFunc("/example3", example3)

	fmt.Println("http://127.0.01:8090")
	http.ListenAndServe("127.0.0.1:8090", nil)
}
