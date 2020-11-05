package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Config struct storage config
type Config struct {
	GetIPURL       string `toml:"ip_url"`
	DNSFormatURL   string `toml:"dns_url"`
	ZoneIdentifier string `toml:"zone_identifier"`
	Email          string `toml:"email"`
	DomainType     string `toml:"domain_type"`
	Name           string `toml:"name"`
	TTL            int    `toml:"ttl"`
	Key            string `toml:"key"`
	RedisServer    string `toml:"redis_server"`
	RedisPassword  string `toml:"redis_password"`
	RedisDB        int    `toml:"redis_db"`
	RedisKey       string `toml:"redis_key"`
	DNSURL         string
	Payload        string
	Identifier     string
	OldIP          string
}

// Conf is Config struct's point
var Conf *Config

// Setup is read toml config file for init
func Setup() {
	filepath := "./config.toml"
	_, err := toml.DecodeFile(filepath, &Conf)
	if err != nil {
		log.Fatal(err)
	}
}
