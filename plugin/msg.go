package plugin

// 消息状态定义
type StatusType uint8

const (
	Ok StatusType = iota
	Error
)

type Message struct {
	Type    PluginType
	Content string `json:"content"`
	Status  StatusType
}
