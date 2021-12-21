package main

import (
	"log"
	"net/http"
)

func main() {
	// body
	http.HandleFunc("/",
					func(
					w http.ResponseWriter,
					r *http.Request) {
		w.Write([]byte("hello world"))
	})

	err := http.ListenAndServe(":3000", nil)

	if err !=nil {
		log.Fatal(err)
	}
}
