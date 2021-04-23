package main

import (
	"fmt"
	"net/http"

	"github.com/eitah/go-url-shortener/urlshort"
)

func main() {
	mux := defaultMux()

	// Build the map handler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"yaml-godoc":      "https://godoc.org/gopkg.in/yaml.v2",
		"tag":             "https://www.google.com",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// build the yaml handler using the mapHandler as a fallback
	// 	yaml := `
	// - path: /urlshort
	//   url: https://github.com/gophercises/urlshort
	// - path: /urlshort-final
	//   url: https://github.com/gophercises/urlshort/tree/solution
	// - path: /dogs
	//   url: https://www.somesite.com/a-story-about-dogs
	// `
	// 	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	fmt.Println("starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler) //nolint:errcheck
	// http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
