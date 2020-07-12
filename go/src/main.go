package main

import (
	"fmt"
	"log"
	"net/http"
)

func greeting(s string) string {
	return "<b>" + s + "</b>"
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, greeting("Code.education Rocks"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
