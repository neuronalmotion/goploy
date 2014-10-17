package core

import (
	"log"
	"os/exec"
	"strings"
)

type Project struct {
	Url    string `json:"url"`
	Ref    string `json:"ref"`
	Path   string `json:"path"`
	Deploy string `json:"deploy"`
}

func (p *Project) UpdateRepo() error {
	cmd := exec.Cmd{
		Path: "/usr/bin/git",
		Args: []string{"git", "pull", p.Url},
		Dir:  p.Path,
	}
	out, err := cmd.Output()
	if err != nil {
		log.Printf("%v", err)
	}
	log.Printf("%s", out)
	return err
}

func (p *Project) DeployCmd() error {
    if p.Deploy == "" {
        return nil
    }
	parts := strings.Fields(p.Deploy)
	head := parts[0]
	parts = parts[1:len(parts)]

	cmd := exec.Command(head, parts...)
	cmd.Dir = p.Path
	out, err := cmd.Output()
	if err != nil {
		log.Printf("%v", err)
	}
	log.Printf("%s", out)
	return err
}
