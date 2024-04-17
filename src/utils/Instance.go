package utils

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Instance struct {
	*http.Request
}

func CreateInstance(uri string, key string) (*Instance, error) {
	// http.NewRequest returns a pointer to a Request and an error
	if !strings.Contains(uri, "http://") && !strings.Contains(uri, "https://") {
		return nil, errors.New("please enter valid uri")
	}
	if len(key) == 0 {
		return nil, errors.New("please enter valid key")
	}
	var dataType string = ""
	var isXml = strings.Contains(strings.ToLower(uri), strings.ToLower("/xml"))
	var isJson = strings.Contains(strings.ToLower(uri), strings.ToLower("/json"))
	if !isJson && !isXml {
		uriLen := len(uri)
		lastChar := uri[uriLen-1]
		if lastChar == '/' {
			uri = uri[:len(uri)-1]
		}
		dataType = "/json"
	}

	if isXml {
		uri = strings.Replace(uri, "/xml", "/json", -1)
	}
	baseURL, err := url.Parse(uri + dataType)
	if err != nil {
		return nil, err
	}
	params := baseURL.Query() // Get a copy of the query values.
	params.Add("Key", key)    // Add your key
	baseURL.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
		return nil, err
	}
	return &Instance{req}, nil
}
