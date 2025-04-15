package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Pritam12F/url-shortener/store"
)

type InputLink struct {
	Url string `json:"url"`
}

func ShortenerHandler(w http.ResponseWriter, req *http.Request){
	var inputLink InputLink

	err := json.NewDecoder(req.Body).Decode(&inputLink)

	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}

	defer req.Body.Close()

	_, err2 := url.ParseRequestURI(inputLink.Url)	

	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	val, err3 := store.GetUrl(inputLink.Url)

	if err3 == nil {
		fmt.Fprintf(w, "Your url is %s", val.ShortenedUrl)
		return
	}
	urlEntry := store.AddUrl(inputLink.Url)

	fmt.Fprintf(w, "Your url is %s", urlEntry)
}

func RedirectHandler(w http.ResponseWriter, req *http.Request){
	id := req.PathValue("id")

	val, err := store.GetUrl(id)

	if err != nil {
		http.Error(w, err.Error(),http.StatusNotFound)
	}

	http.Redirect(w, req, val.OriginalUrl, http.StatusAccepted)
}