package main

import (
	"fmt"
	"net/http"
	"os"
)

func listing(examples []example) func(w http.ResponseWriter, req *http.Request) {

	listing := ""
	for _, e := range examples {
		listing = listing + fmt.Sprintf("<a href=\"/%s\">Example %s - %s</a><br/>", e.url, e.id, e.description)
	}

	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "<html><body>"+
			listing+
			"</body></html>")
	}
}

type examplefunc func(w http.ResponseWriter, req *http.Request)

type example struct {
	id          string
	description string
	url         string
	f           examplefunc
}

func wrap(fileName string, f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		data, err := os.ReadFile(fileName)
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(w, "<!DOCTYPE html><html><body>")
		fmt.Fprintf(w, "<a href=\"/\"> <-- Back</a>")
		fmt.Fprintf(w, "<h2>Example Form</h2>")

		f(w, req)

		fmt.Fprintf(w, "<h2 style=\"margin-top: 4rem;\">%s</h2>\n", fileName)
		fmt.Fprintf(w, "<pre style=\"background:#e0e0e0;padding: 1rem;\">\n"+string(data)+"\n</pre>")
		fmt.Fprintf(w, "</body></html>")
	}
}

func main() {
	examples := []example{
		{
			id:          "1",
			description: "Simple",
			url:         "example1",
			f:           wrap("cmd/example/example1.go", example1),
		},
		{
			id:          "2",
			description: "With Errors",
			url:         "example2",
			f:           wrap("cmd/example/example2.go", example2),
		},
		{
			id:          "3",
			description: "Nested",
			url:         "example3",
			f:           wrap("cmd/example/example3.go", example3),
		},
	}
	for _, e := range examples {
		http.HandleFunc("/"+e.url, e.f)
	}
	http.HandleFunc("/", listing(examples))

	fmt.Println("http://127.0.01:8090")
	http.ListenAndServe("127.0.0.1:8090", nil)
}
