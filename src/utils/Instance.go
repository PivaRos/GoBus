package utils

import (
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
	var dataType string = ""
	var isXml = strings.Contains(strings.ToLower(uri), strings.ToLower("/xml"))
	var isJson = strings.Contains(strings.ToLower(uri), strings.ToLower("/json"))
	if !isJson || !isXml {
		dataType = "/json"
	}

	baseURL, err := url.Parse(uri + dataType)

	if err != nil {
		return nil, nil
	}
	params := baseURL.Query() // Get a copy of the query values.
	params.Add("Key", key)    // Add your key
	baseURL.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
		return nil, nil
	}
	return &Instance{req}, nil
}
