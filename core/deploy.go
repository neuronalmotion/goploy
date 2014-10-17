package core

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func HandleDeploy(w http.ResponseWriter, r *http.Request) {
	if LogLevel() >= LOG_INFO {
		log.Printf("Request on deploy")
	}
	if r.Method == "GET" {
		fmt.Fprint(w, "Got GET, needs POST request\n")
		return
	}

	// parse json
	pe, err := parsePushEvent(r.Body)
	if err != nil {
		log.Printf("Deploy KO (parsing pushevent): %v\n", err)
		fmt.Fprintf(w, "Deploy KO (parsing pushevent): %v\n", err)
		return
	}

	// match to project
	p, err := pushEventProject(pe)
	if err != nil {
		fmt.Fprintf(w, "Deploy KO (finding project): %v\n", err)
		return
	}

	// start the real work
	if LogLevel() >= LOG_VERBOSE {
		log.Printf("Got %v\n", p)
	}
	err = p.UpdateRepo()
	if err != nil {
		fmt.Fprintf(w, "Deploy KO (updating project repo): %v\n", err)
		return
	}

	err = p.DeployCmd()
	if err != nil {
		fmt.Fprintf(w, "Deploy KO (running deploy script): %v\n", err)
		return
	}

	fmt.Fprint(w, "OK\n")
}

func parsePushEvent(stream io.ReadCloser) (*PushEvent, error) {
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
		if p.Url == pe.Repository.Url && p.Ref == pe.Ref {
			return &p, nil
		}
	}
	return nil, fmt.Errorf("URL %v not found", pe.Repository.Url)
}
