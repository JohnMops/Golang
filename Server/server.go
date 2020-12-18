package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	//"os"
	//"github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseFiles("./static/index.html"))
var dogTpl = template.Must(template.ParseFiles("./static/indexDog.html"))

func indexHandler(w http.ResponseWriter, r *http.Request)  {
	tpl.Execute(w, nil)
}

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
	dogTpl.Execute(w, nil)

	if r.Method != "GET" {
		http.Error(w, "Method is not allowed", http.StatusForbidden)
		return
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHanlder)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/index", indexHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
