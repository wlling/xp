package module

import (
	"fmt"
	"reflect"

	. "github.com/devopsxp/xp/plugin"
)

func init() {
	// 初始化shell filter插件映射关系表
	AddFilter("shell", reflect.TypeOf(ShellFilter{}))
}

// shell 命令运行filter插件
type ShellFilter struct {
	status StatusPlugin
}

func (s *ShellFilter) Process(msgs *Message) *Message {
	if s.status != Started {
		fmt.Println("Shell filter plugin is not running,filter nothing.")
		return msgs
	}

	// 解析yaml结果
	// 1. 解析stage步骤
	stage := msgs.Data.Items["stage"].([]string)
	fmt.Println(stage)
	return msgs
}

func (s *ShellFilter) Start() {
	s.status = Started
	fmt.Println("Shell filter plugin started.")
}

func (s *ShellFilter) Stop() {
	s.status = Stopped
	fmt.Println("Shell filter plugin stopped.")
}

func (s *ShellFilter) Status() StatusPlugin {
	return s.status
}

func (s *ShellFilter) Init() {}
