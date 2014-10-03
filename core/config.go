package core

import (
    "fmt"
    "encoding/json"
    "os"
)

var GoployCtx GoployContext

type GoployContext struct {
    Cfg Config
}

type Config struct {
    App struct {
        Port    int
    }
    Projects    []Project
}

type Project struct {
    Path    string
    Deploy  string
}

func init() {
    file, err := os.Open("goploy_conf.json")
    decoder := json.NewDecoder(file)
    decoder.Decode(&GoployCtx.Cfg)

    if err != nil {
        fmt.Println("error: ", err)
    }
}
