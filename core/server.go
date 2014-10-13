package core

import (
	"encoding/json"
	"fmt"
	"io"
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

func HandleDeploy(w http.ResponseWriter, r *http.Request) {
	if LogLevel() >= LOG_INFO {
		log.Printf("Request on deploy")
	}
	if r.Method == "GET" {
		fmt.Fprint(w, "Got GET, needs POST request\n")
		return
	}

	// parse json
	pe, err := ParsePushEvent(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Deploy KO: %v\n", err)
		return
	}

	// match to project
	p, err := pushEventProject(pe)
	if err != nil {
		fmt.Fprintf(w, "Deploy KO: %v\n", err)
		return
	}

	// start the real work
	if LogLevel() >= LOG_VERBOSE {
		log.Printf("Got %v\n", p)
	}

	fmt.Fprint(w, "OK\n")
}

func ParsePushEvent(stream io.ReadCloser) (*PushEvent, error) {
	if LogLevel() >= LOG_VERBOSE {
		log.Println("Trying to parse PushEvent json")
	}
	decoder := json.NewDecoder(stream)
	p := new(PushEvent)
	err := decoder.Decode(p)
	return p, err
}

func pushEventProject(pe *PushEvent) (*Project, error) {
	for _, p := range GoployCtx.Cfg.Projects {
		if p.Url == pe.Repository.Url {
			return &p, nil
		}
	}
	return nil, fmt.Errorf("URL %v not found", pe.Repository.Url)
}
