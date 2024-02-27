package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
)

func highlight(source string) string {
	lexer := lexers.Get("go")
	if lexer == nil {
		lexer = lexers.Fallback
	}
	lexer = chroma.Coalesce(lexer)
	style := styles.Get("onedark")
	if style == nil {
		style = styles.Fallback
	}
	formatter := html.New(html.Standalone(false))

	iterator, err := lexer.Tokenise(nil, source)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	if err := formatter.Format(&buf, style, iterator); err != nil {
		panic(err)
	}
	return buf.String()
}

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

func removeSpecialLines(input string) string {
	var output bytes.Buffer
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)
		if !strings.HasPrefix(trimmedLine, "// S:") && !strings.HasPrefix(trimmedLine, "// E:") {
			output.WriteString(line + "\n")
		}
	}

	return output.String()
}

func wrap(fileName string, f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		data, err := os.ReadFile(fileName)
		if err != nil {
			panic(err)
		}

		content := highlight(removeSpecialLines(string(data)))
		fmt.Fprintf(w, "<!DOCTYPE html><html><body>")
		fmt.Fprintf(w, "<a href=\"/\"> <-- Back</a>")
		fmt.Fprintf(w, "<div style=\"display: flex;\"><div>")
		fmt.Fprintf(w, "<h2>Example Form</h2>")
		f(w, req)
		fmt.Fprintf(w, "</div><div style=\"width:5rem;\">")
		fmt.Fprintf(w, "</div><div>")
		fmt.Fprintf(w, "<h2 style=\"\">%s</h2>\n", fileName)
		fmt.Fprintf(w, "<div style=\"background:#282a36;padding: 1rem;\">\n"+string(content)+"\n</div>")
		fmt.Fprintf(w, "</div>")
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
