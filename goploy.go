package main

import (
	"fmt"
	"github.com/neuronalmotion/goploy/core"
)

func main() {
	fmt.Print(core.GoployCtx.Cfg.App.Port)

}
