package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// TODO: CLI+Viper
func main_backup() {
	// 获取项目的执行路径
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	config := viper.New()

	config.AddConfigPath(path)      // 设置读取文件的路径
	config.SetConfigName("example") // 设置读取的文件名
	config.SetConfigType("yaml")    // 设置文件的类型

	// 尝试进行配置读取
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	// 打印文件读取的内容
	fmt.Println(config.AllSettings())
}
