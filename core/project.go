package core

import (
	"log"
	"os/exec"
	"strings"
    "syscall"
)

type Project struct {
	Url    string `json:"url"`
	Ref    string `json:"ref"`
	Path   string `json:"path"`
	Uid    int `json:"uid,omitempty"`
	Gid    int `json:"gid,omitempty"`
	Deploy string `json:"deploy,omitempty"`
}

func (p *Project) UpdateRepo() error {
	cmd := exec.Cmd{
		Path: "/usr/bin/git",
		Args: []string{"git", "pull", p.Url},
		Dir:  p.Path,
	}

    // gid has to be set before, if not, we might get
    // a permission denied while setting uid
    if p.Gid != 0 {
        if err := syscall.Setgid(p.Gid); err != nil {
            log.Fatalf("failed to Setgid(%d): %v", p.Gid, err)
        }
    }

    if p.Uid != 0 {
        if err := syscall.Setuid(p.Uid); err != nil {
            log.Fatalf("failed to SetUid(%d): %v", p.Uid, err)
        }
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
