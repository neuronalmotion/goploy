package core

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

var GoployCtx GoployContext

const (
	configFile string = "goploy_conf.json"
)

var flagconf *string = flag.String("conf", configFile, "path to config file")

func init() {
	log.SetPrefix("goploy ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func ParseArgs() {
	flag.Parse()
}

func LoadConfig(conf string) {
	config := conf
	if config == "" {
		config = *flagconf
	}
	file, err := os.Open(config)
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
