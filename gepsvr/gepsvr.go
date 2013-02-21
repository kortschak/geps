package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"net/http"
	"html/template"
	"github.com/daviddengcn/go-villa"
	"github.com/russross/blackfriday"
)

var webRoot villa.Path = `./web`

var processors map[string]http.HandlerFunc = map[string]http.HandlerFunc{}

func registerPath(path string, f http.HandlerFunc) {
	log.Println("Register path:", path)
	processors[path] = f
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if proc, ok := processors[path]; ok {
		proc(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	host := ":8081"
	if len(os.Args) > 1 {
		host = os.Args[1]
	}
	http.HandleFunc("/", handler)
	log.Println("Back server listening at", host)
	http.ListenAndServe(host, nil)
}


// for gep files
func __print__(response http.ResponseWriter, s interface{}) {
    response.Write([]byte(fmt.Sprint(s)))
}

/* <html>$text</html> */
func Html(text interface{}) string {
	return strings.Replace(template.HTMLEscapeString(fmt.Sprint(text)), "\n", "<br>", -1)
}


/* <input attr='$text'> */
func Value(text interface{}) string {
	return template.HTMLEscapeString(fmt.Sprint(text))
}

/* http://xxx.xxx/?xxx=$text */
func Query(text interface{}) string {
	return template.URLQueryEscaper(fmt.Sprint(text))
}

/* <script> s='$text' </script> */
func JS(text interface{}) string {
	return template.JSEscaper(fmt.Sprint(text))
}

// Markdown converts a markdown markup text into HTML
func Markdown(text interface{}) string {
	return string(blackfriday.MarkdownCommon([]byte(fmt.Sprint(text))))
}
