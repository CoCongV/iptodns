package main

import (
	"log"
	"time"

	"iptodns/config"
	"iptodns/dns"
	"iptodns/ip"
)

func init() {
	config.Setup()
	dns.Setup()
}

func main() {
	for {
		localIP := ip.GetIP()
		log.Println(localIP)
		log.Println(config.Conf.OldIP)
		if config.Conf.OldIP != localIP {
			dns.UpdateCloudflare(config.Conf.Key, localIP, config.Conf.Name)
			config.Conf.OldIP = localIP
			log.Println("IP address has updated")
		} else {
			log.Println("IP address has not changed")
		}
		time.Sleep(time.Duration(10) * time.Second)
	}
}
