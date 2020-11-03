package dns

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"iptodns/config"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// requestJson is request json body
type requestJson struct {
	DomainType string `json:"type"`
	Name       string `json:"name"`
	Content    string `json:"content"`
	TTL        int    `json:"ttl"`
}

// DNSRecordResp is Query dns records response
type DNSRecordResp struct {
	Success bool        `json:"success"`
	Result  []DNSRecord `json:"result"`
}

// DNSRecord is dns detail
type DNSRecord struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// PrintFalseInfo is print false requst info
func PrintFalseInfo(resp *http.Response) {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	body := string(bodyBytes)
	log.Println(body)
}

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

// AddAuthHeader add auth info into request Header
func AddAuthHeader(req *http.Request) {
	req.Header.Add("X-Auth-Email", "cong.lv.yx@gmail.com")
	req.Header.Add("Authorization", "Bearer "+config.Conf.Key)
}

// GenerateURL is generate cloudflare dns url
func GenerateURL(dnsURL, zoneIdentifier string) string {
	r := strings.NewReplacer("{zone_identifier}", zoneIdentifier)
	return r.Replace(dnsURL)
}

// GetCloudflareDNSList get dns list
func GetCloudflareDNSList(name string) []DNSRecord {
	var results DNSRecordResp

	url := CreateQuery(config.Conf.DNSURL, map[string]string{"name": name})
	req, err := http.NewRequest("GET", url, nil)
	AddAuthHeader(req)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := config.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Println(resp.StatusCode)
		PrintFalseInfo(resp)
		return nil
	}
	CustomUnmarshal(resp, &results)
	return results.Result
}

// UpdateCloudflare for sync new ip address to cloudflare
func UpdateCloudflare(key, ip, name string) bool {
	jsonBody, err := json.Marshal(requestJson{
		DomainType: config.Conf.DomainType,
		Content:    ip,
		Name:       name,
		TTL:        config.Conf.TTL,
	})
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("PUT", config.Conf.DNSURL+config.Conf.Identifier, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatal()
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	AddAuthHeader(req)

	resp, err := config.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		PrintFalseInfo(resp)
		return false
	}
	return true
}

// Setup cache
func Setup() {
	config.Conf.DNSURL = GenerateURL(config.Conf.DNSFormatURL, config.Conf.ZoneIdentifier)
	config.Conf.Identifier = GetCloudflareDNSList(config.Conf.Name)[0].ID
}
