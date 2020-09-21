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

// Pipeline Config 工厂模式
func DefaultPipeConfig(name string) *PipeConfig {
	return &PipeConfig{
		Name:   name,
		Check:  plugin.Config{PluginTypes: plugin.CheckType},
		Input:  plugin.Config{PluginTypes: plugin.InputType},
		Filter: plugin.Config{PluginTypes: plugin.FilterType},
		Output: plugin.Config{PluginTypes: plugin.OutputType},
	}
}

func (p *PipeConfig) WithCheckName(name string) *PipeConfig {
	p.Check.Name = name
	return p
}

func (p *PipeConfig) WithInputName(name string) *PipeConfig {
	p.Input.Name = name
	return p
}

func (p *PipeConfig) WithFilterName(name string) *PipeConfig {
	p.Filter.Name = name
	return p
}

func (p *PipeConfig) WithOutputName(name string) *PipeConfig {
	p.Output.Name = name
	return p
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

// 最后定义pipeline的工厂方法，调用plugin.Factory抽象工厂完成pipelien对象的实例化：
// 保存用于创建Plugin的工厂实例，其中map的key为插件类型，value为抽象工厂接口
var pluginFactories = make(map[plugin.PluginType]plugin.Factory)

// 根据plugin.PluginType返回对应Plugin类型的工厂实例
func factoryOf(t plugin.PluginType) plugin.Factory {
	factory, _ := pluginFactories[t]
	return factory
}

// pipeline工厂方法，根据配置创建一个Pipeline实例
func Of(conf PipeConfig) *Pipeline {
	p := &Pipeline{}
	p.check = factoryOf(plugin.CheckType).Create(conf.Check).(plugin.Check)
	p.input = factoryOf(plugin.InputType).Create(conf.Input).(plugin.Input)
	p.filter = factoryOf(plugin.FilterType).Create(conf.Filter).(plugin.Filter)
	p.output = factoryOf(plugin.OutputType).Create(conf.Output).(plugin.Output)
	return p
}

// 初始化插件工厂对象
func init() {
	pluginFactories[plugin.CheckType] = &plugin.CheckFactory{}
	pluginFactories[plugin.InputType] = &plugin.InputFactory{}
	pluginFactories[plugin.FilterType] = &plugin.FilterFactory{}
	pluginFactories[plugin.OutputType] = &plugin.OutputFactory{}
}
