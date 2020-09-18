// 工厂模式生产具体的类
package plugin

import "reflect"

// ============================根据接口实现struct==============================
// 接着，我们定义input、filter、output三类插件接口的具体实现：
// input插件名称与类型的映射关系，主要用于通过反射创建input对象
var inputNames = make(map[string]reflect.Type

// 插件抽象工厂接口
type Factory interface {
	Create(conf Config) []Plugin
}

// check插件工厂，实现Factory接口
type CheckFactory struct{}

// 读取配置，通过反射机制进行对象实例化
func (i *CheckFactory) Create(conf Config) Plugin {
	t, _ := inputNames[conf.Name]
	p := reflect.New(t).Interface().(Plugin)
	// 返回插件实例前调用Init函数，完成相关初始化方法
	p.Init()
	return p
}

// input插件工厂对象，实现Factory接口
type InputFactory struct{}

// 读取配置，通过反射机制进行对象实例化
func (i *InputFactory) Create(conf Config) Plugin {
	t, _ := inputNames[conf.Name]
	p := reflect.New(t).Interface().(Plugin)
	// 返回插件实例前调用Init函数，完成相关初始化方法
	p.Init()
	return p
}

// filter插件工厂对象，实现Factory接口
type FilterFactory struct{}

// 读取配置，通过反射机制进行对象实例化
func (i *FilterFactory) Create(conf Config) Plugin {
	t, _ := inputNames[conf.Name]
	p := reflect.New(t).Interface().(Plugin)
	// 返回插件实例前调用Init函数，完成相关初始化方法
	p.Init()
	return p
}

// output插件工厂对象，实现Factory接口
type OutputFactory struct{}

// 读取配置，通过反射机制进行对象实例化
func (i *OutputFactory) Create(conf Config) Plugin {
	t, _ := inputNames[conf.Name]
	p := reflect.New(t).Interface().(Plugin)
	// 返回插件实例前调用Init函数，完成相关初始化方法
	p.Init()
	return p
}
