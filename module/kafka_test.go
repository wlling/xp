package module

import (
	"testing"

	"github.com/devopsxp/xp/pipeline"
)

func TestKafkaInput(t *testing.T) {
	config := pipeline.DefaultPipeConfig("pipeline1").
		WithCheckName("ssh").
		WithInputName("kafka").
		WithFilterName("upper").
		WithOutputName("console")

	p := pipeline.Of(*config)
	p.Init()
	p.Start()
	p.Exec()
	p.Stop()
}
