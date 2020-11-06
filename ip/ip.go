package ip

import (
	"io/ioutil"
	"iptodns/config"
	"iptodns/utils"
	"log"
)

// GetIP can public ip address
func GetIP() string {
	resp, err := utils.Client.Get(config.Conf.GetIPURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(bodyBytes)
}
