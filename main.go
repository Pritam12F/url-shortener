package main

import (
	"net/http"

	"github.com/Pritam12F/url-shortener/handlers"
)

func main(){
	
	http.HandleFunc("/shorten", handlers.ShortenerHandler)
	http.HandleFunc("/r/{id}", handlers.RedirectHandler)

	http.ListenAndServe(":8080", nil)
}