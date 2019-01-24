package main

import (
	"net/http"

	"github.com/hemv/url-shortener/handler"
)

func main() {
	mux := defaultMux()
	pathToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := handler.MapHandler(pathToUrls, mux)
	http.ListenAndServe(":8000", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
