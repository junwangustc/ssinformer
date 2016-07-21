package main

import (
	"log"
	"os"

	"github.com/junwangustc/ssinformer/crawl"
	"github.com/junwangustc/ssinformer/smtp"
	"github.com/naoina/toml"
)

type Config struct {
	SMTP    smtp.Config  `toml:"smtp"`
	Crawler crawl.Config `toml:"crawler"`
}

func NewConfig() *Config {
	c := &Config{}
	c.SMTP = smtp.NewConfig()
	c.Crawler = crawl.NewConfig()
	return c
}
func ParseConfig(path string) (cfg *Config, err error) {
	if path == "" {
		log.Fatalln("no configuration provided, using default settings")
	}
	log.Printf("Using configuration at: %s\n", path)
	config := NewConfig()
	f, err := os.Open(path)
	if err != nil {
		return cfg, err
	}
	defer f.Close()
	return config, toml.NewDecoder(f).Decode(&config)
}
