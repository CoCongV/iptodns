package ip

import (
	"iptodns/utils"
	"log"
)

// RespJSON is access ip api response
type RespJSON struct {
	Country     string `json:"country"`
	CountryCode string `json:"countryCode"`
	Region      string `json:"region"`
	RegionName  string `json:"regionName"`
	City        string `json:"city"`
	IP          string `json:"query"`
	ISP         string `json:"isp"`
}

// GetIP can public ip address
func GetIP() string {
	var result RespJSON
	resp, err := utils.Client.Get("http://ip-api.com/json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	utils.CustomUnmarshal(resp, &result)
	return result.IP
}
