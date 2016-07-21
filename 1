package main

import (
	"flag"
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
)

var configPath string
var logfile string

func init() {
	flag.StringVar(&configPath, "config", "/tmp/ssinformer.toml", "judge config file path ")
	flag.StringVar(&logfile, "logfile", "/tmp/ssinformer.log", "log file")
	flag.Parse()
}
func main() {
	output := &lumberjack.Logger{
		Filename:   logfile,
		MaxSize:    100,
		MaxBackups: 5,
		MaxAge:     7,
	}
	log.SetOutput(output)
	cfg, err := ParseConfig(configPath)
	if err != nil {
		log.Println("[Error]", err)
		return
	}
	srv := NewServer(cfg)
	if err := srv.Open(); err != nil {
		log.Println("[Error]", err)
		return
	}
	srv.StartCrawlAndSend()
	select {}
}
