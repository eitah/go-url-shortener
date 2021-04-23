package urlshort

import (
	"fmt"
	"net/http"
)

// MapHandler will return an http.HandlerFunc which also implements http.Handler that will attempt
// to map any paths (keys) to their corresponding URL string values. If the path is not provided in
// the map, then the fallback http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) (handlerFunc http.HandlerFunc) {
	// todo implement this
	return func(writer http.ResponseWriter, req *http.Request) {
		fmt.Printf("redirect to %s...\n", req.URL.Path)
		for key, redirectUrl := range pathsToUrls {
			if req.URL.Path == key {
				// implementation 1: try just redirect which works but is not what fallback does
				// http.Redirect(writer, req, redirectUrl, 301)

				mux := http.NewServeMux()
				fmt.Printf("redirect to %s bc matched %s!\n", redirectUrl, key)
				mux.HandleFunc(key, func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintln(writer, redirectUrl)
				})
				mux.ServeHTTP(writer, req)
				return
			}
		}
		fallback.ServeHTTP(writer, req)
	}
}

// YAMLHandler will parse the provided YAML and return a handler func which also implements http.Handler
// that will attempt to map any paths to their corresponding URL. If the path is not provided in this
// YAML, then the fallback http.Handler will be called instead. Yaml is expected int he format
//
// - path: /some-path
//   url: https://www.some-url.com/demo
//
// the only errors that can be returned all related to having invalid YAML  data.
//
// See MapHandler to create a similar http.HandlerFunc via a mapping of paths to URLs.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// todo implement this
	return nil, nil
}
