package core

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "strings"
)

func ServeHttp() {
    http.HandleFunc("/", handler)
    port := GoployCtx.Cfg.App.Port
    log.Printf("Webserver listening on %v...", port)
    err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
    if err != nil {
        log.Fatal("ListenAndServe error: ", err)
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path
    routed := false
    if strings.HasPrefix(path, "deploy") {
        routed = true
        HandlePushEvent(w, r)
    }

    if !routed {
        fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
    }
}

func HandlePushEvent(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Deploy route")
}

func ParsePushEvent(stream io.ReadCloser) *PushEvent {
    log.Println("Trying to parse PushEvent json")
    decoder := json.NewDecoder(stream)
    p := new(PushEvent)
    decoder.Decode(p)
    return p
}

func TestBla() {
}

