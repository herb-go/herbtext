package texttemplate_test

import (
	"strings"
	"testing"

	"github.com/herb-go/herbtext"
	"github.com/herb-go/herbtext/texttemplate"
)

func TestConfig(t *testing.T) {
	var p *texttemplate.ParamConfig
	var err error
	p = &texttemplate.ParamConfig{}
	_, err = p.CreateParam(herbtext.DefaultEnvironment())
	if err == nil || !strings.Contains(err.Error(), "empty param") {
		t.Fatal(err)
	}
	p = &texttemplate.ParamConfig{
		Source: "Test",
		Parser: "notexist",
	}
	_, err = p.CreateParam(herbtext.DefaultEnvironment())
	if err == nil || !strings.Contains(err.Error(), " not found") {
		t.Fatal(err)
	}
	p = &texttemplate.ParamConfig{
		Source: "Test",
	}
	parser, err := p.CreateParam(herbtext.DefaultEnvironment())
	if err != nil {
		panic(err)
	}
	if parser.Source != "Test" {
		t.Fatal(parser)
	}
}
