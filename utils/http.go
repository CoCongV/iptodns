package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//Client is global http client
var Client = &http.Client{}

// CustomUnmarshal is undump json body
func CustomUnmarshal(resp *http.Response, r interface{}) {
	body, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, r); err != nil {
		log.Fatal(err)
	}
}

// CreateQuery create request url
func CreateQuery(baseURL string, params map[string]string) string {
	r := url.Values{}
	for k, v := range params {
		r.Set(k, v)
	}
	return baseURL + "?" + r.Encode()
}
