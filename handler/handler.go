package handler

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

// ShortPath will serve as container for YAML string input
type ShortPath struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if url, ok := pathToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, url, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yamlData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYAML, err := parseYAML(yamlData)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYAML)
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(yamlData []byte) ([]ShortPath, error) {
	parsedYAML := []ShortPath{}
	err := yaml.Unmarshal(yamlData, &parsedYAML)
	if err != nil {
		return nil, err
	}
	return parsedYAML, nil
}

func buildMap(parsedYAML []ShortPath) map[string]string {
	pathMap := make(map[string]string)
	for _, shortPath := range parsedYAML {
		pathMap[shortPath.Path] = shortPath.URL
	}
	return pathMap
}
