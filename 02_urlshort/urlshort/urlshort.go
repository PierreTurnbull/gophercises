package urlshort

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

type HandlerMap map[string]http.HandlerFunc

type UrlMap map[string]string

type YamlMap struct {
	ShortUrl string `yaml:"shortUrl"`
	DestUrl  string `yaml:"destUrl"`
}

func GetHandlerMap() HandlerMap {
	var pathsToUrls UrlMap

	for i, v := range os.Args {
		isYamlFlag := v == "--yaml"
		hasNextArg := i+1 < len(os.Args)

		var err error

		if isYamlFlag && hasNextArg {
			yamlPath := os.Args[i+1]

			pathsToUrls, err = YAMLHandler(yamlPath)

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if pathsToUrls == nil {
		pathsToUrls = UrlMap{
			"/urlshort-godoc1": "https://godoc.org/github.com/gophercises/urlshort",
			"/yaml-godoc1":     "https://godoc.org/gopkg.in/yaml.v2",
		}
	}

	handlerMap := make(HandlerMap)

	for key, value := range pathsToUrls {
		handlerMap[key] = func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, value, http.StatusMovedPermanently)
		}
	}

	return handlerMap
}

func YAMLHandler(yamlPath string) (UrlMap, error) {
	bytes, err := os.ReadFile(yamlPath)

	var yamlMap []YamlMap
	urlMap := make(UrlMap)

	yaml.Unmarshal(bytes, &yamlMap)

	for _, value := range yamlMap {
		urlMap[value.ShortUrl] = value.DestUrl
	}

	return urlMap, err
}
