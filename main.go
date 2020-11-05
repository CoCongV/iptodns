package main

import (
	"context"
	"log"
	"time"

	"iptodns/config"
	"iptodns/dns"
	"iptodns/ip"
	"iptodns/rediscli"
)

func init() {
	config.Setup()
	dns.Setup()
	rediscli.Setup()
}

var ctx = context.Background()

func main() {
	for {
		localIP := ip.GetIP()
		log.Println(localIP)
		log.Println(config.Conf.OldIP)
		if config.Conf.OldIP != localIP {
			dns.UpdateCloudflare(config.Conf.Key, localIP, config.Conf.Name)
			config.Conf.OldIP = localIP
			err := rediscli.RDB.Set(ctx, config.Conf.RedisKey, localIP, 0).Err()
			if err != nil {
				log.Println(err)
				continue
			}
			log.Println("IP address has updated")
		} else {
			log.Println("IP address has not changed")
		}
		time.Sleep(time.Duration(10) * time.Second)
	}
}
