/**
  @author: Zero
  @date: 2023/3/19 21:10:11
  @desc: 服务配置文件结构体

**/

package properties

import "fmt"

// Application 全局配置
type Application struct {
	Server Server `yaml:"server"`
	Redis  Redis  `yaml:"redis"`
}

// Server 服务配置
type Server struct {
	Name string `yaml:"name"` //服务名
	Addr string `yaml:"addr"` //服务IP与端口
}

// Redis 缓存配置
type Redis struct {
	Host     string `yaml:"host"`     //域名
	Port     int    `yaml:"port"`     //端口
	Password string `yaml:"password"` //登录密码
	Db       int    `yaml:"db"`       //要使用的库索引
}

func (r *Redis) GetAddr() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}
