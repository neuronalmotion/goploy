package core

// God object to access important variables
// you should put as little as you can here
type GoployContext struct {
	Cfg Config
}

const (
	LOG_SILENT int = iota
	LOG_INFO
	LOG_VERBOSE
)

// ----------------------------------------------------------------------------
// goploy
// ----------------------------------------------------------------------------

type Config struct {
	App struct {
		Port     int `json:"port"`
		LogLevel int `json:"log_level"`
	} `json:"app"`
	Projects []Project `json:"projects"`
}

// ----------------------------------------------------------------------------
// github/gitlab
// ----------------------------------------------------------------------------

type PushEvent struct {
	Ref        string `json:"ref"`
	Repository struct {
		FullName string `json:"full_name"`
		Url      string `json:"url"`
	} `json:"repository"`
}
