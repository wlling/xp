package tests

import (
	"testing"

	_ "github.com/devopsxp/xp/module"
	"github.com/devopsxp/xp/pipeline"
)

func TestDemo(t *testing.T) {
	// config := pipeline.PipeConfig{
	// 	Name: "pipeline1",
	// 	Check: plugin.Config{
	// 		Name:        "ssh",
	// 		PluginTypes: plugin.CheckType,
	// 	},
	// 	Input: plugin.Config{
	// 		Name:        "kafka",
	// 		PluginTypes: plugin.InputType,
	// 	},
	// 	Filter: plugin.Config{
	// 		Name:        "upper",
	// 		PluginTypes: plugin.FilterType,
	// 	},
	// 	Output: plugin.Config{
	// 		Name:        "console",
	// 		PluginTypes: plugin.OutputType,
	// 	},
	// }

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
