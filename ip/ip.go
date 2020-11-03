package ip

import (
	"io/ioutil"
	"log"
	"net/http"
)

// GetIP can public ip address
func GetIP() string {
	resp, err := http.Get("https://api.ip.sb/ip")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	body := string(bodyBytes)
	return body
}
