package pipeline

import (
	"fmt"

	"github.com/devopsxp/xp/plugin"
)

// Pipeline Config
type PipeConfig struct {
	Name   string
	Check  plugin.Config
	Input  plugin.Config
	Filter plugin.Config
	Output plugin.Config
}

// 对于插件化的系统，一切皆是插件，因此将pipeline也设计成一个插件，实现plugin接口
// pipeline管道的定义
type Pipeline struct {
	status plugin.StatusPlugin
	check  plugin.Check
	input  plugin.Input
	filter plugin.Filter
	output plugin.Output
}

// 一个消息的处理流程 check -> input -> filter -> output
func (p *Pipeline) Exec() {
	msg := p.check.Conn()
	if msg.Status == plugin.Ok {
		msg = p.input.Receive()
		msg = p.filter.Process(msg)
	}
	p.output.Send(msg)
}

// 启动的顺序 output -> filter -> input -> check
func (p *Pipeline) Start() {
	p.output.Start()
	p.filter.Start()
	p.input.Start()
	p.check.Start()
	p.status = plugin.Started
	fmt.Println("Pipeline started.")
}

// 停止的顺序 check -> input -> filter -> output
func (p *Pipeline) Stop() {
	p.check.Stop()
	p.input.Stop()
	p.filter.Stop()
	p.output.Stop()
	p.status = plugin.Stopped
	fmt.Println("Pipeline stopped.")
}

func (p *Pipeline) Status() plugin.StatusPlugin {
	return p.status
}

func (p *Pipeline) Init() {
	p.check.Init()
	p.input.Init()
	p.filter.Init()
	p.output.Init()
}
