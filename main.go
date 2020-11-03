package main

import (
	"log"

	"iptodns/config"
	"iptodns/dns"
	"iptodns/ip"
)

func init() {
	config.Setup()
	dns.Setup()
}

func main() {
	localIP := ip.GetIP()
	log.Print(localIP)
	log.Print(config.Conf.Identifier)
	dns.UpdateCloudflare(config.Conf.Key, localIP, config.Conf.Name)
}
