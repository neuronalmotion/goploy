package main

import (
	"github.com/neuronalmotion/goploy/core"
)

func main() {
	core.ParseArgs()
	core.LoadConfig()
	core.ServeHttp()
}
