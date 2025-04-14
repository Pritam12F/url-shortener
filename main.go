package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, req *http.Request){
	

	fmt.Fprintf(w, "hello\n")
}

func main(){
	http.HandleFunc("/", rootHandler)

	http.ListenAndServe(":8080", nil)
}