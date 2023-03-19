/**
  @author: Zero
  @date: 2023/3/19 20:41:39
  @desc: 系统级别配置文件结构体

**/

package properties

// Bootstrap 系统级别配置文件信息
type Bootstrap struct {
	Nacos NacosConfig `yaml:"nacos"`
}

// NacosConfig 配置中心信息
type NacosConfig struct {
	Enable    bool   `yaml:"enable"`    //是否开启读取配置中心
	Addr      string `yaml:"addr"`      //地址
	UserName  string `yaml:"username"`  //用户名
	PassWord  string `yaml:"password"`  //用户密码
	Namespace string `yaml:"namespace"` //命名空间
	Group     string `yaml:"group"`     //分组
	Name      string `yaml:"name"`      //配置文件名称
}
