package core

import (
	"fmt"
	"log"
	"net/http"
)

func ServeHttp() {
	// configure routes
	http.HandleFunc("/", handler)
	http.HandleFunc("/deploy", HandleDeploy)

	// start webserver
	port := GoployCtx.Cfg.App.Port
	log.Printf("Webserver listening on %v...", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if LogLevel() >= LOG_VERBOSE {
		log.Printf("%v on path %v", r.Method, path)
	}

	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
