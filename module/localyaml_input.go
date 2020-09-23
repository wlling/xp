package module

import (
	"fmt"
	"reflect"

	. "github.com/devopsxp/xp/plugin"
	"github.com/spf13/viper"
)

func init() {
	AddInput("localyaml", reflect.TypeOf(LocalYamlInput{}))
}

type LocalYaml struct {
	data map[string]interface{}
}

func (l *LocalYaml) Get() {
	l.data = viper.AllSettings()
}

type LocalYamlInput struct {
	status StatusPlugin
	yaml   LocalYaml
}

func (l *LocalYamlInput) Receive() *Message {
	l.yaml.Get()
	if l.status != Started {
		fmt.Println("LocalYaml input plugin is not running,input nothing.")
		return nil
	}

	return Builder().WithItemInterface(l.yaml.data).Build()
}

func (l *LocalYamlInput) Start() {
	l.status = Started
	fmt.Println("LocalYamlInput plugin started.")
}

func (l *LocalYamlInput) Stop() {
	l.status = Stopped
	fmt.Println("LocalYamlInput plugin stopped.")
}

func (l *LocalYamlInput) Status() StatusPlugin {
	return l.status
}

// LocalYamlInput的Init函数实现
func (l *LocalYamlInput) Init() {
	l.yaml.data = make(map[string]interface{})
}
