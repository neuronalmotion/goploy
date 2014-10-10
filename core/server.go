package core

import (
    "fmt"
    "log"
    "net/http"
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
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
