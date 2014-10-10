package core

import (
    "testing"
    "os"
    "github.com/stretchr/testify/assert"
)

const (
    pushEventFile string = "testdata/pushevent.json"
)

func TestParsePushEvent(t *testing.T){
    assert := assert.New(t)
    file, err := os.Open(pushEventFile)
    if err != nil {
        t.Errorf("Could not load input test %v, err: %v", pushEventFile, err)
    }
    p := ParsePushEvent(file)
    assert.NotNil(p)
    assert.Equal(p.Ref, "refs/heads/gh-pages")

    r := p.Repository
    assert.Equal(r.FullName, "baxterthehacker/public-repo")
    assert.Equal(r.Url, "https://github.com/baxterthehacker/public-repo")
}

