/**
  @author: Zero
  @date: 2023/3/19 20:24:10
  @desc: 项目配置文件初始化,读取配置中心或本地的配置.

**/

package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	. "zero.com/gtask-common/properties"
	"zero.com/gtask-common/session"
)

// InitConfig 初始化配置文件
func InitConfig() {
	// 加载配置文件
	loadConfig()
}

// 加载配置文件
func loadConfig() {
	if boot := loadBootstrapConfig(); boot.Nacos.Enable {
		//加载远程配置中心Nacos的配置
		loadNacosConfig(boot.Nacos)
	} else {
		// 加载本地配置
		loadLocalConfig()
	}
}

// 加载系统级别配置文件,并且返回
func loadBootstrapConfig() *Bootstrap {
	// 读取系统配置文件
	boot := &Bootstrap{}
	configUnmarshal("bootstrap", "yaml", "config/", boot)
	return boot
}

// 读取本地的配置文件
func loadLocalConfig() {
	app := &Application{}
	v := configUnmarshal("application", "yaml", "config/", app)
	//将配置属性对象 挂载在全局作用域
	session.Configure = app

	// 注册监听回调,监听配置文件的修改.
	v.OnConfigChange(func(in fsnotify.Event) {
		// 修改后实时刷新到程序中
		if err := v.Unmarshal(app); err != nil {
			// flushed 刷新
			fmt.Printf("application Config Change Flushed Error %s \n", err)
		}
	})
	// 开启监听
	v.WatchConfig()
}

// 读取远程配置中心Nacos的配置
func loadNacosConfig(nacos NacosConfig) {
	//TODO 后续再实现
}

// 读取配置文件,映射到指定的结构体中
// fileName 文件名
// fileType 文件类型
// filePath 配置文件所在目录
// payload  要映射的结构体指针
func configUnmarshal(fileName, fileType, filePath string, payload interface{}) *viper.Viper {
	instance := viper.New()
	instance.SetConfigName(fileName)
	instance.SetConfigType(fileType)
	instance.AddConfigPath(filePath)
	// 加载配置文件
	if err := instance.ReadInConfig(); err != nil {
		panic(fmt.Errorf("load %s Config Error: %w", fileName, err))
	}
	// 加载到结构体上
	if err := instance.Unmarshal(payload); err != nil {
		panic(fmt.Errorf("unmarshal %s Config Error: %w", fileName, err))
	}
	return instance
}
