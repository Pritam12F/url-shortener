package store

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
)

type Link struct {
	Id string	`json:"id"`
	OriginalUrl string `json:"originalUrl"`
	ShortenedUrl string `json:"shortenedUrl"`
}

var UrlStore map[string]Link = make(map[string]Link)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[0:3])
}

func AddUrl(originalUrl string) string {
	id := GetMD5Hash(originalUrl)

	UrlStore[id] = Link{
		Id: id,
		OriginalUrl: originalUrl,
		ShortenedUrl: "http://localhost:8080/r/"+id,
	}

	return "http://localhost:8080/r/"+id
}

func GetUrl(id string) (Link, error) {
	val, ok := UrlStore[id]

	if !ok {
		return Link{}, errors.New("URL isn't registered")
	}

	return val, nil
}