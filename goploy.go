package main

import (
	"github.com/neuronalmotion/goploy/core"
)

func main() {
	core.ParseArgs()
	core.LoadConfig("") // we use what has been parsed as conf path
	core.ServeHttp()
}
