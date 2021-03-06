package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type info struct {
	Hostname string
	Message  string
}

func indexHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Handling request...")
	tmplt := template.New("index.html")
	tmplt, _ = tmplt.ParseFiles("index.html")

	hostname, _ := os.Hostname()
	message := "Hello World"

	if os.Getenv("message") != "" {
		message = os.Getenv("message")
	}

	p := info{
		Hostname: hostname,
		Message:  message,
	}

	tmplt.Execute(response, p)
}

func simpleHelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s!", r.URL.Path[1:])
}

func main() {
	fmt.Println("Server Starting")
	fmt.Println("Setting up handlers")
	http.HandleFunc("/hello-app", simpleHelloHandler)
	http.HandleFunc("/hello", indexHandler)
	fmt.Println("Serving on port 8080")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
