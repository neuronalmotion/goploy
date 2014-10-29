package core

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

const (
	configFile string = "goploy_conf.json"
)

var conf *string = flag.String("conf", configFile, "path to config file")

var GoployCtx GoployContext

func ParseArgs() {
	flag.Parse()
}

func LoadConfig() {
	// configure logging
	log.SetPrefix("goploy ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// load config file
	file, err := os.Open(*conf)
	decoder := json.NewDecoder(file)
	decoder.Decode(&GoployCtx.Cfg)

	if err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}
	if LogLevel() > LOG_SILENT {
		log.Println("Config file loaded")
	}
}

func LogLevel() int {
	return GoployCtx.Cfg.App.LogLevel
}
