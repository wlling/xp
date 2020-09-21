package module

import (
	"testing"

	"github.com/devopsxp/xp/pipeline"
)

func TestDemo1(t *testing.T) {
	config := pipeline.DefaultPipeConfig("pipeline1").
		WithCheckName("ssh").
		WithInputName("hello").
		WithFilterName("upper").
		WithOutputName("console")

	p := pipeline.Of(*config)
	p.Init()
	p.Start()
	p.Exec()
	p.Stop()
}
