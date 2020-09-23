package plugin

import "sync"

// 消息状态定义
type StatusType uint8

const (
	Ok StatusType = iota
	Error
)

type Message struct {
	Env    *Env
	Data   *Data
	Status StatusType
}

type Env struct {
	Type PluginType
}

type Data struct {
	Raw    string                 `json:"raw"` // 原始文档
	Items  map[string]interface{} // 详细参数配置
	Target []string               // 目标服务器
}

// Message对象的Builder对象
type builder struct {
	once *sync.Once
	msg  *Message
}

// 返回Builder对象，工厂模式
func Builder() *builder {
	return &builder{
		once: &sync.Once{},
		msg: &Message{
			Data: &Data{},
			Env:  &Env{},
		},
	}
}

// 建造者模式
func (b *builder) WithRaw(info string) *builder {
	b.msg.Data.Raw = info
	return b
}

func (b *builder) WithStatus(status StatusType) *builder {
	b.msg.Status = status
	return b
}

func (b *builder) WithItemInterface(data map[string]interface{}) *builder {
	b.msg.Data.Items = data
	return b
}

func (b *builder) WithItems(key string, value interface{}) *builder {
	b.once.Do(func() {
		b.msg.Data.Items = make(map[string]interface{})
	})
	b.msg.Data.Items[key] = value
	return b
}

func (b *builder) WithTarget(info []string) *builder {
	b.msg.Data.Target = info
	return b
}

func (b *builder) WithEnv(info PluginType) *builder {
	b.msg.Env.Type = info
	return b
}

// 创建Message对象，在最后一步调用
func (b *builder) Build() *Message {
	return b.msg
}
