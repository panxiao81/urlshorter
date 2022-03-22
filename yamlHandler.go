package urlshorter

import (
	"gopkg.in/yaml.v2"
	"net/http"
)

// YAMLHandler will parse the provided YAML and then return
// a http.HandlerFunc (which also implements http.Handler)
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
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yml, &pathUrls)
	if err != nil {
		return nil, err
	}

	// convert YAML into map
	pathToUrls := make(map[string]string)
	for _, p := range pathUrls {
		pathToUrls[p.Path] = p.Url
	}
	return MapHandler(pathToUrls, fallback), nil
}

type pathUrl struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}
