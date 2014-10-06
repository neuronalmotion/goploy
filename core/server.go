package core

import (
    "net/http"
    "fmt"
)

func ServeHttp() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(fmt.Sprintf(":%v", GoployCtx.Cfg.App.Port), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
