package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHanlder(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello")
}

func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Dog Page")
	return
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHanlder)
	http.HandleFunc("/dog", dog)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
