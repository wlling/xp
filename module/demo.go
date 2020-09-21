package module

import (
	"fmt"
	"reflect"
	"strings"

	. "github.com/devopsxp/xp/plugin"
)

func init() {
	// 初始化check插件映射关系表
	AddCheck("ssh", reflect.TypeOf(SshCheck{}))
	// 初始化input插件映射关系表
	AddInput("hello", reflect.TypeOf(HelloInput{}))
	// 初始化filter插件映射关系表
	AddFilter("upper", reflect.TypeOf(UpperFilter{}))
	// 初始化output插件映射关系表
	AddOutput("console", reflect.TypeOf(ConsoleOutput{}))
}

// ssh主机check插件
type SshCheck struct {
	status StatusPlugin
}

func (s *SshCheck) Conn() *Message {
	if s.status != Started {
		fmt.Println("Hello input plugin is not running,input nothing.")
		return nil
	}

	fmt.Println("Check target is connect")

	// 造假数据
	return Builder().WithRaw("{'name':'xp'}").WithItems("thisis", "world").WithTarget([]string{"127.0.0.1", "192.168.0.1"}).WithStatus(Ok).Build()
}

func (s *SshCheck) Start() {
	s.status = Started
	fmt.Println("Check SshCheck plugin started.")
}

func (s *SshCheck) Stop() {
	s.status = Stopped
	fmt.Println("Check SshCheck plugin stopped.")
}

func (s *SshCheck) Status() StatusPlugin {
	return s.status
}

func (s *SshCheck) Init() {
	fmt.Println("Get machine and connecting test init")
}

// Hello input插件，接收“Hello World”消息
type HelloInput struct {
	status StatusPlugin
}

func (h *HelloInput) Receive() *Message {
	// 如果插件未启动，则返回nil
	if h.status != Started {
		fmt.Println("Hello input plugin is not running,input nothing.")
		return nil
	}
	return Builder().WithRaw("{'name':'xp'}").WithItems("thisis", "world").WithTarget([]string{"127.0.0.1", "192.168.0.1"}).WithStatus(Ok).Build()
}

func (h *HelloInput) Start() {
	h.status = Started
	fmt.Println("Hello input plugin started.")
}

func (h *HelloInput) Stop() {
	h.status = Stopped
	fmt.Println("Hello input plugin stopped.")
}

func (h *HelloInput) Status() StatusPlugin {
	return h.status
}

func (h *HelloInput) Init() {}

// Upper filter插件，将消息全部字母转成大写
type UpperFilter struct {
	status StatusPlugin
}

func (u *UpperFilter) Process(msgs *Message) *Message {
	if u.status != Started {
		fmt.Println("Upper filter plugin is not running ,filter nothing.")
		return msgs
	}

	for i, val := range msgs.Data.Target {
		msgs.Data.Target[i] = strings.ToUpper(val)
	}
	return msgs
}

func (u *UpperFilter) Start() {
	u.status = Started
	fmt.Println("Upper filter plugin started.")
}

func (u *UpperFilter) Stop() {
	u.status = Stopped
	fmt.Println("Upper filter plugin stopped.")
}

func (u *UpperFilter) Status() StatusPlugin {
	return u.status
}

func (u *UpperFilter) Init() {}

// Console output插件，将消息输出到控制台上
type ConsoleOutput struct {
	status StatusPlugin
}

func (c *ConsoleOutput) Send(msgs *Message) {
	if c.status != Started {
		fmt.Println("Console output is not running, output nothing.")
		return
	}
	fmt.Printf("Output:\n\tHeader: %+v, Body: %+v\n", msgs.Data.Raw, msgs.Data.Target)
}

func (c *ConsoleOutput) Start() {
	c.status = Started
	fmt.Println("Console output plugin started.")
}

func (c *ConsoleOutput) Stop() {
	c.status = Stopped
	fmt.Println("Console output plugin stopped.")
}

func (c *ConsoleOutput) Status() StatusPlugin {
	return c.status
}

func (c *ConsoleOutput) Init() {}
