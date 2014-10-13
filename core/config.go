package core

import (
	"encoding/json"
	"log"
	"os"
)

const (
	configFile string = "/home/robin/devel/go/src/github.com/neuronalmotion/goploy/goploy_conf.json"
)

var GoployCtx GoployContext

func init() {
	// configure logging
	log.SetPrefix("goploy ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// load config file
	file, err := os.Open(configFile)
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
