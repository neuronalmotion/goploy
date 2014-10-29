package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	conf string = "testdata/conf.json"
)

func TestParseArgs(t *testing.T) {
	assert := assert.New(t)
	ParseArgs()

	assert.Equal(*flagconf, configFile, "Default value of --conf is bad")
}

func TestLoadConfig(t *testing.T) {
	assert := assert.New(t)
	LoadConfig(conf)

	cfg := GoployCtx.Cfg
	assert.NotNil(cfg)

	assert.Equal(cfg.App.Port, 8000, "invalid App.Port value")
	assert.Equal(cfg.App.LogLevel, 2, "invalid App.LogLevel value")
	assert.Equal(len(cfg.Projects), 1, "invalid Projects length")

	p := cfg.Projects[0]
	assert.Equal(p.Url, "https://github.com/baxterthehacker/public-repo", "invalid project url")
	assert.Equal(p.Ref, "refs/heads/gh-pages", "invalid project ref")
	assert.Equal(p.Path, "/home/robin/devel/public-repo", "invalid project path")
	assert.Equal(p.Deploy, "echo hello world", "invalid project deploy")
}
