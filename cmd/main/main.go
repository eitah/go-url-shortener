package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/eitah/go-url-shortener/urlshort"
)

var yamlFilePathFlag string

func init() {
	flag.StringVar(&yamlFilePathFlag, "y", "", "Path to a custom map to direct your path to urls.")
	flag.Parse()
}

func main() {
	mux := defaultMux()

	// Build the map handler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
		"/tag":            "https://www.google.com",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// indenting this file messes up the yml. The override is useful.
	yaml := []byte(`
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /fish
  url: https://www.google.com
`)

	if yamlFilePathFlag != "" {
		var err error
		// override yml default
		yaml, err = ioutil.ReadFile(yamlFilePathFlag)
		if err != nil {
			panic(err)
		}
	}

	yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler) //nolint:errcheck
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
