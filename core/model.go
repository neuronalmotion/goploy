package core

// God object to access important variables
// you should put as little as you can here
type GoployContext struct {
    Cfg Config
}

// ----------------------------------------------------------------------------
// goploy
// ----------------------------------------------------------------------------

type Config struct {
    App struct {
        Port    int         `json:"port"`
    }                       `json:"app"`
    Projects    []Project   `json:"projects"`
}

type Project struct {
    Path    string          `json:"path"`
    Deploy  string          `json:"deploy"`
}

// ----------------------------------------------------------------------------
// github/gitlab
// ----------------------------------------------------------------------------

type PushEvent struct {
    Ref         string       `json:"ref"`
    Repository struct {
        FullName    string   `json:"full_name"`
        Url         string   `json:"url"`
    }                        `json:"repository"`
}

